package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"

	"knight-hacks-2024/services"
)

var validProjectColumns = map[string]bool{
	"id":          true,
	"name":        true,
	"description": true,
	"is_remote":   true,
	"location":    true,
}

func (s *postgreSQL) PostProject(ctx context.Context, project *services.Project) error {
	query := `
	INSERT INTO project (id, name, description, is_remote, location)
	VALUES ($1, $2, $3, $4, $5);`

	if _, err := s.db.ExecContext(ctx, query, project.ID, project.Name, project.Description, project.IsRemote, project.Location); err != nil {
		log.Printf("error posting project to Postgres: %s", err)
		return err
	}

	if err := s.insertProjectRoles(ctx, project.ID, project.Roles); err != nil {
		return err
	}

	return nil
}

func (s *postgreSQL) insertProjectRoles(ctx context.Context, projectID int, roles []*services.Role) error {
	for _, role := range roles {
		query := `
		INSERT INTO project_roles (project_id, role_id)
		VALUES ($1, $2);`
		if _, err := s.db.ExecContext(ctx, query, projectID, role.ID); err != nil {
			log.Printf("error associating project with role in Postgres: %s", err)
			return err
		}
	}
	return nil
}

func (s *postgreSQL) GetProject(ctx context.Context, filter map[string]string) (*services.Project, error) {
	query := `SELECT id, name, description, is_remote, location FROM project`

	conditions := []string{}
	var args []interface{}
	i := 1
	for col, val := range filter {
		if !validProjectColumns[col] {
			return nil, fmt.Errorf("invalid filter column %s", col)
		}
		conditions = append(conditions, fmt.Sprintf("%s = $%d", col, i))
		args = append(args, val)
		i++
	}

	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	project := &services.Project{}

	row := s.db.QueryRowContext(ctx, query, args...)
	if err := row.Scan(
		&project.ID,
		&project.Name,
		&project.Description,
		&project.IsRemote,
		&project.Location,
	); err != nil {
		if err == sql.ErrNoRows {
			log.Printf("error: no project found for the given filters")
			return nil, err
		}
		log.Printf("error scanning project: %s", err)
		return nil, err
	}

	if err := s.getAssociatedRoles(ctx, project); err != nil {
		return nil, err
	}

	return project, nil
}

func (s *postgreSQL) getAssociatedRoles(ctx context.Context, project *services.Project) error {
	query := `
	SELECT r.id, r.name  -- Adjust this if the Role struct has different fields
	FROM role r
	JOIN project_roles pr ON r.id = pr.role_id
	WHERE pr.project_id = $1;`

	rows, err := s.db.QueryContext(ctx, query, project.ID)
	if err != nil {
		log.Printf("error retrieving roles for project: %s", err)
		return err
	}
	defer rows.Close()

	var roles []*services.Role
	for rows.Next() {
		var role services.Role
		if err := rows.Scan(&role.ID, &role.Name); err != nil {
			log.Printf("error scanning role: %s", err)
			return err
		}
		roles = append(roles, &role)
	}
	project.Roles = roles

	return nil
}

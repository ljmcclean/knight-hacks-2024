package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"knight-hacks-2024/services"
	"log"
	"strings"
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

	_, err := s.db.ExecContext(ctx, query, project.ID, project.Name, project.Description, project.IsRemote, project.Location)
	if err != nil {
		log.Printf("error posting project to Postgres: %s", err)
	}
	return err
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
	err := row.Scan(
		&project.ID,
		&project.Name,
		&project.Description,
		&project.IsRemote,
		&project.Location,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("error: no project found for the given filters")
			return nil, err
		}
		log.Printf("error scanning project: %s", err)
		return nil, err
	}

	return project, nil
}

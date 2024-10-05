package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/ljmcclean/knight-hacks-2024/services"

	"github.com/lib/pq"
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
	INSERT INTO project (id, name, description, is_remote, location, skills)
	VALUES ($1, $2, $3, $4, $5, $6);`

	if _, err := s.db.ExecContext(ctx, query, project.ID, project.Name, project.Description, project.IsRemote, project.Location, "{"+strings.Join(project.Skills, ", ")+"}"); err != nil {
		log.Printf("error posting project to Postgres: %s", err)
		return err
	}

	return nil
}

func (s *postgreSQL) GetProject(ctx context.Context, filter map[string]string) (*services.Project, error) {
	query := `SELECT id, name, description, is_remote, location, skills FROM project`

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
		pq.Array(&project.Skills),
	); err != nil {
		if err == sql.ErrNoRows {
			log.Printf("error: no project found for the given filters")
			return nil, err
		}
		log.Printf("error scanning project: %s", err)
		return nil, err
	}

	return project, nil
}

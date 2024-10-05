package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"knight-hacks-2024/services"
	"log"
	"strings"
)

var validRoleColumns = map[string]bool{
	"id":          true,
	"name":        true,
	"description": true,
}

func (s *postgreSQL) PostRole(ctx context.Context, role *services.Role) error {
	query := `
	INSERT INTO role (id, name, description)
	VALUES ($1, $2, $3);`

	_, err := s.db.ExecContext(ctx, query, role.ID, role.Name, role.Description)
	if err != nil {
		log.Printf("error posting role to Postgres: %s", err)
		return err
	}
	return nil
}

func (s *postgreSQL) GetRole(ctx context.Context, filter map[string]string) (*services.Role, error) {
	query := `SELECT id, name, description FROM role`

	conditions := []string{}
	var args []interface{}
	i := 1
	for col, val := range filter {
		if !validRoleColumns[col] {
			return nil, fmt.Errorf("invalid filter column %s", col)
		}
		conditions = append(conditions, fmt.Sprintf("%s = $%d", col, i))
		args = append(args, val)
		i++
	}

	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	role := &services.Role{}

	row := s.db.QueryRowContext(ctx, query, args...)
	err := row.Scan(
		&role.ID,
		&role.Name,
		&role.Description,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("error: no role found for the given filters")
			return nil, err
		}
		log.Printf("error scanning role: %s", err)
		return nil, err
	}

	return role, nil
}

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

var validProfileColumns = map[string]bool{
	"email":       true,
	"id":          true,
	"name":        true,
	"password":    true,
	"description": true,
	"location":    true,
}

func (s *postgreSQL) PostProfile(ctx context.Context, profile *services.Profile) error {
	query := `
	INSERT INTO profile (name, email, password, description, location, skills, project_ids)
	VALUES ($1, $2, $3, $4, $5, $6, $7);`

	if _, err := s.db.ExecContext(ctx, query, profile.Name, profile.Email, profile.Password, profile.Description, profile.Location, "{"+strings.Join(profile.Skills, ", ")+"}", "{"+strings.Join(profile.Projects, ", ")+"}"); err != nil {
		log.Printf("error posting profile to Postgres: %s", err)
		return err
	}

	return nil
}

func (s *postgreSQL) GetProfile(ctx context.Context, filter map[string]string) (*services.Profile, error) {
	query := `SELECT id, name, email, description, location, password, skills, project_ids FROM profile`

	conditions := []string{}
	var args []interface{}
	i := 1
	for col, val := range filter {
		if !validProfileColumns[col] {
			return nil, fmt.Errorf("invalid filter column %s", col)
		}
		conditions = append(conditions, fmt.Sprintf("%s = $%d", col, i))
		args = append(args, val)
		i++
	}
	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	profile := &services.Profile{}

	row := s.db.QueryRowContext(ctx, query, args...)
	if err := row.Scan(
		&profile.ID,
		&profile.Name,
		&profile.Email,
		&profile.Description,
		&profile.Location,
		&profile.Password,
		pq.Array(&profile.Skills),
		pq.Array(&profile.Projects),
	); err != nil {
		if err == sql.ErrNoRows {
			log.Printf("error: no profile found for the given filters")
			return nil, err
		}
		log.Printf("error scanning profile: %s", err)
		return nil, err
	}

	return profile, nil
}

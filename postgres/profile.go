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
	VALUES ($1, $2, $3, $4, $5, $6);`

	if _, err := s.db.ExecContext(ctx, query, profile.Name, profile.Email, profile.Password, profile.Description, profile.Location, "{"+strings.Join(profile.Skills, ", ")+"}"); err != nil {
		log.Printf("error posting profile to Postgres: %s", err)
		return err
	}

	return nil
}

func (s *postgreSQL) UpdateProfile(ctx context.Context, profile *services.Profile) error {
	query := `
	UPDATE profile
	SET name = $2, email = $3, password = $4, description = $5, location = $6, skills = $7
	WHERE id = $1;`

	if _, err := s.db.ExecContext(ctx, query, profile.ID, profile.Name, profile.Email, profile.Password, profile.Description, profile.Location, pq.Array(profile.Skills)); err != nil {
		log.Printf("error updating profile in Postgres: %s", err)
		return err
	}

	return nil
}

func (s *postgreSQL) GetProfile(ctx context.Context, filter map[string]string) (*services.Profile, error) {
	query := `SELECT id, name, email, description, location, password, skills FROM profile`

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

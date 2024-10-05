package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"

	"knight-hacks-2024/services"
)

// Valid column names for selection
var validColumns = map[string]bool{
	"email":    true,
	"id":       true,
	"name":     true,
	"password": true,
}

// TODO: Verify email is valid
func (s *postgreSQL) PostProfile(ctx context.Context, profile *services.Profile) error {
	query := `
	INSERT INTO profile (name, email, password)
	VALUES ($1, $2, $3);`
	_, err := s.db.ExecContext(ctx, query, profile.Name, profile.Email, profile.Password)
	if err != nil {
		log.Printf("error posting profile to Postgres: %s", err)
	}
	return err
}

func (s *postgreSQL) GetProfile(ctx context.Context, fields []string, filter map[string]string) (*services.Profile, error) {
	for _, field := range fields {
		if !validColumns[field] {
			return nil, fmt.Errorf("invalid field column %s", field)
		}
	}

	query := fmt.Sprintf("SELECT %s FROM profile ", strings.Join(fields, ", "))

	// Build conditions into query
	conditions := []string{}
	var args []interface{}
	i := 1
	for col, val := range filter {
		if !validColumns[col] {
			return nil, fmt.Errorf("invalid filter column %s", col)
		}
		conditions = append(conditions, fmt.Sprintf("%s = $%d", col, i))
		args = append(args, val)
		i++
	}
	if len(conditions) > 0 {
		query += " WHERE "
		query += strings.Join(conditions, " AND ")
	}

	profile := &services.Profile{}
	var scanDest []interface{}

	for _, field := range fields {
		switch field {
		case "id":
			scanDest = append(scanDest, &profile.ID)
		case "name":
			scanDest = append(scanDest, &profile.Name)
		case "email":
			scanDest = append(scanDest, &profile.Email)
		case "password":
			scanDest = append(scanDest, &profile.Password)
		}
	}

	row := s.db.QueryRowContext(ctx, query, args...)
	err := row.Scan(scanDest...)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("error no profile found for given filters")
		} else {
			log.Printf("error scanning profile: %s", err)
		}
		return nil, err
	}

	return profile, nil
}

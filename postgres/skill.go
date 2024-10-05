package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"knight-hacks-2024/services"
	"log"
	"strings"
)

var validSkillColumns = map[string]bool{
	"id":               true,
	"name":             true,
	"proficiency":      true,
	"years_experience": true,
}

func (s *postgreSQL) PostSkill(ctx context.Context, skill *services.Skill) error {
	query := `
	INSERT INTO skill (id, name, proficiency, years_experience)
	VALUES ($1, $2, $3, $4);`

	_, err := s.db.ExecContext(ctx, query, skill.ID, skill.Name, skill.Proficiency, skill.YearsExperience)
	if err != nil {
		log.Printf("error posting skill to Postgres: %s", err)
		return err
	}
	return nil
}

func (s *postgreSQL) GetSkill(ctx context.Context, filter map[string]string) (*services.Skill, error) {
	query := `SELECT id, name, proficiency, years_experience FROM skill`

	conditions := []string{}
	var args []interface{}
	i := 1
	for col, val := range filter {
		if !validSkillColumns[col] {
			return nil, fmt.Errorf("invalid filter column %s", col)
		}
		conditions = append(conditions, fmt.Sprintf("%s = $%d", col, i))
		args = append(args, val)
		i++
	}

	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	skill := &services.Skill{}

	row := s.db.QueryRowContext(ctx, query, args...)
	err := row.Scan(
		&skill.ID,
		&skill.Name,
		&skill.Proficiency,
		&skill.YearsExperience,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("error: no skill found for the given filters")
			return nil, err
		}
		log.Printf("error scanning skill: %s", err)
		return nil, err
	}

	return skill, nil
}

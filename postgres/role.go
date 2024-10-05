package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"

	"knight-hacks-2024/services"
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

	if _, err := s.db.ExecContext(ctx, query, role.ID, role.Name, role.Description); err != nil {
		log.Printf("error posting role to Postgres: %s", err)
		return err
	}

	if err := s.insertRoleSkills(ctx, role.ID, role.Skills); err != nil {
		return err
	}

	return nil
}

func (s *postgreSQL) insertRoleSkills(ctx context.Context, roleID int, skills []*services.Skill) error {
	for _, skill := range skills {
		query := `
		INSERT INTO role_skills (role_id, skill_id)
		VALUES ($1, $2);`
		if _, err := s.db.ExecContext(ctx, query, roleID, skill.ID); err != nil {
			log.Printf("error associating role with skill in Postgres: %s", err)
			return err
		}
	}
	return nil
}

func (s *postgreSQL) GetRole(ctx context.Context, filter map[string]string) (*services.Role, error) {
	query := `SELECT r.id, r.name, r.description FROM role r`

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
	if err := row.Scan(
		&role.ID,
		&role.Name,
		&role.Description,
	); err != nil {
		if err == sql.ErrNoRows {
			log.Printf("error: no role found for the given filters")
			return nil, err
		}
		log.Printf("error scanning role: %s", err)
		return nil, err
	}

	if err := s.getRoleAssociatedSkills(ctx, role); err != nil {
		return nil, err
	}

	return role, nil
}

func (s *postgreSQL) getRoleAssociatedSkills(ctx context.Context, role *services.Role) error {
	query := `
	SELECT s.id, s.name, s.proficiency, s.years_experience
	FROM skill s
	JOIN role_skills rs ON s.id = rs.skill_id
	WHERE rs.role_id = $1;`

	rows, err := s.db.QueryContext(ctx, query, role.ID)
	if err != nil {
		log.Printf("error retrieving skills for role: %s", err)
		return err
	}
	defer rows.Close()

	var skills []*services.Skill
	for rows.Next() {
		skill := &services.Skill{}
		if err := rows.Scan(
			&skill.ID,
			&skill.Name,
			&skill.Proficiency,
			&skill.YearsExperience,
		); err != nil {
			log.Printf("error scanning skill: %s", err)
			return err
		}
		skills = append(skills, skill)
	}

	role.Skills = skills

	return nil
}

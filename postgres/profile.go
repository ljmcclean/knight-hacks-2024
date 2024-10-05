package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"

	"knight-hacks-2024/services"

	"github.com/google/uuid"
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
	INSERT INTO profile (id, name, email, password, description, location)
	VALUES ($1, $2, $3, $4, $5, $6);`

	if _, err := s.db.ExecContext(ctx, query, profile.ID, profile.Name, profile.Email, profile.Password, profile.Description, profile.Location); err != nil {
		log.Printf("error posting profile to Postgres: %s", err)
		return err
	}

	if err := s.insertProfileProjects(ctx, profile.ID, profile.Projects); err != nil {
		return err
	}

	if err := s.insertProfileSkills(ctx, profile.ID, profile.Skills); err != nil {
		return err
	}

	return nil
}

func (s *postgreSQL) insertProfileProjects(ctx context.Context, profileID uuid.UUID, projects []*services.Project) error {
	for _, project := range projects {
		query := `
		INSERT INTO profile_projects (profile_id, project_id)
		VALUES ($1, $2);`
		if _, err := s.db.ExecContext(ctx, query, profileID, project.ID); err != nil {
			log.Printf("error associating profile with project in Postgres: %s", err)
			return err
		}
	}
	return nil
}

func (s *postgreSQL) insertProfileSkills(ctx context.Context, profileID uuid.UUID, skills []*services.Skill) error {
	for _, skill := range skills {
		query := `
		INSERT INTO profile_skills (profile_id, skill_id)
		VALUES ($1, $2);`
		if _, err := s.db.ExecContext(ctx, query, profileID, skill.ID); err != nil {
			log.Printf("error associating profile with skill in Postgres: %s", err)
			return err
		}
	}
	return nil
}

func (s *postgreSQL) GetProfile(ctx context.Context, filter map[string]string) (*services.Profile, error) {
	query := `SELECT id, name, email, description, location, password FROM profile`

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
	); err != nil {
		if err == sql.ErrNoRows {
			log.Printf("error: no profile found for the given filters")
			return nil, err
		}
		log.Printf("error scanning profile: %s", err)
		return nil, err
	}

	if err := s.getAssociatedProjects(ctx, profile); err != nil {
		return nil, err
	}

	if err := s.getAssociatedSkills(ctx, profile); err != nil {
		return nil, err
	}

	return profile, nil
}

func (s *postgreSQL) getAssociatedProjects(ctx context.Context, profile *services.Profile) error {
	query := `
	SELECT p.id, p.name, p.description, p.is_remote, p.location
	FROM project p
	JOIN profile_projects pp ON p.id = pp.project_id
	WHERE pp.profile_id = $1;`

	rows, err := s.db.QueryContext(ctx, query, profile.ID)
	if err != nil {
		log.Printf("error retrieving projects for profile: %s", err)
		return err
	}
	defer rows.Close()

	var projects []*services.Project
	for rows.Next() {
		var project services.Project
		if err := rows.Scan(&project.ID, &project.Name, &project.Description, &project.IsRemote, &project.Location); err != nil {
			log.Printf("error scanning project: %s", err)
			return err
		}
		projects = append(projects, &project)
	}
	profile.Projects = projects

	return nil
}

func (s *postgreSQL) getAssociatedSkills(ctx context.Context, profile *services.Profile) error {
	query := `
	SELECT s.id, s.name, s.proficiency, s.years_experience
	FROM skill s
	JOIN profile_skills ps ON s.id = ps.skill_id
	WHERE ps.profile_id = $1;`

	rows, err := s.db.QueryContext(ctx, query, profile.ID)
	if err != nil {
		log.Printf("error retrieving skills for profile: %s", err)
		return err
	}
	defer rows.Close()

	var skills []*services.Skill
	for rows.Next() {
		var skill services.Skill
		if err := rows.Scan(&skill.ID, &skill.Name, &skill.Proficiency, &skill.YearsExperience); err != nil {
			log.Printf("error scanning skill: %s", err)
			return err
		}
		skills = append(skills, &skill)
	}
	profile.Skills = skills

	return nil
}

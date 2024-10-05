package services

import "context"

type Skill struct {
	ID              int
	Name            string
	Proficiency     int
	YearsExperience int
}

type SkillService interface {
	PostSkill(context.Context, *Skill) error
	GetSkill(ctx context.Context, filter map[string]string) (*Skill, error)
}

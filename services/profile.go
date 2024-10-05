package services

import (
	"context"

	"github.com/google/uuid"
)

type Profile struct {
	ID             uuid.UUID
	Name           string
	Description    string
	Email          string
	Password       string
	Location       string
	Skills         []*Skill
	OwnedProjects  []*Project
	JoinedProjects []*Project
}

type ProfileService interface {
	PostProfile(context.Context, *Profile) error
	GetProfile(ctx context.Context, fields []string, filter map[string]string) (*Profile, error)
}

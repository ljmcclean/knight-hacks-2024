package services

import (
	"context"

	"github.com/google/uuid"
)

type Project struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	IsRemote    int       `json:"is_remote"`
	Location    string    `json:"location"`
	Skills      []string  `json:"skills"`
	UserID      uuid.UUID `json:"user_id"`
}

type ProjectService interface {
	PostProject(context.Context, *Project) error
	UpdateProject(ctx context.Context, profile *Project) error
	GetProject(ctx context.Context, filter map[string]string) (*Project, error)
	GetUserProjects(ctx context.Context, userID uuid.UUID) ([]*Project, error)
	GetMatchingProjects(ctx context.Context, skills []string) ([]*Project, error)
}

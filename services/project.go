package services

import (
	"context"
)

type Project struct {
	ID          int
	Name        string
	Description string
	IsRemote    bool
	Location    string
	Roles       []*Role
}

type ProjectService interface {
	PostProject(context.Context, *Project) error
	GetProject(ctx context.Context, filter map[string]string) (*Project, error)
}

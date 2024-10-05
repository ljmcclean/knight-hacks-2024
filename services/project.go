package services

import (
	"context"
)

type Project struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	IsRemote    bool     `json:"is_remote"`
	Location    string   `json:"location"`
	Skills      []string `json:"skills"`
}

type ProjectService interface {
	PostProject(context.Context, *Project) error
	GetProject(ctx context.Context, filter map[string]string) (*Project, error)
}

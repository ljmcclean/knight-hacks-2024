package services

import (
	"context"

	"github.com/google/uuid"
)

type Profile struct {
	ID          uuid.UUID
	Name        string
	Description string
	Email       string
	Password    string
	Location    string
	Skills      []*Skill
	Projects    []*Project
}

type ProfileService interface {
	PostProfile(context.Context, *Profile) error
	GetProfile(ctx context.Context, filter map[string]string) (*Profile, error)
}

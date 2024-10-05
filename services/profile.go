package services

import (
	"context"

	"github.com/google/uuid"
)

type Profile struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Description string    `json:"description"`
	Password    string    `json:"password"`
	Location    string    `json:"location"`
	Skills      []string  `json:"skills"`
}

type ProfileService interface {
	PostProfile(context.Context, *Profile) error
	UpdateProfile(context.Context, *Profile) error
	GetProfile(ctx context.Context, filter map[string]string) (*Profile, error)
}

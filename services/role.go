package services

import "context"

type Role struct {
	ID          int
	Name        string
	Description string
	Skills      []*Skill
}

type RoleService interface {
	PostRole(context.Context, *Role) error
	GetRole(ctx context.Context, filter map[string]string) (*Role, error)
}

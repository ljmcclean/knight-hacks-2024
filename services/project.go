package services

import "github.com/google/uuid"

type Project struct {
	ID          uuid.UUID
	Name        string
	Description string
	IsRemote    bool
	Location    string
	Roles       []*Role
}

package services

import (
	"context"
	"github.com/ljmcclean/knight-hacks-2024/config"
	"time"

	"github.com/google/uuid"
)

type Session struct {
	SessionID  string
	ProfileID  uuid.UUID
	AuthLevel  int
	LastAccess time.Time
}

type SessionService interface {
	PostSession(context.Context, *Session) error
	InvalidateSession(ctx context.Context, sessionID string) error
	GetSession(ctx context.Context, sessionID string) (*Session, error)
	CleanupSessions(context.Context, config.Config) error
}

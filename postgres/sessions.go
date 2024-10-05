package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"knight-hacks-2024/config"
	"knight-hacks-2024/services"
	"time"
)

func (s *postgreSQL) PostSession(ctx context.Context, session *services.Session) error {
	query := `
	INSERT INTO session (session_id, profile_id, auth_level, last_access)
	VALUES ($1, $2, $3, $4);`
	_, err := s.db.QueryContext(ctx, query, session.SessionID, session.ProfileID, session.AuthLevel, time.Now())
	if err != nil {
		log.Printf("error posting session to Postgres: %s", err)
	}
	return err
}

func (s *postgreSQL) InvalidateSession(ctx context.Context, sessionID string) error {
	query := `
	UPDATE session 
	SET valid = false
	WHERE session_id = $1`
	_, err := s.db.ExecContext(ctx, query, sessionID)
	return err
}

func (s *postgreSQL) CleanupSessions(ctx context.Context, cfg config.Config) error {
	query := `
	DELETE FROM session
	WHERE valid = false or last_access < $1`
	_, err := s.db.ExecContext(ctx, query, time.Now().Add(-cfg.Session.Lifespan))
	return err
}

func (s *postgreSQL) GetSession(ctx context.Context, sessionID string) (*services.Session, error) {
	query := `
	SELECT profile_id, session_id, auth_level
	FROM session
	WHERE session_id = $1 AND valid = true;`

	var session services.Session

	err := s.db.QueryRowContext(ctx, query, sessionID).Scan(&session.ProfileID, &session.SessionID, &session.AuthLevel)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("session not found or invalid")
		}
		return nil, fmt.Errorf("error querying session: %s", err)
	}

	return &session, nil
}

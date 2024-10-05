package auth

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
	"knight-hacks-2024/services"
	"time"

	"github.com/google/uuid"
)

type contextKey string

const SessionKey = contextKey("session")

func Authenticate(next http.Handler, ctx context.Context, ss services.SessionService, authLevel int) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session_id")
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		sessionID := cookie.Value

		session, err := ss.GetSession(ctx, sessionID)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		if session.AuthLevel >= authLevel {
			// Attach the session to authenticated routes
			ctx := context.WithValue(ctx, SessionKey, session)
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
	})
}

func RegisterSession(ctx context.Context, ss services.SessionService, profileID uuid.UUID) (sessionID string, err error) {
	sessionID, err = generateSessionID()
	if err != nil {
		return "", err
	}

	session := &services.Session{
		ProfileID:  profileID,
		SessionID:  sessionID,
		AuthLevel:  1,
		LastAccess: time.Now(),
	}
	err = ss.PostSession(ctx, session)
	return sessionID, nil
}

func generateSessionID() (string, error) {
	randomBytes := make([]byte, 32)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", fmt.Errorf("failed to generate random bytes")
	}
	sessionID := base64.RawURLEncoding.EncodeToString(randomBytes)
	return sessionID, nil
}

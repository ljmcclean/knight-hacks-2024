package handlers

import (
	"context"
	"net/http"

	"github.com/ljmcclean/knight-hacks-2024/server/auth"
	"github.com/ljmcclean/knight-hacks-2024/services"
)

func PostLogout(ctx context.Context, ss services.SessionService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session := r.Context().Value(auth.SessionKey).(*services.Session)

		err := ss.InvalidateSession(ctx, session.SessionID)
		if err != nil {
			http.Error(w, "Could not invalidate the active session", http.StatusInternalServerError)
		}

		http.SetCookie(w, &http.Cookie{
			Name:     "session_id",
			Value:    "",
			MaxAge:   -1, // Delete the cookie
			Path:     "/",
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteStrictMode,
		})

		http.Redirect(w, r, "/", http.StatusSeeOther)
	})
}

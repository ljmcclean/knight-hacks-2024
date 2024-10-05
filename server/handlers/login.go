package handlers

import (
	"context"
	"log"
	"net/http"
	"knight-hacks-2024/config"
	"knight-hacks-2024/services"

	"github.com/a-h/templ"
	"golang.org/x/crypto/bcrypt"

	"knight-hacks-2024/server/auth"
	"knight-hacks-2024/server/templates"
)

func GetLogin() http.Handler {
	return templ.Handler(templates.Login())
}

func PostLogin(ctx context.Context, cfg *config.Config, ps services.ProfileService, ss services.SessionService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		email := r.Form.Get("email")
		givenPassword := r.Form.Get("password")

		profile, err := ps.GetProfile(ctx,
			[]string{"id", "password"},
			map[string]string{
				"email": email,
			})
		if err != nil {
			log.Printf("error retrieving profile: %s", err)
			http.Error(w, "Account not found", http.StatusUnauthorized)
			return
		}
		actualPassword := profile.Password

		// Verify passwords match
		err = bcrypt.CompareHashAndPassword([]byte(actualPassword), []byte(givenPassword))
		if err != nil {
			http.Error(w, "Invalid login credentials", http.StatusUnauthorized)
			return
		}

		// Create session for user
		sessionID, err := auth.RegisterSession(ctx, ss, profile.ID)
		if err != nil {
			http.Error(w, "Could not generate a session ID", http.StatusInternalServerError)
		}

		http.SetCookie(w, &http.Cookie{
			Name:     "session_id",
			Value:    sessionID,
			Path:     "/",
			MaxAge:   int(cfg.Session.Lifespan),
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteStrictMode,
		})

		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
	})
}

package handlers

import (
	"context"
	"github.com/ljmcclean/knight-hacks-2024/server/templates"
	"github.com/ljmcclean/knight-hacks-2024/services"
	"net/http"

	"github.com/a-h/templ"

	"github.com/ljmcclean/knight-hacks-2024/server/auth"
)

// Protected route with session passed through via context
func GetDashboard(ctx context.Context, ps services.ProfileService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session := r.Context().Value(auth.SessionKey).(*services.Session)

		profile, err := ps.GetProfile(ctx,
			map[string]string{
				"id": session.ProfileID.String(),
			},
		)
		if err != nil {
			http.Error(w, "Account not found", http.StatusUnauthorized)
		}

		templ.Handler(templates.Dashboard(profile)).ServeHTTP(w, r)
	})
}

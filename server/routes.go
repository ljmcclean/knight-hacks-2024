package server

import (
	"context"
	"net/http"
	"knight-hacks-2024/config"
	"knight-hacks-2024/server/auth"
	"knight-hacks-2024/server/handlers"
	"knight-hacks-2024/server/templates"
	"knight-hacks-2024/services"

	"github.com/a-h/templ"
)

func addRoutes(mux *http.ServeMux, ctx context.Context, cfg *config.Config, db services.Database) {
	mux.Handle("/", templ.Handler(templates.Index()))

	mux.Handle("GET /sign-up", handlers.GetSignUp())
	mux.Handle("POST /sign-up", handlers.PostSignUp(ctx, db))

	mux.Handle("GET /login", handlers.GetLogin())
	mux.Handle("POST /login", handlers.PostLogin(ctx, cfg, db, db))

	mux.Handle("GET /dashboard", auth.Authenticate(handlers.GetDashboard(ctx, db), ctx, db, 1))

	mux.Handle("POST /logout", auth.Authenticate(handlers.PostLogout(ctx, db), ctx, db, 0))
}

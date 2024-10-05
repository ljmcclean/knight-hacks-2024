package server

import (
	"context"
	"knight-hacks-2024/config"
	"knight-hacks-2024/server/auth"
	"knight-hacks-2024/server/handlers"
	"knight-hacks-2024/server/templates"
	"knight-hacks-2024/services"
	"net/http"

	"github.com/a-h/templ"
)

func addRoutes(mux *http.ServeMux, ctx context.Context, cfg *config.Config, db services.Database) {
	assetsFS := http.FileServer(http.Dir("server/public"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", assetsFS))

	mux.Handle("/", templ.Handler(templates.Index()))

	mux.Handle("GET /sign-up", handlers.GetSignUp())
	mux.Handle("POST /sign-up", handlers.PostSignUp(ctx, db))

	mux.Handle("GET /login", handlers.GetLogin())
	mux.Handle("POST /login", handlers.PostLogin(ctx, cfg, db, db))

	mux.Handle("GET /dashboard", auth.Authenticate(handlers.GetDashboard(ctx, db), ctx, db, 1))

	mux.Handle("POST /logout", auth.Authenticate(handlers.PostLogout(ctx, db), ctx, db, 0))
}

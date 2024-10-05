package server

import (
	"context"
	"net/http"

	"github.com/ljmcclean/knight-hacks-2024/config"
	"github.com/ljmcclean/knight-hacks-2024/server/auth"
	"github.com/ljmcclean/knight-hacks-2024/server/handlers"
	"github.com/ljmcclean/knight-hacks-2024/server/templates"
	"github.com/ljmcclean/knight-hacks-2024/services"

	"github.com/a-h/templ"

	"github.com/ljmcclean/knight-hacks-2024/server/handlers/api"
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

	mux.Handle("GET /discover", auth.Authenticate(handlers.GetDiscover(ctx, db), ctx, db, 1))

	mux.Handle("POST /logout", auth.Authenticate(handlers.PostLogout(ctx, db), ctx, db, 0))

	// API Endpoints
	mux.Handle("GET /api/profile/{id}", auth.Authenticate(api.GetProfile(ctx, db), ctx, db, 1))
	mux.Handle("POST /api/profile/", auth.Authenticate(api.PostProfile(ctx, db), ctx, db, 1))

	mux.Handle("GET /api/project/{id}", auth.Authenticate(api.GetProject(ctx, db), ctx, db, 1))
	mux.Handle("GET /api/all-user-projects", auth.Authenticate(api.GetUserProjects(ctx, db), ctx, db, 1))
	mux.Handle("GET /api/all-matching-projects", auth.Authenticate(api.GetMatchingProjects(ctx, db), ctx, db, 1))
	mux.Handle("POST /api/project/", auth.Authenticate(api.PostProject(ctx, db), ctx, db, 1))
	mux.Handle("POST /api/new-project/", auth.Authenticate(api.PostNewProject(ctx, db), ctx, db, 1))
}

package server

import (
	"context"
	"net/http"
	"knight-hacks-2024/config"
	"knight-hacks-2024/services"
)

func New(cfg *config.Config, ctx context.Context, db services.Database) *http.Server {
	mux := http.NewServeMux()

	addRoutes(mux, ctx, cfg, db)

	// var handler http.Handler = mux
	// handler = middleware(handler)
	// handler = auth(handler)

	return &http.Server{
		Addr:    cfg.Server.Port,
		Handler: mux,
	}
}

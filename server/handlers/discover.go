package handlers

import (
	"context"
	"github.com/ljmcclean/knight-hacks-2024/server/templates"
	"net/http"

	"github.com/a-h/templ"
)

// Protected route with session passed through via context
func GetDiscover(ctx context.Context) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		templ.Handler(templates.Discover()).ServeHTTP(w, r)
	})
}

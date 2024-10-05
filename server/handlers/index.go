package handlers

import (
	"net/http"
	"knight-hacks-2024/server/templates"

	"github.com/a-h/templ"
)

func GetIndex() http.Handler {
	return templ.Handler(templates.Index())
}

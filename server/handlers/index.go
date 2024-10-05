package handlers

import (
	"github.com/ljmcclean/knight-hacks-2024/server/templates"
	"net/http"

	"github.com/a-h/templ"
)

func GetIndex() http.Handler {
	return templ.Handler(templates.Index())
}

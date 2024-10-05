package handlers

import (
	"context"
	"knight-hacks-2024/server/auth"
	"knight-hacks-2024/server/templates"
	"knight-hacks-2024/services"
	"net/http"
	"regexp"

	"github.com/a-h/templ"
)

var emailRegex, _ = regexp.Compile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

func GetSignUp() http.Handler {
	return templ.Handler(templates.SignUp())
}

func PostSignUp(ctx context.Context, ps services.ProfileService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		name := r.Form.Get("name")
		if name == "" {
			http.Error(w, "Bad Request", http.StatusBadRequest)
		}

		email := r.Form.Get("email")
		if email == "" || !emailRegex.MatchString(email) {
			http.Error(w, "Bad Request", http.StatusBadRequest)
		}

		password := r.Form.Get("password")
		if len(password) < 8 || len(password) > 50 {
			http.Error(w, "Bad Request", http.StatusBadRequest)
		}

		// Securely posts profile to ProfileService
		// TODO: Ensure this action can't fail
		err := auth.RegisterProfile(ctx, ps, name, email, password)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}

		http.Redirect(w, r, "/login", http.StatusFound)
	})
}

package api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/ljmcclean/knight-hacks-2024/server/auth"
	"github.com/ljmcclean/knight-hacks-2024/services"
)

func GetProfile(ctx context.Context, ps services.ProfileService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimPrefix(r.URL.Path, "/api/profile/")
		id := strings.Split(path, "/")[0]

		if id == "" {
			http.Error(w, "Profile ID is missing", http.StatusBadRequest)
			return
		}

		profile, err := ps.GetProfile(ctx, map[string]string{
			"id": id,
		})
		if err != nil {
			http.Error(w, "Profile could not be found", http.StatusNotFound)
		}

		json, err := json.Marshal(profile)
		if err != nil {
			log.Printf("error marshalling profile")
			http.Error(w, "Couldn't marshal profile", http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(json)
	})
}

func PostProfile(ctx context.Context, ps services.ProfileService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session := r.Context().Value(auth.SessionKey).(*services.Session)
		id := session.ProfileID

		r.ParseForm()
		name := r.Form.Get("name")
		email := r.Form.Get("email")
		desc := r.Form.Get("description")
		loc := r.Form.Get("location")
		skillStr := r.Form.Get("skills")
		skills := strings.Split(skillStr, ",")

		profile := &services.Profile{
			ID:          id,
			Name:        name,
			Email:       email,
			Description: desc,
			Location:    loc,
			Skills:      skills,
		}

		err := ps.UpdateProfile(ctx, profile)
		if err != nil {
			log.Printf("error posting profile: %s", err)
			return
		}

		w.WriteHeader(http.StatusOK)
	})
}

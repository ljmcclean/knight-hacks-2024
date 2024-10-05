package api

import (
	"context"
	"encoding/json"
	"github.com/ljmcclean/knight-hacks-2024/services"
	"log"
	"net/http"
	"strings"
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

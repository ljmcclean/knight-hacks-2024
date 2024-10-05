package api

import (
	"context"
	"encoding/json"
	"github.com/ljmcclean/knight-hacks-2024/services"
	"log"
	"net/http"
	"strings"
)

func GetProject(ctx context.Context, ps services.ProjectService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimPrefix(r.URL.Path, "/api/project/")
		id := strings.Split(path, "/")[0]

		if id == "" {
			http.Error(w, "Project ID is missing", http.StatusBadRequest)
			return
		}

		profile, err := ps.GetProject(ctx, map[string]string{
			"id": id,
		})
		if err != nil {
			http.Error(w, "Project could not be found", http.StatusNotFound)
		}

		json, err := json.Marshal(profile)
		if err != nil {
			log.Printf("error marshalling profile")
			http.Error(w, "Couldn't marshal project", http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(json)
	})
}

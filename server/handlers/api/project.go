package api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/ljmcclean/knight-hacks-2024/server/auth"
	"github.com/ljmcclean/knight-hacks-2024/services"
)

func GetProject(ctx context.Context, ps services.ProjectService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimPrefix(r.URL.Path, "/api/project/")
		id := strings.Split(path, "/")[0]

		if id == "" {
			http.Error(w, "Project ID is missing", http.StatusBadRequest)
			return
		}

		project, err := ps.GetProject(ctx, map[string]string{
			"id": id,
		})
		if err != nil {
			http.Error(w, "Project could not be found", http.StatusNotFound)
		}

		json, err := json.Marshal(project)
		if err != nil {
			log.Printf("error marshalling project")
			http.Error(w, "Couldn't marshal project", http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(json)
	})
}

func GetUserProjects(ctx context.Context, ps services.ProjectService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session := r.Context().Value(auth.SessionKey).(*services.Session)
		id := session.ProfileID

		projects, err := ps.GetUserProjects(ctx, id)
		if err != nil {
			http.Error(w, "Projects could not be found", http.StatusNotFound)
		}

		json, err := json.Marshal(projects)
		if err != nil {
			log.Printf("error marshalling profile")
			http.Error(w, "Couldn't marshal project", http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(json)
	})
}

func GetMatchingProjects(ctx context.Context, db services.Database) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session := r.Context().Value(auth.SessionKey).(*services.Session)
		id := session.ProfileID

		profile, err := db.GetProfile(ctx, map[string]string{
			"id": id.String(),
		})
		if err != nil {
			http.Error(w, "Profile could not be found", http.StatusNotFound)
		}

		projects, err := db.GetMatchingProjects(ctx, profile.Skills)
		if err != nil {
			http.Error(w, "Projects could not be found", http.StatusNotFound)
		}

		json, err := json.Marshal(projects)
		if err != nil {
			log.Printf("error marshalling profile")
			http.Error(w, "Couldn't marshal project", http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(json)
	})
}

func PostNewProject(ctx context.Context, ps services.ProjectService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session := r.Context().Value(auth.SessionKey).(*services.Session)
		id := session.ProfileID

		r.ParseForm()
		name := r.Form.Get("name")
		desc := r.Form.Get("description")
		isRemoteStr := r.Form.Get("is_remote")
		isRemote, err := strconv.Atoi(isRemoteStr)
		if err != nil {
			log.Printf("invalid isRemote field")
			http.Error(w, "Could not parse is_remote field", http.StatusBadRequest)
			return
		}
		loc := r.Form.Get("location")
		skillStr := r.Form.Get("skills")
		skills := strings.Split(skillStr, ",")

		project := &services.Project{
			UserID:      id,
			Name:        name,
			Description: desc,
			IsRemote:    isRemote,
			Location:    loc,
			Skills:      skills,
		}

		err = ps.PostProject(ctx, project)
		if err != nil {
			log.Printf("error posting project: %s", err)
			return
		}
	})
}

func PostProject(ctx context.Context, ps services.ProjectService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session := r.Context().Value(auth.SessionKey).(*services.Session)
		id := session.ProfileID

		r.ParseForm()
		name := r.Form.Get("name")
		desc := r.Form.Get("description")
		isRemoteStr := r.Form.Get("is_remote")
		isRemote, err := strconv.Atoi(isRemoteStr)
		if err != nil {
			log.Printf("invalid isRemote field")
			http.Error(w, "Could not parse is_remote field", http.StatusBadRequest)
			return
		}
		loc := r.Form.Get("location")
		skillStr := r.Form.Get("skills")
		skills := strings.Split(skillStr, ",")

		project := &services.Project{
			UserID:      id,
			Name:        name,
			Description: desc,
			IsRemote:    isRemote,
			Location:    loc,
			Skills:      skills,
		}

		err = ps.UpdateProject(ctx, project)
		if err != nil {
			log.Printf("error posting project: %s", err)
			return
		}
	})
}

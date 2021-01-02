// Package projects - Project handlers: GET, POST, PUT and DELETE
package projects

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/JSONhilder/overseer_api/internal/application"
	"github.com/gorilla/mux"
)

type Project struct {
	ID               string `json: "id"`
	UID              string `json: "uid"`
	ProjectName      string `json: "project_name"`
	ProjectDesc      string `json: "project_desc"`
	ProjectCompleted string `json: "project_completed"`
	ProjectTime      string `json: "project_time"`
}

type httpError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func jsonError(w http.ResponseWriter, err string, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	jsonErr := httpError{Code: 400, Message: err}
	json.NewEncoder(w).Encode(jsonErr)
}

// GetProjects - List all users projects
func GetProjects(app *application.Application) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		app.LOG.Info("GetAll endpoint hit: /api/projects")

		var projects []Project
		sqlQuery := `SELECT * FROM projects;`
		rows, err := app.DB.Client.Query(sqlQuery)

		defer rows.Close()
		for rows.Next() {
			var p Project

			err = rows.Scan(&p.ID, &p.UID, &p.ProjectName, &p.ProjectDesc, &p.ProjectTime, &p.ProjectCompleted)
			if err != nil {
				jsonError(w, err.Error(), 400)
			}

			projects = append(projects, p)
		}

		json.NewEncoder(w).Encode(projects)
	}
}

// GetProject - get single project with id param
func GetProject(app *application.Application) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		app.LOG.Info("GetProject endpoint hit: /api/projects/:id")

		// Get all variable/parameters passed returns a map
		vars := mux.Vars(r)
		key := vars["id"]

		sqlQuery := `SELECT * FROM projects WHERE id=$1`
		row, err := app.DB.Client.Query(sqlQuery, key)
		if err != nil {
			jsonError(w, err.Error(), 400)
		}

		var p Project

		for row.Next() {
			err = row.Scan(&p.ID, &p.UID, &p.ProjectName, &p.ProjectDesc, &p.ProjectTime, &p.ProjectCompleted)
			if err != nil {
				jsonError(w, err.Error(), 400)
			}
		}

		json.NewEncoder(w).Encode(p)
	}
}

// CreateProject - Create new project
func CreateProject(app *application.Application) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		app.LOG.Info("CreateProject endpoint hit: /api/projects")

		sqlQuery := `
		INSERT INTO projects (u_id, project_name, project_desc, project_time, project_completed)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id`

		id := 0
		err := app.DB.Client.QueryRow(sqlQuery, false, "test proj2", "testing2", "11 minutes", false).Scan(&id)
		if err != nil {
			jsonError(w, err.Error(), 400)
		} else {
			fmt.Println("New record ID is: ", id)
		}
	}
}

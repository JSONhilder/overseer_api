// Package projects - Project handlers: GET, POST, PUT and DELETE
package projects

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/JSONhilder/overseer_api/internal/application"
	"github.com/JSONhilder/overseer_api/internal/utils"
	"github.com/gorilla/mux"
)

/*
	@TODO
	- Add uid to query filters
	- Possibly get rid of *client* in db.Client.Query
	- Set serverRes to a boolean value
*/

// Project - model for project entries
type Project struct {
	ID               string `json: "id"`
	UID              string `json: "uid"`
	ProjectName      string `json: "project_name"`
	ProjectDesc      string `json: "project_desc"`
	ProjectCompleted bool   `json: "project_completed"`
	ProjectTime      string `json: "project_time"`
}

type serverRes struct {
	Res string `json: res`
}

// GetProjects - List all users projects
func GetProjects(app *application.Application) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		app.Logger.Info("GetAll endpoint hit: /api/projects")

		var projects []Project
		sqlQuery := `SELECT * FROM projects;`
		rows, err := app.Db.Client.Query(sqlQuery)

		defer rows.Close()
		for rows.Next() {
			var p Project

			err = rows.Scan(&p.ID, &p.UID, &p.ProjectName, &p.ProjectDesc, &p.ProjectTime, &p.ProjectCompleted)
			if err != nil {
				utils.JSONError(w, err.Error(), 400)
			}

			projects = append(projects, p)
		}

		json.NewEncoder(w).Encode(projects)
	}
}

// GetProject - get single project with id param
func GetProject(app *application.Application) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		app.Logger.Info("GetProject endpoint hit: /api/projects/:id")

		// Get all variable/parameters passed returns a map
		vars := mux.Vars(r)
		key := vars["id"]

		sqlQuery := `SELECT * FROM projects WHERE id=$1`
		row, err := app.Db.Client.Query(sqlQuery, key)
		if err != nil {
			utils.JSONError(w, err.Error(), 400)
		}

		var p Project

		for row.Next() {
			err = row.Scan(&p.ID, &p.UID, &p.ProjectName, &p.ProjectDesc, &p.ProjectTime, &p.ProjectCompleted)
			if err != nil {
				utils.JSONError(w, err.Error(), 400)
			}
		}

		json.NewEncoder(w).Encode(p)
	}
}

// CreateProject - Create new project
func CreateProject(app *application.Application) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		app.Logger.Info("CreateProject endpoint hit: /api/projects")

		var id int
		var p Project
		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			utils.JSONError(w, "Bad Request", 400)
			return
		}
		defer r.Body.Close()

		sqlQuery := `
		INSERT INTO projects (u_id, project_name, project_desc, project_time, project_completed)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id`

		err = app.Db.Client.QueryRow(sqlQuery, p.UID, p.ProjectName, p.ProjectDesc, p.ProjectTime, p.ProjectCompleted).Scan(&id)
		if err != nil {
			utils.JSONError(w, err.Error(), 400)
		} else {
			str := fmt.Sprintf("Successfully created project: %v", id)
			res := serverRes{Res: str}
			json.NewEncoder(w).Encode(res)
		}
	}
}

// DeleteProject - Delete project with passed Id param
func DeleteProject(app *application.Application) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		app.Logger.Info("DeleteProject endpoint hit: /api/projects/:id")

		// Get all variable/parameters passed returns a map
		vars := mux.Vars(r)
		key := vars["id"]

		sqlQuery := `DELETE FROM projects WHERE id=$1`
		_, err := app.Db.Client.Query(sqlQuery, key)
		if err != nil {
			utils.JSONError(w, err.Error(), 400)
			return
		}

		str := fmt.Sprintf("Successfully deleted project: %v", key)
		res := serverRes{Res: str}
		json.NewEncoder(w).Encode(res)
	}
}

// UpdateProject - Update project with passed Id param
func UpdateProject(app *application.Application) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		app.Logger.Info("UpdateProject endpoint hit: /api/projects/:id")

		var p Project
		// Get all variable/parameters passed returns a map
		vars := mux.Vars(r)
		key := vars["id"]

		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			utils.JSONError(w, "Bad Request", 400)
			return
		}
		defer r.Body.Close()

		sqlQuery := `UPDATE projects
		SET u_id = $1, project_name = $2, project_desc = $3, project_time = $4, project_completed = $5
		WHERE id = $6`

		_, err = app.Db.Client.Query(sqlQuery, p.UID, p.ProjectName, p.ProjectDesc, p.ProjectTime, p.ProjectCompleted, key)
		if err != nil {
			utils.JSONError(w, err.Error(), 400)
		} else {
			str := fmt.Sprintf("Successfully updated project: %v", key)
			res := serverRes{Res: str}
			json.NewEncoder(w).Encode(res)
		}
	}
}

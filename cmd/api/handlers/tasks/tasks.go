package tasks

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
	- Fetching tasks for only the id of the current project
	- Filter queries with uid and project id
*/

// Task - model for task entries
type Task struct {
	ID            string `json: id`
	ProjectID     int    `json: project_id`
	TaskName      string `json: task_name`
	TaskDesc      string `json: task_desc`
	TaskCompleted bool   `json: task_completed`
	TaskTime      string `json: task_time`
}

type serverRes struct {
	Res string `json: res`
}

// GetTasks - Get tasks for specific project
func GetTasks(app *application.Application) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		//w.Header().Set("Content-Type", "application/json")
		app.Logger.Info("GetTasks endpoint hit: /api/tasks")

		var tasks []Task
		sqlQuery := `SELECT * FROM tasks;`
		rows, err := app.Db.Client.Query(sqlQuery)
		defer rows.Close()

		for rows.Next() {
			var t Task
			err = rows.Scan(&t.ID, &t.ProjectID, &t.TaskName, &t.TaskDesc, &t.TaskTime, &t.TaskCompleted)
			if err != nil {
				//jsonError(w, err.Error(), 400)
				utils.JSONError(w, err.Error(), 400)
			}

			tasks = append(tasks, t)
		}

		json.NewEncoder(w).Encode(tasks)
	}
}

// CreateTask - Create new tasks and save to database
func CreateTask(app *application.Application) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		app.Logger.Info("CreateTask endpoint hit: /api/tasks")
		w.Header().Set("Content-Type", "application/json")

		var id int
		var t Task
		err := json.NewDecoder(r.Body).Decode(&t)
		if err != nil {
			utils.JSONError(w, "Bad Request", 400)
			return
		}
		defer r.Body.Close()

		// typo on task_desc
		sqlQuery := `
		INSERT INTO tasks (p_id, task_name, task_desc, task_time, task_completed)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id`

		err = app.Db.Client.QueryRow(sqlQuery, t.ProjectID, t.TaskName, t.TaskDesc, t.TaskTime, t.TaskCompleted).Scan(&id)
		if err != nil {
			utils.JSONError(w, err.Error(), 500)
		} else {
			str := fmt.Sprintf("Successfully created task: %v", id)
			res := serverRes{Res: str}
			json.NewEncoder(w).Encode(res)
		}
	}
}

// DeleteTask - Delete task with passed Id param
func DeleteTask(app *application.Application) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		app.Logger.Info("DeleteProject endpoint hit: /api/projects/:id")
		w.Header().Set("Content-Type", "application/json")

		// Get all variable/parameters passed returns a map
		vars := mux.Vars(r)
		key := vars["id"]

		sqlQuery := `DELETE FROM tasks WHERE id=$1`
		_, err := app.Db.Client.Query(sqlQuery, key)
		if err != nil {
			utils.JSONError(w, err.Error(), 500)
			return
		}

		str := fmt.Sprintf("Successfully deleted task: %v", key)
		res := serverRes{Res: str}
		json.NewEncoder(w).Encode(res)
	}
}

// UpdateTask - Update task with passed Id param
func UpdateTask(app *application.Application) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		app.Logger.Info("UpdateProject endpoint hit: /api/projects/:id")
		w.Header().Set("Content-Type", "application/json")

		var t Task
		// Get all variable/parameters passed returns a map
		vars := mux.Vars(r)
		key := vars["id"]

		err := json.NewDecoder(r.Body).Decode(&t)
		if err != nil {
			utils.JSONError(w, "Bad Request", 400)
			return
		}
		defer r.Body.Close()

		sqlQuery := `UPDATE tasks
		SET task_name = $1, task_desc = $2, task_time = $3, task_completed = $4
		WHERE id = $5`

		_, err = app.Db.Client.Query(sqlQuery, t.TaskName, t.TaskDesc, t.TaskTime, t.TaskCompleted, key)
		if err != nil {
			utils.JSONError(w, err.Error(), 500)
		} else {
			str := fmt.Sprintf("Successfully updated project: %v", key)
			res := serverRes{Res: str}
			json.NewEncoder(w).Encode(res)
		}
	}
}

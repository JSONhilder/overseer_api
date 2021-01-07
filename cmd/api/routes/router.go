package router

import (
	"log"
	"net/http"
	"os"

	// Apllication dependancy
	"github.com/JSONhilder/overseer_api/internal/application"
	"github.com/JSONhilder/overseer_api/internal/middleware"

	// Handlers
	"github.com/JSONhilder/overseer_api/cmd/api/handlers/projects"
	"github.com/JSONhilder/overseer_api/cmd/api/handlers/server"
	"github.com/JSONhilder/overseer_api/cmd/api/handlers/tasks"

	// import libraries
	"github.com/gorilla/mux"
)

// StartRouter - initials the routers with their respective handlers
func StartRouter(app *application.Application) {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", server.Check(app))

	authRoutes := router.PathPrefix("/api").Subrouter()
	authRoutes.Use(middleware.VerifyJwt)
	authRoutes.HandleFunc("/projects", projects.CreateProject(app)).Methods("POST")
	authRoutes.HandleFunc("/projects/{id}", projects.UpdateProject(app)).Methods("PUT")
	authRoutes.HandleFunc("/projects/{id}", projects.DeleteProject(app)).Methods("DELETE")
	authRoutes.HandleFunc("/projects/{id}", projects.GetProject(app))
	authRoutes.HandleFunc("/projects", projects.GetProjects(app))

	authRoutes.HandleFunc("/tasks", tasks.CreateTask(app)).Methods("POST")
	authRoutes.HandleFunc("/tasks/{id}", tasks.DeleteTask(app)).Methods("DELETE")
	authRoutes.HandleFunc("/tasks/{id}", tasks.UpdateTask(app)).Methods("PUT")
	authRoutes.HandleFunc("/tasks", tasks.GetTasks(app))

	log.Fatal(http.ListenAndServe(os.Getenv("SERVER_PORT"), router))
}

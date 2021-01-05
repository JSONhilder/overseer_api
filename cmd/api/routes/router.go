package router

import (
	"log"
	"net/http"
	"os"

	// Apllication dependancy
	"github.com/JSONhilder/overseer_api/internal/application"
	"github.com/JSONhilder/overseer_api/internal/middleware/authjwt"

	// Handlers
	"github.com/JSONhilder/overseer_api/cmd/api/handlers/projects"
	"github.com/JSONhilder/overseer_api/cmd/api/handlers/server"

	// import libraries
	"github.com/gorilla/mux"
)

// StartRouter - initials the routers with their respective handlers
func StartRouter(app *application.Application) {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", server.Check(app))

	authRoutes := router.PathPrefix("/api").Subrouter()
	authRoutes.Use(authjwt.Verify)
	authRoutes.HandleFunc("/projects", projects.CreateProject(app)).Methods("POST")
	authRoutes.HandleFunc("/projects/{id}", projects.GetProject(app)).Methods("GET")
	authRoutes.HandleFunc("/projects/{id}", projects.UpdateProject(app)).Methods("PUT")
	authRoutes.HandleFunc("/projects/{id}", projects.DeleteProject(app)).Methods("DELETE")
	authRoutes.HandleFunc("/projects", projects.GetProjects(app))

	log.Fatal(http.ListenAndServe(os.Getenv("SERVER_PORT"), router))
}

package router

import (
	"log"
	"net/http"
	"os"

	// Apllication dependancy
	"github.com/JSONhilder/overseer_api/internal/application"

	// Handlers
	"github.com/JSONhilder/overseer_api/cmd/api/handlers/projects"
	"github.com/JSONhilder/overseer_api/cmd/api/handlers/server"

	// import libraries
	"github.com/gorilla/mux"
)

// StartRouter - initials the routers with their respective handlers
func StartRouter(app *application.Application) {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", server.Check(app))
	myRouter.HandleFunc("/api/projects", projects.CreateProject(app)).Methods("POST")
	myRouter.HandleFunc("/api/projects/{id}", projects.GetProject(app)).Methods("GET")
	myRouter.HandleFunc("/api/projects", projects.GetProjects(app))

	log.Fatal(http.ListenAndServe(os.Getenv("SERVER_PORT"), myRouter))
}

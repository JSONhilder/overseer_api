package main

import (
	"log"

	"github.com/JSONhilder/overseer_api/internal/application"

	// import my directories
	router "github.com/JSONhilder/overseer_api/cmd/api/routes"
	"github.com/joho/godotenv"
)

/*
	load env file -> create application instance -> pass application instance to server
*/
func main() {
	// load env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Has instance of db and config in "app"
	app, err := application.Get()
	if err != nil {
		log.Fatal(err.Error())
	}

	app.LOG.Info("API running on port http:/localhost:8080")
	router.StartRouter(app)
}

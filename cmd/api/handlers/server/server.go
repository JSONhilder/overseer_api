package server

import (
	"fmt"
	"net/http"

	"github.com/JSONhilder/overseer_api/internal/application"
)

// Check - current status of server
func Check(app *application.Application) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "status: up")
		app.LOG.Info("Check endpoint hit: /")
	}
}

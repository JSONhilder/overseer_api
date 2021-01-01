package router

import (
	"fmt"
	"log"
	"net/http"

	// import libraries
	"github.com/gorilla/mux"
)

func healthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Server is running.")
}

// StartRouter - initials the routers with their respective handlers
func StartRouter() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", healthCheck)

	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

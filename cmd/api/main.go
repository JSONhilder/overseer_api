package main

import (
	"fmt"
	"log"

	// import my directories
	router "github.com/JSONhilder/overseer_api/cmd/api/routes"
	"github.com/JSONhilder/overseer_api/internal/db"
)

func main() {
	fmt.Println("Main go file.")

	db.ConnectDB()

	log.Println("API running on port http:/localhost:8080")
	router.StartRouter()
}

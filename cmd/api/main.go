package main

import (
	"fmt"

	// import my directories
	"github.com/JSONhilder/overseer_api/internal/db"
)

func main() {
	fmt.Println("Main go file.")

	db.ConnectDB()
}

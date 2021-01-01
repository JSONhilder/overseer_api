package db

import (
	"database/sql"
	"fmt"

	// postgres sql driver
	_ "github.com/lib/pq"
)

const (
	// @todo - move to env file
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "overseer_main"
)

// ConnectDB - Create db connection function
func ConnectDB() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// opening connection to db
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}

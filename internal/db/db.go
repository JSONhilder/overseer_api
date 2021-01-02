package db

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/JSONhilder/overseer_api/internal/config"

	// postgres sql driver
	_ "github.com/lib/pq"
)

// DB - DB struct to hold instance of db to pass around app
type DB struct {
	Client *sql.DB
}

// Get - get db instance
func Get(cfg *config.Config) (*DB, error) {
	// opening connection to db
	db, err := initDB(cfg)
	if err != nil {
		return nil, err
	}

	//defer db.Close()

	return &DB{
		Client: db,
	}, nil
}

func initDB(cfg *config.Config) (*sql.DB, error) {
	psqlInfo := config.GetDBConnString(cfg)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	fmt.Println("Successfully connected!")

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	setupTables(db)

	return db, nil
}

// Close - Helper function to close instance of db
func (d *DB) Close() error {
	return d.Client.Close()
}

// Setup tables
func setupTables(db *sql.DB) {

	query, err := ioutil.ReadFile("/home/jason/dev/github/overseer_api/internal/db/setup_db.sql")
	if err != nil {
		panic(err)
	}

	requests := strings.Split(string(query), ";")

	for _, request := range requests {
		_, err := db.Exec(request)

		if err != nil {
			panic(err)
		}
	}
}

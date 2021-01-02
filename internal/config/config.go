package config

import (
	"fmt"
	"os"
)

// Config struct for db comnection string
type Config struct {
	dbUser string
	dbPswd string
	dbHost string
	dbPort string
	dbName string
}

// Get - Return a pointer to a Config struct
func Get() *Config {
	conf := &Config{}

	conf.dbUser = os.Getenv("POSTGRES_USER")
	conf.dbPswd = os.Getenv("POSTGRES_PASSWORD")
	conf.dbHost = os.Getenv("POSTGRES_HOST")
	conf.dbPort = os.Getenv("POSTGRES_PORT")
	conf.dbName = os.Getenv("POSTGRES_DB")

	return conf
}

// GetDBConnString - Builds and returns connection string from .env file
func GetDBConnString(c *Config) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		c.dbUser,
		c.dbPswd,
		c.dbHost,
		c.dbPort,
		c.dbName,
	)
}

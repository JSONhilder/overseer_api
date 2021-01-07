package application

import (
	"github.com/JSONhilder/overseer_api/internal/config"
	"github.com/JSONhilder/overseer_api/internal/db"
	"go.uber.org/zap"
)

// Application - struct with pointers to db instance and config object
type Application struct {
	Db     *db.DB
	Conf   *config.Config
	Logger *zap.SugaredLogger
}

// Get - Return an application state with access to db, config and logger for dependancy injection
func Get() (*Application, error) {
	cfg := config.Get()

	// Init the logger and add to apllication struct
	logger, _ := zap.NewDevelopment()
	defer logger.Sync() // flushes buffer, if any
	log := logger.Sugar()

	db, err := db.Get(cfg)
	if err != nil {
		return nil, err
	}
	log.Info("Database successfully connected.")

	return &Application{
		Db:     db,
		Conf:   cfg,
		Logger: log,
	}, nil
}

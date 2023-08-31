package db

import (
	"english/backend/config"
	"english/backend/shared/errorlog"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Handler struct {
	DB *sqlx.DB
}

func Init(cfg *config.SVCConfig) (Handler, error) {
	var err error

	if cfg.DBConfig, err = config.NewDBConfig(); err != nil {
		return Handler{}, errorlog.NewError(err.Error())
	}

	cfg.DBConfig.SetDataBaseURL(cfg.IsDocker)

	db, err := sqlx.Connect("postgres", cfg.DBConfig.DataBaseURL)
	if err != nil {
		return Handler{}, errorlog.NewError(err.Error())
	}

	return Handler{db}, nil
}

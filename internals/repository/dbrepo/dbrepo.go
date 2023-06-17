package dbrepo

import (
	"database/sql"

	"github.com/ishanshre/goFiberApiExample/internals/config"
	"github.com/ishanshre/goFiberApiExample/internals/repository"
)

type postgresDBRepo struct {
	Global *config.AppConfig
	DB     *sql.DB
}

func NewPostgresRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
	return &postgresDBRepo{
		Global: a,
		DB:     conn,
	}
}

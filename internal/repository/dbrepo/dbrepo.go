package dbrepo

import (
	"database/sql"
	"github.com/esmaeilmirzaee/bookings/internal/config"
	"github.com/esmaeilmirzaee/bookings/internal/repository"
)

type postgresDBRepo struct {
	App *config.AppConfig
	DB *sql.DB
}

// NewPostgresRepo returns a repository database
func NewPostgresRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
	return &postgresDBRepo{
		App: a,
		DB: conn,
	}
}
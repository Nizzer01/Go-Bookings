package dbrepo

import (
	"database/sql"
	"github.com/Nizzer01/Go-Bookings/internal/config"
	"github.com/Nizzer01/Go-Bookings/internal/repository"
)

//Setup Database types eg PG, SQL, Mongo
type postgresDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

type testDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

func NewPostgresRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
	return &postgresDBRepo{
		App: a,
		DB:  conn,
	}
}

func NewTestingsRepo(a *config.AppConfig) repository.DatabaseRepo {
	return &testDBRepo{
		App: a,
	}
}

package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Postgres struct {
	*sql.DB
}

func NewPostgres(host, port, user, password, name, sslmode string, maxIdleCons, maxOpenCons int) (pg *Postgres, err error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, name, sslmode)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return
	}

	db.SetMaxIdleConns(maxIdleCons)
	db.SetMaxOpenConns(maxOpenCons)

	err = db.Ping()
	if err != nil {
		return
	}
	pg = &Postgres{db}
	return
}

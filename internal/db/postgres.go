package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type PostgresOptions struct {
	Host     string
	Post     string
	User     string
	Password string
	Name     string
}

func NewPostgresDB(opt PostgresOptions) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		opt.User, opt.Password, opt.Host, opt.Post, opt.Name))

	if err != nil {
		return nil, err
	}

	return db, nil
}

package app

import (
	"ananhartanto/lion-backend/helper"
	"database/sql"
	"time"
)

func NewDB() *sql.DB {
	db, err := sql.Open("postgres", "postgres://<USERNAME>:<PASSWORD>@<HOST>:<PORT>/<DB_NAME>")
	helper.Error(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
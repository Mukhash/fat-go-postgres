package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "qwerty"
	dbname   = "postgres"
)

var conn *sql.DB

func CreateConnection() error {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}
	if err = db.Ping(); err != nil {
		return err
	}
	conn = db
	return nil
}

func GetDB() *sql.DB {
	if conn != nil {
		return conn
	}
	return nil
}

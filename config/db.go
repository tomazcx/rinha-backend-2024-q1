package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var dbConn *sql.DB

func ConnectToDb(conf *AppConfig) (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", conf.DBHost, conf.DBUser, conf.DBPassword, conf.DBName)
	db, _ := sql.Open("postgres", connStr)
	err := db.Ping()

	if err != nil {
		return db, err
	}

	db.SetMaxOpenConns(30)
	db.SetMaxIdleConns(30)

	dbConn = db

	return dbConn, nil
}

func GetDBConn() *sql.DB {
	return dbConn
}

package mysql

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var connection *sql.DB

func Query(query string, args ...interface{}) (*sql.Rows, error) {
	db := getConnection()
	return db.Query(query, args...)
}

func QueryRow(query string, args ...interface{}) *sql.Row {
	db := getConnection()
	return db.QueryRow(query, args...)
}

func getConnection() *sql.DB {
	if connection == nil {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
		connection, _ = sql.Open("mysql", dsn)
	}

	return connection
}

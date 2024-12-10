package db_sql_pkg

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() (*sql.DB, error) {
	connStr := "host=localhost port=5432 user=student password=securepass dbname=student sslmode=disable"
	DB, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("DB Connection error:", err)
	}

	if err := DB.Ping(); err != nil {
		return nil, fmt.Errorf("DB Connection error: %w", err)
	} else {
		return DB, nil
	}
}

func RunTask3() {
	if _, err := ConnectDB(); err == nil {
		fmt.Println("Connection Successful")
	}
}

package sql_intro

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
	connStr := "host=localhost port=5432 user=student password=securepass dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}
	DB = db
}

func RunTask2() {
  ConnectDB()
  query,err:= DB.Query("CREATE TABLE students (id SERIAL PRIMARY KEY, name TEXT, age INT)")
  if err != nil{
    fmt.Println(err)
  
  ry.Columns
}

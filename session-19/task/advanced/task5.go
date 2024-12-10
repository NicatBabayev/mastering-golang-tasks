package advanced

import (
	"fmt"
	"session-19/task/db_sql_pkg"
)

func preparedStmt() {
	DB, err := db_sql_pkg.ConnectDB()
	if err != nil {
		fmt.Println(err)
	}

	stmt, err := DB.Prepare("INSERT INTO students (name, age) VALUES ($1, $2)")
	if err != nil {
		fmt.Println(err)
	}
	res, err := stmt.Exec("Koroglu", 7777)
	if err != nil {
		fmt.Println(err)
	}
	affected, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err)
	}
	if affected > 0 {
		fmt.Println("Insert successful")
	}
}

func RunTask5() {
	preparedStmt()
}

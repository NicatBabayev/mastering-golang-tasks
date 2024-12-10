package advanced

import (
	"fmt"
	"session-19/task/db_sql_pkg"
)

func dbTx() {
	DB, err := db_sql_pkg.ConnectDB()
	if err != nil {
		fmt.Println(err)
	}
	tx, err := DB.Begin()
	if err != nil {
		fmt.Println(err)
	}
	res, err := tx.Exec("UPDATE students SET age=21 WHERE name='Ali'")
	if err != nil {
		fmt.Println(err)
	}
	if affected, err := res.RowsAffected(); err == nil && affected == 1 {
		fmt.Println("Transaction successful")
	}
	err = tx.Commit()
	if err != nil {
		fmt.Println(err)
	}
}

func RunTask6() {
	dbTx()
}

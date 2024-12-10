package db_sql_pkg

import "fmt"

type student struct {
	ID   int
	name string
	age  int
}

func queryDB() {
	student := student{}
	DB, err := ConnectDB()
	if err != nil {
		fmt.Println(err)
		return
	}
	rows, err := DB.Query("SELECT * FROM students")
	if err != nil {
		fmt.Println(err)
		return
	}
	for rows.Next() {
		err = rows.Scan(&student.ID, &student.name, &student.age)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("ID: %d, Name: %s, Age: %d\n", student.ID, student.name, student.age)
	}
}

func RunTask4() {
	queryDB()

}

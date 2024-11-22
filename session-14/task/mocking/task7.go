package mocking

import "fmt"

type User struct {
	ID   int
	Name string
	Age  int
}

type DB struct {
	users []User
}

type Database interface {
	GetUserByID(id int) (User, error)
}

func (db *DB) GetUserByID(id int) (User, error) {
	for _, user := range db.users {
		if user.ID == id {
			return user, nil
		}
	}
	return User{}, fmt.Errorf("User not found.")
}

func GetUser(db Database, id int) (User, error) {
	user, err := db.GetUserByID(id)
	if err != nil {
		return User{}, fmt.Errorf("getUser error: %w", err)
	}
	return user, nil

}

func RunTask7() {
	db := DB{
		users: []User{
			{
				ID:   1,
				Name: "Koroglu",
				Age:  7777,
			},
		},
	}

	user, err := GetUser(&db, 1)
	if err != nil {
		fmt.Println(fmt.Errorf("function error: %w", err))
		return
	}

	fmt.Println("user found: ", user.Name)
}

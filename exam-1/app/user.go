package app

import "fmt"

type User struct {
	UserID        int
	Name          string
	BorrowedBooks []int
	BorrowLimit   int //limit = 3
}

func listAllUsers(db *DB) {
	db.listAllUsers()
}
func selectUser(db *DB) int {
	var selectionID int
	fmt.Println("All users: ")
	for i, user := range db.users {
		fmt.Printf("- [%d] %s.\n", i, user.Name)
	}
	fmt.Print("Select the user: ")
	fmt.Scanln(&selectionID)
	res, _ := db.searchUser(db.users[selectionID])

	return res
}

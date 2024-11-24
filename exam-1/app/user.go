package app

type User struct {
	UserID        int
	Name          string
	BorrowedBooks []int
	BorrowLimit   int
}

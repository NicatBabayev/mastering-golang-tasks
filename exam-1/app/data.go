package app

func LoadInitialData(db *DB) {
	books := []Book{
		{
			BookID: 1,
			Title:  "The Go Programming Guide",
			Author: "John Doe",
			Status: "available",
		},
		{
			BookID: 2,
			Title:  "Concurrency in Practice",
			Author: "Jane Smith",
			Status: "available",
		},
		{
			BookID: 3,
			Title:  "Microservices Design",
			Author: "Alan Turing",
			Status: "available",
		},
		{
			BookID: 4,
			Title:  "Data Structures",
			Author: "Grace Hopper",
			Status: "available",
		},
		{
			BookID: 5,
			Title:  "The Art of Computer Prog",
			Author: "Donald Knuth",
			Status: "available",
		},
		{
			BookID: 6,
			Title:  "Clean Code",
			Author: "Robert Martin",
			Status: "available",
		},
		{
			BookID: 7,
			Title:  "Algorithms Unlocked",
			Author: "Thomas Gormen",
			Status: "available",
		},
		{
			BookID: 8,
			Title:  "Software Architescture",
			Author: "Martin Fowler",
			Status: "available",
		},
		{
			BookID: 9,
			Title:  "System Design",
			Author: "Scott Meyers",
			Status: "available",
		},
		{
			BookID: 10,
			Title:  "Effective Go",
			Author: "Rob Pike",
			Status: "available",
		},
	}
	users := []User{
		{
			UserID:        1,
			Name:          "Ali",
			BorrowedBooks: []int{},
			BorrowLimit:   0,
		},
		{
			UserID:        2,
			Name:          "Babek",
			BorrowedBooks: []int{},
			BorrowLimit:   0,
		},
		{
			UserID:        3,
			Name:          "Cabir",
			BorrowedBooks: []int{},
			BorrowLimit:   0,
		},
		{
			UserID:        4,
			Name:          "Diana",
			BorrowedBooks: []int{},
			BorrowLimit:   0,
		},
		{
			UserID:        5,
			Name:          "Elvira",
			BorrowedBooks: []int{},
			BorrowLimit:   0,
		},
	}
	for _, book := range books {
		db.AddBook(book)
	}
	for _, user := range users {
		db.AddUser(user)
	}
}

func DisplayData(db *DB) {
}

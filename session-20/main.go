package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

type Book struct {
	ID            uint   `gorm:"primaryKey"`
	Title         string `gorm:"size:255;not null,unique"`
	Author        string `gorm:"size:100"`
	PublishedYear uint
	CreatedAt     time.Time
}
type User struct {
	ID      uint
	Name    string `gorm:"size:255;unique"`
	Profile Profile
}
type Profile struct {
	ID     uint
	Bio    string
	UserID uint
}
type Customer struct {
	gorm.Model
	Name   string
	Orders []Order
}
type Order struct {
	gorm.Model
	CustomerID uint
	ItemName   string
}
type Post struct {
	gorm.Model
	PostOrder string
	PUser     PUser
	PUserID   uint
}
type PUser struct {
	gorm.Model
	Author string
}

type GUser struct {
	gorm.Model
	UserName string
	Groups   []Group `gorm:"many2many:gusers_groups"`
}

type Group struct {
	gorm.Model
	GroupName string
	GUsers    []GUser `gorm:"many2many:gusers_groups"`
}

func main() {
	// Task 1
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Baku"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("DB Connection error:", err)
		return
	}

	dbPing, _ := db.DB()
	if err = dbPing.Ping(); err == nil {
		fmt.Println("Connected to PostgreSQL database successfully!")
	}

	//	Task 2
	err = db.AutoMigrate(&Book{})
	if err != nil {
		fmt.Println("Migration error")
		return
	} else {
		fmt.Println("Database migration completed successfully!")
	}
	// Task 3
	book := Book{
		Title:         "Go Programming",
		Author:        "John Doe",
		PublishedYear: 2023,
	}
	db.Create(&book)
	fmt.Println("Book record inserted successfully!")
	var books []Book
	err = db.Find(&books).Error
	if err != nil {
		fmt.Println("No books found")
	}
	fmt.Println("Books in database:")
	for _, book := range books {
		fmt.Printf("ID: %d, Title: %s, Author: %s, Year: %d\n", book.ID, book.Title, book.Author, book.PublishedYear)
	}
	// Associations
	// HasOne
	err = db.AutoMigrate(&User{}, &Profile{})
	if err != nil {
		return
	}
	user := User{
		Name: "John Doe",
		Profile: Profile{
			Bio: "Profile",
		},
	}
	var users []User
	db.Create(&user)
	err = db.Model(&User{}).Preload("Profile").Find(&users).Error

	for _, user := range users {
		fmt.Printf("User: %s, Bio: %s\n", user.Name, user.Profile.Bio)
	}
	//	HasMany
	err = db.AutoMigrate(&Customer{}, &Order{})
	if err != nil {
		return
	}
	customer := Customer{
		Name:   "Ali",
		Orders: []Order{{ItemName: "aa"}, {ItemName: "bb"}},
	}
	db.Create(&customer)
	var customers []Customer
	err = db.Model(&Customer{}).Preload("Orders").Find(&customers).Error
	for _, customer := range customers {
		var s string
		for _, order := range customer.Orders {
			s += fmt.Sprintf("Order %s ", order.ItemName)
		}
		fmt.Printf("Customer: %s %s\n", customer.Name, s)
	}
	// BelongsTo
	err = db.AutoMigrate(&Post{}, &PUser{})
	if err != nil {
		return
	}
	post := Post{
		PostOrder: "First",
		PUser:     PUser{Author: "Ali"},
	}
	db.Create(&post)
	var posts []Post
	err = db.Model(&Post{}).Preload("PUser").Find(&posts).Error

	for _, post := range posts {
		fmt.Printf("Post: %s Post, Author: %s\n", post.PostOrder, post.PUser.Author)
	}
	//	ManyToMany
	err = db.AutoMigrate(&GUser{}, &Group{})
	if err != nil {
		return
	}
	gUser := GUser{
		UserName: "Ali",
		Groups:   []Group{{GroupName: "Admin"}, {GroupName: "Developer"}},
	}
	db.Create(&gUser)
	var gUsers []GUser
	err = db.Model(&GUser{}).Preload("Groups").Find(&gUsers).Error
	for _, gUser := range gUsers {
		var s string
		for _, group := range gUser.Groups {
			s += fmt.Sprintf("Group: %s ", group.GroupName)
		}
		fmt.Printf("User: %s %s\n", gUser.UserName, s)
	}
	//	Transactions
	err = db.Transaction(func(tx *gorm.DB) error {
		if userErr1 := db.Create(&User{Name: "Ali1"}).Error; userErr1 != nil {
			return userErr1
		}
		if userErr2 := db.Create(&User{Name: "Ali2"}).Error; userErr2 != nil {
			return userErr2
		}
		if userErr3 := db.Create(&User{ID: 18, Name: "Ali2"}).Error; userErr3 != nil {
			return userErr3
		}

		return nil
	})
	if err != nil {
		fmt.Println("Transaction rolled back")
	} else {
		fmt.Println("Transaction committed successfully")
	}
}

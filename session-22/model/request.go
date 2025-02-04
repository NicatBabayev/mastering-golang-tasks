package model

import (
	"time"
)

type BookRequest struct {
	ID            uint      `json:"id"`
	Title         string    `json:"title" gorm:"size:255,unique"`
	Author        string    `json:"author" gorm:"size:255"`
	Price         float64   `json:"price" gorm:"decimal(10,2)"`
	PublishedYear time.Time `json:"published_year"`
}

type LoginRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

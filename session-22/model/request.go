package model

import (
	"time"
)

type BookRequest struct {
	ID            uint
	Title         string  `gorm:"size:255,unique"`
	Author        string  `gorm:"size:255"`
	Price         float64 `gorm:"decimal(10,2)"`
	PublishedYear time.Time
}

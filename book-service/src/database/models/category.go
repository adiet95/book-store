package models

import "time"

type Category struct {
	CategoryId   string    `gorm:"primaryKey;" json:"id,omitempty"`
	CategoryName string    `json:"category_name,omitempty"`
	Description  string    `json:"description,omitempty"`
	CreatedAt    time.Time `gorm:"default:now(); not null" json:"created_at"`
	UpdatedAt    time.Time `gorm:"default:now(); not null" json:"updated_at"`
}

type Categories []Category

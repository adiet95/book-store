package models

import "time"

type Author struct {
	AuthorId  uint      `gorm:"primaryKey" json:"id,omitempty"`
	FullName  string    `json:"full_name,omitempty"`
	Country   string    `json:"country,omitempty"`
	CreatedAt time.Time `gorm:"default:now(); not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:now(); not null" json:"updated_at"`
}

type Authors []Author

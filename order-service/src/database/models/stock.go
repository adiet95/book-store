package models

import "time"

type Stock struct {
	StockId   uint      `gorm:"primaryKey" json:"stock_id"`
	StockName string    `json:"stock_name"`
	BookId    uint      `json:"book_id"`
	Book      Book      `json:"order"`
	Qty       int       `json:"qty"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Stocks []Stock

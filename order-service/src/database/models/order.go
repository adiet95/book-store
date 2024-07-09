package models

import "time"

type Order struct {
	OrderId          uint      `gorm:"primaryKey" json:"id,omitempty"`
	OrderName        string    `json:"order_name"`
	OrderDescription string    `json:"order_description"`
	UserId           int       `json:"user_id"`
	User             User      `json:"user"`
	StockId          int       `json:"stock_id"`
	Stock            Stock     `json:"stock"`
	Qty              int       `json:"qty"`
	Status           string    `json:"status"`
	StartDate        time.Time `json:"start_date"`
	EndDate          time.Time `json:"end_date"`
}

type Orders []Order

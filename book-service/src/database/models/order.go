package models

type Order struct {
	OrderId    uint     `gorm:"primaryKey" json:"id,omitempty"`
	OrderName  string   `json:"order_name"`
	UserId     string   `json:"user_id"`
	User       User     `json:"user"`
	CategoryId string   `json:"category_id"`
	Category   Category `json:"category"`
	AuthorId   string   `json:"author_id"`
	Author     Author   `json:"author"`
	Qty        string   `json:"qty"`
	Status     string   `json:"status"`
}

type Orders []Order

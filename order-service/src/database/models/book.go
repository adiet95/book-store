package models

type Book struct {
	BookId      uint     `gorm:"primaryKey" json:"id,omitempty"`
	BookName    string   `json:"order_name"`
	CategoryId  uint     `json:"category_id"`
	Category    Category `json:"category"`
	AuthorId    uint     `json:"author_id"`
	Author      Author   `json:"author"`
	Qty         int      `json:"qty"`
	Status      string   `json:"status"`
	QtyOrder    int      `json:"qty_order"`
	StatusOrder string   `json:"status_order"`
}

type Books []Book

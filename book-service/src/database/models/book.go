package models

type Book struct {
	BookId          uint     `gorm:"primaryKey" json:"id,omitempty"`
	BookName        string   `json:"book_name"`
	BookDescription string   `json:"book_description"`
	CategoryId      int      `json:"category_id"`
	Category        Category `json:"category"`
	AuthorId        int      `json:"author_id"`
	Author          Author   `json:"author"`
}

type Books []Book

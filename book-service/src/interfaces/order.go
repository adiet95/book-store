package interfaces

import (
	"github.com/adiet95/book-store/book-service/src/database/models"
	"github.com/adiet95/book-store/book-service/src/libs"
)

type BookRepo interface {
	FindAll(limit, offset int) (*models.Books, error)
	Save(data *models.Book) (*models.Book, error)
	Update(data *models.Book, id int) (*models.Book, error)
	Delete(id int) (*models.Book, error)
	FindByName(name string) (*models.Books, error)
	GetUserId(email string) (*models.User, error)
	FindById(id int) (*models.Books, error)
}

type BookService interface {
	GetAll(limit, offset int) *libs.Response
	Add(data *models.Book) *libs.Response
	Update(data *models.Book, id int) *libs.Response
	Delete(id int) *libs.Response
	Search(name string) *libs.Response
	SearchId(id int) *libs.Response
}

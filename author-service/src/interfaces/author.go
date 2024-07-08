package interfaces

import (
	"github.com/adiet95/book-store/author-service/src/database/models"
	"github.com/adiet95/book-store/author-service/src/libs"
)

type AuthorRepo interface {
	FindAll(limit, offset int) (*models.Authors, error)
	Save(data *models.Author) (*models.Author, error)
	Update(data *models.Author, id int) (*models.Author, error)
	Delete(id int) (*models.Author, error)
	FindByName(name string) (*models.Authors, error)
	FindById(id int) (*models.Authors, error)
}

type AuthorService interface {
	GetAll(limit, offset int) *libs.Response
	Add(data *models.Author) *libs.Response
	Update(data *models.Author, id int) *libs.Response
	Delete(id int) *libs.Response
	Search(name string) *libs.Response
	SearchId(id int) *libs.Response
}

package interfaces

import (
	"github.com/adiet95/book-store/category-service/src/database/models"
	"github.com/adiet95/book-store/category-service/src/libs"
)

type CategoryRepo interface {
	FindAll(limit, offset int) (*models.Categories, error)
	Save(data *models.Category) (*models.Category, error)
	Update(data *models.Category, id int) (*models.Category, error)
	Delete(id int) (*models.Category, error)
	FindByName(name string) (*models.Categories, error)
	FindById(id int) (*models.Category, error)
}

type CategoryService interface {
	GetAll(limit, offset int) *libs.Response
	Add(data *models.Category) *libs.Response
	Update(data *models.Category, id int) *libs.Response
	Delete(id int) *libs.Response
	Search(name string) *libs.Response
	SearchId(id int) *libs.Response
}

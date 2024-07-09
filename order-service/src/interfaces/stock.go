package interfaces

import (
	"github.com/adiet95/book-store/order-service/src/database/models"
	"github.com/adiet95/book-store/order-service/src/libs"
)

type StockRepo interface {
	FindAll(limit, offset int) (*models.Stocks, error)
	Save(data *models.Stock) (*models.Stock, error)
	Update(data *models.Stock, id int) (*models.Stock, error)
	Delete(id int) (*models.Stock, error)
	FindByName(name string) (*models.Stocks, error)
	GetUserId(email string) (*models.User, error)
	FindById(id int) (*models.Stocks, error)
}

type StockService interface {
	GetAll(limit, offset int) *libs.Response
	Add(data *models.Stock) *libs.Response
	Update(data *models.Stock, id int) *libs.Response
	Delete(id int) *libs.Response
	Search(name string) *libs.Response
	SearchId(id int) *libs.Response
}

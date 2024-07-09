package stock

import (
	"github.com/adiet95/book-store/order-service/src/database/models"
	"github.com/adiet95/book-store/order-service/src/interfaces"
	"github.com/adiet95/book-store/order-service/src/libs"
)

type stock_service struct {
	stock_repo interfaces.StockRepo
}

func NewService(reps interfaces.StockRepo) *stock_service {
	return &stock_service{reps}
}

func (r *stock_service) GetAll(limit, offset int) *libs.Response {
	data, err := r.stock_repo.FindAll(limit, offset)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	return libs.New(data, 200, false)
}

func (re *stock_service) Add(data *models.Stock) *libs.Response {
	result, err := re.stock_repo.Save(data)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	return libs.New(result, 200, false)
}

func (re *stock_service) Update(data *models.Stock, id int) *libs.Response {
	res, err := re.stock_repo.Update(data, id)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	return libs.New(res, 200, false)
}

func (re *stock_service) Delete(id int) *libs.Response {
	data, err := re.stock_repo.Delete(id)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	return libs.New(data, 200, false)
}

func (re *stock_service) Search(name string) *libs.Response {
	data, err := re.stock_repo.FindByName(name)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	return libs.New(data, 200, false)
}

func (re *stock_service) SearchId(id int) *libs.Response {
	data, err := re.stock_repo.FindById(id)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	return libs.New(data, 200, false)
}

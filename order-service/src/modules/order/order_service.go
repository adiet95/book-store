package order

import (
	"github.com/adiet95/book-store/order-service/src/database/models"
	"github.com/adiet95/book-store/order-service/src/interfaces"
	"github.com/adiet95/book-store/order-service/src/libs"
	"strings"
)

type order_service struct {
	order_repo interfaces.OrderRepo
	stock_repo interfaces.StockRepo
}

func NewService(orderRepo interfaces.OrderRepo, stockRepo interfaces.StockRepo) *order_service {
	return &order_service{
		order_repo: orderRepo,
		stock_repo: stockRepo,
	}
}

func (r *order_service) GetAll(limit, offset int) *libs.Response {
	data, err := r.order_repo.FindAll(limit, offset)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	return libs.New(data, 200, false)
}

func (re *order_service) Add(data *models.Order) *libs.Response {
	result, err := re.order_repo.Save(data)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	return libs.New(result, 200, false)
}

func (re *order_service) Update(data *models.Order, id int) *libs.Response {
	res, err := re.order_repo.Update(data, id)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	return libs.New(res, 200, false)
}

func (re *order_service) Delete(id int) *libs.Response {
	data, err := re.order_repo.Delete(id)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	return libs.New(data, 200, false)
}

func (re *order_service) Search(name string) *libs.Response {
	data, err := re.order_repo.FindByName(name)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	return libs.New(data, 200, false)
}

func (re *order_service) SearchId(id int) *libs.Response {
	data, err := re.order_repo.FindById(id)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	return libs.New(data, 200, false)
}

func (re *order_service) SearchByUserId(limit, offset int, email string) *libs.Response {
	dataUser, err := re.order_repo.GetUserId(email)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}

	data, err := re.order_repo.FindByUserId(int(dataUser.UserId), limit, offset)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	return libs.New(data, 200, false)
}

func (re *order_service) Order(data *models.Order, email string) *libs.Response {
	dataUser, err := re.order_repo.GetUserId(email)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}

	data.UserId = int(dataUser.UserId)

	dataStock, err := re.stock_repo.FindById(data.StockId)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	if data.Status != "" {
		if data.Status == strings.ToLower("borrow") {
			dataStock.Qty = dataStock.Qty - data.Qty
			if data.Qty == 0 {
				dataStock.Status = "stock out"
			} else if dataStock.Qty > 0 {
				dataStock.Status = "available"
			} else if dataStock.Qty < 0 {
				return libs.New("out of stock", 400, true)
			}

			_, err = re.stock_repo.Update(dataStock, data.StockId)
			if err != nil {
				return libs.New(err.Error(), 400, true)
			}
		} else if data.Status == strings.ToLower("return") {
			dataStock.Qty = dataStock.Qty + data.Qty
			if data.Qty == 0 {
				dataStock.Status = "stock out"
			} else if dataStock.Qty > 0 {
				dataStock.Status = "available"
			} else if dataStock.Qty < 0 {
				return libs.New("out of stock", 400, true)
			}

			_, err = re.stock_repo.Update(dataStock, data.StockId)
			if err != nil {
				return libs.New(err.Error(), 400, true)
			}
		}
	}
	res, err := re.order_repo.Save(data)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	return libs.New(res, 200, false)
}

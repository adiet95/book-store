package stock

import (
	"errors"

	"github.com/adiet95/book-store/order-service/src/database/models"
	"gorm.io/gorm"
)

type stock_repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *stock_repo {
	return &stock_repo{db}
}

func (r *stock_repo) FindAll(limit, offset int) (*models.Stocks, error) {
	var datas *models.Stocks
	result := r.db.Model(&datas).Limit(limit).Offset(offset).Find(&datas)
	if result.Error != nil {
		return nil, errors.New("failed obtain datas")
	}
	return datas, nil
}

func (r *stock_repo) Save(data *models.Stock) (*models.Stock, error) {
	checkName := r.db.Where("LOWER(stock_name) LIKE ?", "%"+data.StockName+"%").Find(&data)
	if checkName.RowsAffected != 0 {
		return nil, errors.New("stock name is exist")
	}

	var dataBook *models.Book
	checkBook := r.db.Model(&dataBook).Where("book_id = ?", data.BookId).First(&dataBook)
	if checkBook.RowsAffected == 0 {
		return nil, errors.New("book not found")
	}

	res := r.db.Create(data)
	if res.Error != nil {
		return nil, errors.New("failed obtain datas")
	}
	return data, nil
}

func (re *stock_repo) Update(data *models.Stock, id int) (*models.Stock, error) {
	var dataBook *models.Book
	checkBook := re.db.Model(&dataBook).Where("book_id = ?", data.BookId).First(&dataBook)
	if checkBook.RowsAffected == 0 {
		return nil, errors.New("book not found")
	}

	res := re.db.Model(&data).Where("stock_id = ?", id).Updates(&data)

	if res.Error != nil {
		return nil, errors.New("failed to update data")
	}
	return data, nil
}

func (re *stock_repo) Delete(id int) (*models.Stock, error) {
	var data *models.Stock
	var datas *models.Stocks
	res := re.db.Where("stock_id = ?", id).Find(&datas)

	if res.RowsAffected == 0 {
		return nil, errors.New("data not found")
	}
	r := re.db.Model(data).Where("stock_id = ?", id).Delete(&data)
	if r.Error != nil {
		return nil, errors.New("failed to delete data")
	}
	return nil, nil
}

func (re *stock_repo) FindByName(name string) (*models.Stocks, error) {
	var datas *models.Stocks
	res := re.db.Where("LOWER(stock_name) LIKE ?", "%"+name+"%").Find(&datas)
	if res.Error != nil {
		return nil, errors.New("failed to found data")
	}
	if res.RowsAffected == 0 {
		return nil, errors.New("data not found")
	}
	return datas, nil
}

func (re *stock_repo) FindById(id int) (*models.Stock, error) {
	var datas *models.Stock
	res := re.db.Where("stock_id = ?", id).First(&datas)
	if res.Error != nil {
		return nil, errors.New("failed to found data")
	}
	if res.RowsAffected == 0 {
		return nil, errors.New("data not found")
	}
	return datas, nil
}

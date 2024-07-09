package category

import (
	"errors"

	"github.com/adiet95/book-store/category-service/src/database/models"
	"gorm.io/gorm"
)

type category_repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *category_repo {
	return &category_repo{db}
}

func (r *category_repo) FindAll(limit, offset int) (*models.Categories, error) {
	var datas *models.Categories
	result := r.db.Model(&datas).Limit(limit).Offset(offset).Find(&datas)
	if result.Error != nil {
		return nil, errors.New("failed obtain datas")
	}
	return datas, nil
}

func (r *category_repo) Save(data *models.Category) (*models.Category, error) {
	res := r.db.Create(data)
	if res.Error != nil {
		return nil, errors.New("failed obtain datas")
	}
	return data, nil
}

func (re *category_repo) Update(data *models.Category, id int) (*models.Category, error) {
	res := re.db.Model(&data).Where("category_id = ?", id).Updates(&data)

	if res.Error != nil {
		return nil, errors.New("failed to update data")
	}
	return data, nil
}

func (re *category_repo) Delete(id int) (*models.Category, error) {
	var data *models.Category
	var datas *models.Categories
	res := re.db.Where("category_id = ?", id).Find(&datas)

	if res.RowsAffected == 0 {
		return nil, errors.New("data not found")
	}
	r := re.db.Model(data).Where("category_id = ?", id).Delete(&data)
	if r.Error != nil {
		return nil, errors.New("failed to delete data")
	}
	return nil, nil
}

func (re *category_repo) FindByName(name string) (*models.Categories, error) {
	var datas *models.Categories
	res := re.db.Order("category_id asc").Where("LOWER(category_name) LIKE ?", "%"+name+"%").Find(&datas)
	if res.Error != nil {
		return nil, errors.New("failed to found data")
	}
	if res.RowsAffected == 0 {
		return nil, errors.New("data not found")
	}
	return datas, nil
}

func (re *category_repo) FindById(id int) (*models.Categories, error) {
	var datas *models.Categories
	res := re.db.Order("category_id asc").Where("category_id = ?", id).Find(&datas)
	if res.Error != nil {
		return nil, errors.New("failed to found data")
	}
	if res.RowsAffected == 0 {
		return nil, errors.New("data not found")
	}
	return datas, nil
}

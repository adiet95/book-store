package author

import (
	"errors"

	"github.com/adiet95/book-store/author-service/src/database/models"
	"gorm.io/gorm"
)

type author_repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *author_repo {
	return &author_repo{db}
}

func (r *author_repo) FindAll(limit, offset int) (*models.Authors, error) {
	var datas *models.Authors
	result := r.db.Model(&datas).Limit(limit).Offset(offset).Find(&datas)
	if result.Error != nil {
		return nil, errors.New("failed obtain datas")
	}
	return datas, nil
}

func (r *author_repo) Save(data *models.Author) (*models.Author, error) {
	res := r.db.Create(data)
	if res.Error != nil {
		return nil, errors.New("failed obtain datas")
	}
	return data, nil
}

func (re *author_repo) Update(data *models.Author, id int) (*models.Author, error) {
	res := re.db.Model(&data).Where("id = ?", id).Updates(&data)

	if res.Error != nil {
		return nil, errors.New("failed to update data")
	}
	return data, nil
}

func (re *author_repo) Delete(id int) (*models.Author, error) {
	var data *models.Author
	var datas *models.Authors
	res := re.db.Where("id = ?", id).Find(&datas)

	if res.RowsAffected == 0 {
		return nil, errors.New("data not found")
	}
	r := re.db.Model(data).Where("id = ?", id).Delete(&data)
	if r.Error != nil {
		return nil, errors.New("failed to delete data")
	}
	return nil, nil
}

func (re *author_repo) FindByName(name string) (*models.Authors, error) {
	var datas *models.Authors
	res := re.db.Order("id asc").Where("LOWER(order_name) LIKE ?", "%"+name+"%").Find(&datas)
	if res.Error != nil {
		return nil, errors.New("failed to found data")
	}
	if res.RowsAffected == 0 {
		return nil, errors.New("data not found")
	}
	return datas, nil
}

func (re *author_repo) FindById(id int) (*models.Authors, error) {
	var datas *models.Authors
	res := re.db.Order("id asc").Where("id = ?", id).Find(&datas)
	if res.Error != nil {
		return nil, errors.New("failed to found data")
	}
	if res.RowsAffected == 0 {
		return nil, errors.New("data not found")
	}
	return datas, nil
}

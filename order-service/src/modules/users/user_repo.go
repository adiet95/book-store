package users

import (
	"errors"

	"github.com/adiet95/book-store/book-service/src/database/models"
	"gorm.io/gorm"
)

type user_repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *user_repo {
	return &user_repo{db}
}

func (r *user_repo) FindAll(limit, offset int) (*models.Users, error) {
	var data models.Users

	result := r.db.Model(&data).Limit(limit).Offset(offset).Order("email asc").Find(&data)

	if result.Error != nil {
		return nil, errors.New("failed obtain datas")
	}
	return &data, nil
}

func (r *user_repo) Save(data *models.User) (*models.User, error) {
	var datas models.Users
	res := r.db.Where("LOWER(email) = ?", data.Email).Find(&datas)

	if res.RowsAffected != 0 {
		return nil, errors.New("email already registered")
	}

	result := r.db.Create(data)
	if result.Error != nil {
		return nil, errors.New("failled to obtain data")
	}
	return data, nil
}

func (re *user_repo) UpdateUser(data *models.User, email string) (*models.User, error) {
	res := re.db.Model(&data).Where("LOWER(email) = ?", email).Updates(&data)

	if res.Error != nil {
		return nil, errors.New("failed to update data")
	}
	return data, nil
}

func (re *user_repo) DeleteUser(email string) (*models.User, error) {
	var data *models.User
	var datas *models.Users
	res := re.db.Where("LOWER(email) = ?", email).Find(&datas)

	if res.RowsAffected == 0 {
		return nil, errors.New("data not found")
	}
	r := re.db.Model(data).Where("LOWER(email) = ?", email).Delete(&data)
	if r.Error != nil {
		return nil, errors.New("failed to delete data")
	}
	return nil, nil
}

func (re *user_repo) FindByEmail(email string) (*models.User, error) {
	var data *models.User
	var datas *models.Users

	res := re.db.Model(&datas).Where("LOWER(email) = ?", email).Find(&data)
	if res.Error != nil {
		return nil, errors.New("failed to find data")
	}
	if res.RowsAffected == 0 {
		return nil, errors.New("email not found")
	}
	return data, nil
}

func (re *user_repo) FindByName(name string) (*models.Users, error) {
	var datas *models.Users

	res := re.db.Model(&datas).Where("LOWER(full_name) LIKE ?", "%"+name+"%").Find(&datas)
	if res.Error != nil {
		return nil, errors.New("failed to find data")
	}
	if res.RowsAffected == 0 {
		return nil, errors.New("name not found")
	}
	return datas, nil
}

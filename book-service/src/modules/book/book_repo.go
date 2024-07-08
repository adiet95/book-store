package book

import (
	"errors"

	"github.com/adiet95/book-store/book-service/src/database/models"
	"gorm.io/gorm"
)

type book_repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *book_repo {
	return &book_repo{db}
}

func (r *book_repo) FindAll(limit, offset int) (*models.Books, error) {
	var datas *models.Books
	result := r.db.Model(&datas).Limit(limit).Offset(offset).Find(&datas)
	if result.Error != nil {
		return nil, errors.New("failed obtain datas")
	}
	return datas, nil
}

func (r *book_repo) Save(data *models.Book) (*models.Book, error) {
	res := r.db.Create(data)
	if res.Error != nil {
		return nil, errors.New("failed obtain datas")
	}
	return data, nil
}

func (re *book_repo) Update(data *models.Book, id int) (*models.Book, error) {
	res := re.db.Model(&data).Where("id = ?", id).Updates(&data)

	if res.Error != nil {
		return nil, errors.New("failed to update data")
	}
	return data, nil
}

func (re *book_repo) Delete(id int) (*models.Book, error) {
	var data *models.Book
	var datas *models.Books
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

func (re *book_repo) FindByName(name string) (*models.Books, error) {
	var datas *models.Books
	res := re.db.Order("id asc").Where("LOWER(order_name) LIKE ?", "%"+name+"%").Find(&datas)
	if res.Error != nil {
		return nil, errors.New("failed to found data")
	}
	if res.RowsAffected == 0 {
		return nil, errors.New("data not found")
	}
	return datas, nil
}

func (re *book_repo) FindById(id int) (*models.Books, error) {
	var datas *models.Books
	res := re.db.Order("id asc").Where("id = ?", id).Find(&datas)
	if res.Error != nil {
		return nil, errors.New("failed to found data")
	}
	if res.RowsAffected == 0 {
		return nil, errors.New("data not found")
	}
	return datas, nil
}

func (r *book_repo) GetUserId(email string) (*models.User, error) {
	var users *models.Users
	var user *models.User

	result := r.db.Model(&users).Where("email = ?", email).Find(&user)
	if result.Error != nil {
		return nil, errors.New("invalid user_id")
	}
	return user, nil
}

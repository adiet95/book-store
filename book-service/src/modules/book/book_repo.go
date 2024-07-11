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
	var datas models.Books
	var dataAuthor models.Author
	var dataCategory models.Category

	checkName := r.db.Where("LOWER(book_name) LIKE ?", "%"+data.BookName+"%").Find(&datas)
	if checkName.RowsAffected != 0 {
		return nil, errors.New("book name is exist")
	}

	checkAuthor := r.db.Where("author_id = ?", data.AuthorId).First(&dataAuthor)
	if checkAuthor.Error != nil {
		return nil, errors.New("failed to found data author")
	}
	if checkAuthor.RowsAffected == 0 {
		return nil, errors.New("author not found")
	}

	checkCategory := r.db.Where("category_id = ?", data.CategoryId).First(&dataCategory)
	if checkCategory.Error != nil {
		return nil, errors.New("failed to found data category")
	}
	if checkCategory.RowsAffected == 0 {
		return nil, errors.New("category not found")
	}

	res := r.db.Create(data)
	if res.Error != nil {
		return nil, errors.New("failed obtain datas")
	}
	return data, nil
}

func (re *book_repo) Update(data *models.Book, id int) (*models.Book, error) {
	var dataAuthor models.Author
	var dataCategory models.Category

	checkAuthor := re.db.Where("author_id = ?", data.AuthorId).First(&dataAuthor)
	if checkAuthor.Error != nil {
		return nil, errors.New("failed to found data author")
	}
	if checkAuthor.RowsAffected == 0 {
		return nil, errors.New("author not found")
	}

	checkCategory := re.db.Where("category_id = ?", data.CategoryId).First(&dataCategory)
	if checkCategory.Error != nil {
		return nil, errors.New("failed to found data category")
	}
	if checkCategory.RowsAffected == 0 {
		return nil, errors.New("category not found")
	}

	res := re.db.Model(&data).Where("book_id = ?", id).Updates(&data)

	if res.Error != nil {
		return nil, errors.New("failed to update data")
	}
	return data, nil
}

func (re *book_repo) Delete(id int) (*models.Book, error) {
	var data *models.Book
	var datas *models.Books
	res := re.db.Where("book_id = ?", id).Find(&datas)

	if res.RowsAffected == 0 {
		return nil, errors.New("data not found")
	}
	r := re.db.Model(data).Where("book_id = ?", id).Delete(&data)
	if r.Error != nil {
		return nil, errors.New("failed to delete data")
	}
	return nil, nil
}

func (re *book_repo) FindByName(name string) (*models.Books, error) {
	var datas *models.Books
	res := re.db.Order("book_id asc").Where("LOWER(book_name) LIKE ?", "%"+name+"%").Find(&datas)
	if res.Error != nil {
		return nil, errors.New("failed to found data")
	}
	if res.RowsAffected == 0 {
		return nil, errors.New("data not found")
	}
	return datas, nil
}

func (re *book_repo) FindById(id int) (*models.Book, error) {
	var datas *models.Book
	res := re.db.Order("book_id asc").Where("book_id = ?", id).First(&datas)
	if res.Error != nil {
		return nil, errors.New("failed to found data")
	}
	if res.RowsAffected == 0 {
		return nil, errors.New("data not found")
	}
	return datas, nil
}

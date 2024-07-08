package book

import (
	"github.com/adiet95/book-store/book-service/src/database/models"
	"github.com/adiet95/book-store/book-service/src/interfaces"
	"github.com/adiet95/book-store/book-service/src/libs"
)

type book_service struct {
	book_repo interfaces.BookRepo
}

func NewService(reps interfaces.BookRepo) *book_service {
	return &book_service{reps}
}

func (r *book_service) GetAll(limit, offset int) *libs.Response {
	data, err := r.book_repo.FindAll(limit, offset)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	return libs.New(data, 200, false)
}

func (re *book_service) Add(data *models.Book) *libs.Response {
	result, err := re.book_repo.Save(data)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	return libs.New(result, 200, false)
}

func (re *book_service) Update(data *models.Book, id int) *libs.Response {
	res, err := re.book_repo.Update(data, id)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	return libs.New(res, 200, false)
}

func (re *book_service) Delete(id int) *libs.Response {
	data, err := re.book_repo.Delete(id)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	return libs.New(data, 200, false)
}

func (re *book_service) Search(name string) *libs.Response {
	data, err := re.book_repo.FindByName(name)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	return libs.New(data, 200, false)
}

func (re *book_service) SearchId(id int) *libs.Response {
	data, err := re.book_repo.FindById(id)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	return libs.New(data, 200, false)
}

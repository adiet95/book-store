package category

import (
	"github.com/adiet95/book-store/category-service/src/database/models"
	"github.com/adiet95/book-store/category-service/src/interfaces"
	"github.com/adiet95/book-store/category-service/src/libs"
)

type category_service struct {
	category_repo interfaces.CategoryRepo
}

func NewService(reps interfaces.CategoryRepo) *category_service {
	return &category_service{reps}
}

func (r *category_service) GetAll(limit, offset int) *libs.Response {
	data, err := r.category_repo.FindAll(limit, offset)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	return libs.New(data, 200, false)
}

func (re *category_service) Add(data *models.Category) *libs.Response {
	result, err := re.category_repo.Save(data)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	return libs.New(result, 200, false)
}

func (re *category_service) Update(data *models.Category, id int) *libs.Response {
	res, err := re.category_repo.Update(data, id)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	return libs.New(res, 200, false)
}

func (re *category_service) Delete(id int) *libs.Response {
	data, err := re.category_repo.Delete(id)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	return libs.New(data, 200, false)
}

func (re *category_service) Search(name string) *libs.Response {
	data, err := re.category_repo.FindByName(name)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	return libs.New(data, 200, false)
}

func (re *category_service) SearchId(id int) *libs.Response {
	data, err := re.category_repo.FindById(id)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	return libs.New(data, 200, false)
}

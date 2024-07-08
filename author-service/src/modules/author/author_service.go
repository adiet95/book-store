package author

import (
	"github.com/adiet95/book-store/author-service/src/database/models"
	"github.com/adiet95/book-store/author-service/src/interfaces"
	"github.com/adiet95/book-store/author-service/src/libs"
)

type author_service struct {
	author_repo interfaces.AuthorRepo
}

func NewService(reps interfaces.AuthorRepo) *author_service {
	return &author_service{reps}
}

func (r *author_service) GetAll(limit, offset int) *libs.Response {
	data, err := r.author_repo.FindAll(limit, offset)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	return libs.New(data, 200, false)
}

func (re *author_service) Add(data *models.Author) *libs.Response {
	result, err := re.author_repo.Save(data)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	return libs.New(result, 200, false)
}

func (re *author_service) Update(data *models.Author, id int) *libs.Response {
	res, err := re.author_repo.Update(data, id)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	return libs.New(res, 200, false)
}

func (re *author_service) Delete(id int) *libs.Response {
	data, err := re.author_repo.Delete(id)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	return libs.New(data, 200, false)
}

func (re *author_service) Search(name string) *libs.Response {
	data, err := re.author_repo.FindByName(name)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	return libs.New(data, 200, false)
}

func (re *author_service) SearchId(id int) *libs.Response {
	data, err := re.author_repo.FindById(id)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	return libs.New(data, 200, false)
}

package interfaces

import (
	"github.com/adiet95/book-store/src/database/models"
	"github.com/adiet95/book-store/src/libs"
)

type UserRepo interface {
	FindAll(limit, offset int) (*models.Users, error)
	Save(data *models.User) (*models.User, error)
	UpdateUser(data *models.User, email string) (*models.User, error)
	DeleteUser(email string) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	FindByName(name string) (*models.Users, error)
}

type UserService interface {
	Add(data *models.User) *libs.Response
	Update(data *models.User, email string) *libs.Response
	Delete(email string) *libs.Response
	FindEmail(email string, limit, offset int) *libs.Response
	Search(email string) *libs.Response
	SearchName(name string) *libs.Response
}

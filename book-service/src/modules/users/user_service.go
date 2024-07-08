package users

import (
	"github.com/adiet95/book-store/book-service/src/database/models"
	"github.com/adiet95/book-store/book-service/src/interfaces"
	"github.com/adiet95/book-store/book-service/src/libs"
)

type user_service struct {
	user_repo interfaces.UserRepo
}

func NewService(reps interfaces.UserRepo) *user_service {
	return &user_service{reps}
}

func (re *user_service) Add(data *models.User) *libs.Response {
	valid := libs.Validation(data.Email, data.Password)
	if valid != nil {
		return libs.New(valid.Error(), 400, true)
	}

	hassPass, err := libs.HashPassword(data.Password)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}

	data.Password = hassPass
	result, err := re.user_repo.Save(data)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	return libs.New(result, 201, false)
}

func (re *user_service) Update(data *models.User, email string) *libs.Response {
	//Get old
	oldData, err := re.user_repo.FindByEmail(email)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}

	if oldData.Role == "admin" {
		valid := libs.Validation(data.Email, data.Password)
		if valid != nil {
			return libs.New(valid.Error(), 400, true)
		}

		//Hasing New Password and update data
		hassPass, err := libs.HashPassword(data.Password)
		if err != nil {
			return libs.New(err.Error(), 400, true)
		}
		data.Password = hassPass

		result, err := re.user_repo.UpdateUser(data, email)
		if err != nil {
			return libs.New(err.Error(), 400, true)
		}
		return libs.New(result, 202, false)
	}

	valid := libs.Validation(data.Email, data.Password)
	if valid != nil {
		return libs.New(valid.Error(), 400, true)
	}

	//Hasing New Password and update data
	hassPass, err := libs.HashPassword(data.Password)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}

	data.Email = oldData.Email
	data.Role = oldData.Role
	data.Password = hassPass

	result, err := re.user_repo.UpdateUser(data, email)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	return libs.New(result, 202, false)
}

func (re *user_service) Delete(email string) *libs.Response {

	data, err := re.user_repo.DeleteUser(email)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	return libs.New(data, 204, false)
}

func (re *user_service) FindEmail(email string, limit, offset int) *libs.Response {
	data, err := re.user_repo.FindByEmail(email)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	if data.Role == "admin" {
		datas, err := re.user_repo.FindAll(limit, offset)
		if err != nil {
			return libs.New(err.Error(), 400, true)
		}
		return libs.New(datas, 200, false)
	}
	return libs.New(data, 200, false)
}

func (re *user_service) Search(email string) *libs.Response {
	data, err := re.user_repo.FindByEmail(email)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	return libs.New(data, 200, false)
}

func (re *user_service) SearchName(name string) *libs.Response {
	data, err := re.user_repo.FindByName(name)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	return libs.New(data, 200, false)
}

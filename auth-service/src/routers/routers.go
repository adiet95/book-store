package routers

import (
	"errors"

	"github.com/adiet95/book-store/auth-service/src/database"
	auth "github.com/adiet95/book-store/auth-service/src/modules/auth"
	"github.com/adiet95/book-store/auth-service/src/modules/users"
	"github.com/gin-gonic/gin"
)

func New() (*gin.Engine, error) {
	mainRoute := gin.Default()

	db, err := database.New()
	if err != nil {
		return nil, errors.New("failed init database")
	}

	auth.New(mainRoute, db)
	users.New(mainRoute, db)

	return mainRoute, nil
}

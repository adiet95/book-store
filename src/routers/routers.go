package routers

import (
	"errors"

	"github.com/adiet95/book-store/src/database"
	auth "github.com/adiet95/book-store/src/modules/auth"
	"github.com/adiet95/book-store/src/modules/order"
	"github.com/adiet95/book-store/src/modules/users"
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
	order.New(mainRoute, db)

	return mainRoute, nil
}

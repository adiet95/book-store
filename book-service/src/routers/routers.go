package routers

import (
	"errors"

	"github.com/adiet95/book-store/book-service/src/database"
	"github.com/adiet95/book-store/book-service/src/modules/book"
	"github.com/gin-gonic/gin"
)

func New() (*gin.Engine, error) {
	mainRoute := gin.Default()

	db, err := database.New()
	if err != nil {
		return nil, errors.New("failed init database")
	}

	book.New(mainRoute, db)

	return mainRoute, nil
}

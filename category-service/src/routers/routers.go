package routers

import (
	"errors"

	"github.com/adiet95/book-store/category-service/src/database"
	"github.com/adiet95/book-store/category-service/src/modules/category"
	"github.com/gin-gonic/gin"
)

func New() (*gin.Engine, error) {
	mainRoute := gin.Default()

	db, err := database.New()
	if err != nil {
		return nil, errors.New("failed init database")
	}

	dbRedis, err := database.NewRedisClient()
	if err != nil {
		return nil, errors.New("failed init database")
	}
	category.New(mainRoute, db, dbRedis)

	return mainRoute, nil
}

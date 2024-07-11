package routers

import (
	"errors"

	"github.com/adiet95/book-store/author-service/src/database"
	"github.com/adiet95/book-store/author-service/src/modules/author"
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
		return nil, errors.New("failed init redis")
	}

	author.New(mainRoute, db, dbRedis)

	return mainRoute, nil
}

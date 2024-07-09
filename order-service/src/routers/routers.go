package routers

import (
	"errors"
	"github.com/adiet95/book-store/order-service/src/modules/stock"

	"github.com/adiet95/book-store/order-service/src/database"
	"github.com/adiet95/book-store/order-service/src/modules/order"
	"github.com/gin-gonic/gin"
)

func New() (*gin.Engine, error) {
	mainRoute := gin.Default()

	db, err := database.New()
	if err != nil {
		return nil, errors.New("failed init database")
	}

	stock.New(mainRoute, db)
	order.New(mainRoute, db)

	return mainRoute, nil
}

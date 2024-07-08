package author

import (
	"github.com/adiet95/book-store/author-service/src/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func New(rt *gin.Engine, db *gorm.DB) {
	repo := NewRepo(db)
	svc := NewService(repo)
	ctrl := NewCtrl(svc)

	route := rt.Group("/category").Use(middleware.CheckAuth())
	{
		route.POST("", ctrl.Add)
		route.PUT("", ctrl.Update)
		route.DELETE("", ctrl.Delete)
		route.GET("", ctrl.GetAll)
		route.GET("/search", ctrl.Search)
		route.GET("/detail", ctrl.SearchId)
	}
}

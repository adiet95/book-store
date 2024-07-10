package category

import (
	"github.com/adiet95/book-store/category-service/src/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func New(rt *gin.Engine, db *gorm.DB) {
	repo := NewRepo(db)
	svc := NewService(repo)
	ctrl := NewCtrl(svc)

	route := rt.Group("/category").Use(middleware.CheckAuth())
	{
		route.POST("", middleware.CheckAuthor(), middleware.CheckAuthor(), ctrl.Add)
		route.PUT("/:id", middleware.CheckAuthor(), middleware.CheckAuthor(), ctrl.Update)
		route.DELETE("/:id", middleware.CheckAuthor(), middleware.CheckAuthor(), ctrl.Delete)
		route.GET("", ctrl.GetAll)
		route.GET("/search", ctrl.Search)
		route.GET("/:id", ctrl.SearchId)
	}
}

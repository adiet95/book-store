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

	route := rt.Group("/author").Use(middleware.CheckAuth())
	{
		route.POST("", middleware.CheckAuthor(), ctrl.Add)
		route.PUT("", middleware.CheckAuthor(), ctrl.Update)
		route.DELETE("", middleware.CheckAuthor(), ctrl.Delete)
		route.GET("", ctrl.GetAll)
		route.GET("/search", ctrl.Search)
		route.GET("/detail", ctrl.SearchId)
	}
}

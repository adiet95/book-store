package users

import (
	"github.com/adiet95/book-store/src/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func New(rt *gin.Engine, db *gorm.DB) {
	repo := NewRepo(db)
	svc := NewService(repo)
	ctrl := NewCtrl(svc)

	route := rt.Group("/user").Use(middleware.CheckAuth())
	{
		route.GET("", ctrl.GetAll)
		route.POST("", middleware.CheckAuthor(), ctrl.Add)
		route.PUT("", ctrl.Update)
		route.DELETE("", middleware.CheckAuthor(), ctrl.Delete)
		route.GET("/detail", middleware.CheckAuthor(), ctrl.Search)
		route.GET("/search", middleware.CheckAuthor(), ctrl.SearchName)
	}

}

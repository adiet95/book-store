package users

import (
	"github.com/adiet95/book-store/auth-service/src/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func New(rt *gin.Engine, db *gorm.DB) {
	repo := NewRepo(db)
	svc := NewService(repo)
	ctrl := NewCtrl(svc)

	route := rt.Group("/user").Use(middleware.CheckAuth(), middleware.CheckAuthor())
	{
		route.GET("", ctrl.GetAll)
		route.POST("", middleware.CheckAuthor(), ctrl.Add)
		route.PUT("", ctrl.Update)
		route.DELETE("/:id", middleware.CheckAuthor(), ctrl.Delete)
		route.GET("/detail", middleware.CheckAuthor(), ctrl.Search)
		route.GET("/search", middleware.CheckAuthor(), ctrl.SearchName)
	}

}

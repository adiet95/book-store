package category

import (
	"github.com/adiet95/book-store/category-service/src/middleware"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

func New(rt *gin.Engine, db *gorm.DB, redisClient *redis.Client) {
	repo := NewRepo(db, redisClient)
	svc := NewService(repo)
	ctrl := NewCtrl(svc)

	route := rt.Group("/category").Use(middleware.CheckAuth())
	{
		route.POST("", middleware.CheckAuthor(), ctrl.Add)
		route.PUT("/:id", middleware.CheckAuthor(), ctrl.Update)
		route.DELETE("/:id", middleware.CheckAuthor(), ctrl.Delete)
		route.GET("", ctrl.GetAll)
		route.GET("/search", ctrl.Search)
		route.GET("/:id", ctrl.SearchId)
	}
}

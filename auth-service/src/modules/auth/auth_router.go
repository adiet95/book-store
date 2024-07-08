package auth

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func New(rt *gin.Engine, db *gorm.DB) {
	repo := NewRepo(db)
	svc := NewService(repo)
	ctrl := NewCtrl(svc)

	route := rt.Group("")
	{
		route.POST("/login", ctrl.SignIn)
		route.POST("/register", ctrl.Register)
	}
}

package middleware

import (
	"strings"

	"github.com/adiet95/book-store/src/libs"
	"github.com/gin-gonic/gin"
)

func CheckAuthor() gin.HandlerFunc {
	return func(c *gin.Context) {
		headerToken := c.GetHeader("Authorization")

		if !strings.Contains(headerToken, "Bearer") {
			libs.New("invalid header type", 401, true).Send(c)
			c.Abort()
		}
		token := strings.Replace(headerToken, "Bearer ", "", -1)

		checkToken, err := libs.CheckToken(token)
		if err != nil {
			libs.New(err.Error(), 401, true).Send(c)
			c.Abort()
		}
		if checkToken.Role != "admin" {
			libs.New("forbidden access", 401, true).Send(c)
			c.Abort()
		}
		c.Set("role", checkToken.Role)
		c.Next()
	}
}

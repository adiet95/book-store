package middleware

import (
	"strings"

	"github.com/adiet95/book-store/auth-service/src/libs"
	"github.com/gin-gonic/gin"
)

func CheckAuth() gin.HandlerFunc {
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

		c.Set("email", checkToken.Email)
		c.Next()
	}
}

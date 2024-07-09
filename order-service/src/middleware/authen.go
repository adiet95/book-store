package middleware

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"os"
	"strings"

	pb "github.com/adiet95/book-store/order-service/src/database/grpc-model"
	"github.com/adiet95/book-store/order-service/src/libs"
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

		conn, err := grpc.NewClient(os.Getenv("GRPC_ADDRESS"), grpc.WithInsecure())
		if err != nil {
			errStr := fmt.Sprintf("Error mysql connection %s", err.Error())
			libs.New(errStr, 401, true)
			panic(err)
		}
		clientGrpc := pb.NewAuthenticationClient(conn)
		req := pb.AuthRequest{Token: token}
		tokenData, err := clientGrpc.ValidateToken(context.Background(), &req)
		if err != nil {
			libs.New(err.Error(), 400, true).Send(c)
			c.Abort()
		}

		if tokenData != nil {
			if tokenData.IsValidate != true {
				libs.New(err.Error(), 401, true).Send(c)
				c.Abort()
			}
		} else {
			libs.New(err.Error(), 401, true).Send(c)
			c.Abort()
		}

		c.Set("validate", tokenData.IsValidate)
		c.Next()
	}
}

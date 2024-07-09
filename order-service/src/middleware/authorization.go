package middleware

import (
	"context"
	"fmt"
	pb "github.com/adiet95/book-store/order-service/src/database/grpc-model"
	"google.golang.org/grpc"
	"os"
	"strings"

	"github.com/adiet95/book-store/order-service/src/libs"
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

		conn, err := grpc.NewClient(os.Getenv("GRPC_ADDRESS"), grpc.WithInsecure())
		if err != nil {
			errStr := fmt.Sprintf("Error mysql connection %s", err.Error())
			libs.New(errStr, 401, true)
			panic(err)
		}
		clientGrpc := pb.NewAuthenticationClient(conn)
		req := pb.AuthRequest{Token: token}
		tokenData, err := clientGrpc.ValidateToken(context.Background(), &req)

		if err != nil || tokenData != nil {
			if tokenData.IsValidate != true {
				libs.New(err.Error(), 401, true).Send(c)
				c.Abort()
			}
			if tokenData.Role != "admin" {
				libs.New("forbidden access", 401, true).Send(c)
				c.Abort()
			}
		} else {
			libs.New(err.Error(), 400, true).Send(c)
			c.Abort()
		}

		c.Set("role", tokenData.Role)
		c.Next()
	}
}

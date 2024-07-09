package grpc_auth

import (
	"context"
	"errors"
	"fmt"
	pb "github.com/adiet95/book-store/auth-service/src/database/grpc-model"
	"github.com/adiet95/book-store/auth-service/src/interfaces"
	"github.com/adiet95/book-store/auth-service/src/libs"
	"sync"
)

type AuthGrpc struct {
	pb.UnimplementedAuthenticationServer
	mx       sync.Mutex
	Data     []*pb.AuthData
	authRepo interfaces.AuthRepo
}

func InitAuthGrpc(authRepo interfaces.AuthRepo) *AuthGrpc {
	return &AuthGrpc{
		authRepo: authRepo,
	}
}

func (s *AuthGrpc) ValidateToken(c context.Context, authReq *pb.AuthRequest) (*pb.AuthResponse, error) {
	result := &pb.AuthResponse{}
	checkToken, err := libs.CheckToken(authReq.Token)
	if err != nil || checkToken == nil {
		message := fmt.Sprintf("Invalid Token err: %v", err)
		err = errors.New(message)
		return nil, err
	}
	result.IsValidate = true
	result.Role = checkToken.Role
	err = nil
	return result, nil
}

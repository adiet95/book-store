package config

import (
	"errors"
	"fmt"
	"github.com/adiet95/book-store/auth-service/src/database"
	pb "github.com/adiet95/book-store/auth-service/src/database/grpc-model"
	"github.com/adiet95/book-store/auth-service/src/modules/auth"
	grpc_auth "github.com/adiet95/book-store/auth-service/src/modules/grpc-auth"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"os"
)

var GrpcServeCmd = &cobra.Command{
	Use:   "serve-grpc",
	Short: "start apllication",
	RunE:  GrpcServerCmd,
}

func GrpcServerCmd(cmd *cobra.Command, args []string) error {
	addr := fmt.Sprintf(":%s", os.Getenv("GRPC_PORT"))
	fmt.Println(addr, "<<ADDR")
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln("Error when listen server address", err.Error())
	}

	grpcServer := grpc.NewServer()
	db, err := database.New()
	if err != nil {
		return errors.New("failed init database")
	}

	repo := auth.NewRepo(db)
	authService := grpc_auth.InitAuthGrpc(repo)

	pb.RegisterAuthenticationServer(grpcServer, authService)

	err = grpcServer.Serve(listen)
	if err != nil {
		return err
	}
	message := fmt.Sprintf("Listen GRPC PORT %v", os.Getenv("GRPC_PORT"))
	status.New(200, message)

	return nil
}

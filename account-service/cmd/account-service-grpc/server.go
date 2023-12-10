package main

import (
	"context"
	"net"

	"github.com/stephanvebrian/expense-manager/account-service/internal/config"
	grpcHandler "github.com/stephanvebrian/expense-manager/account-service/internal/handler/grpc"
	pb "github.com/stephanvebrian/expense-manager/common/proto/account-service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type handler struct {
	grpc grpcHandler.IHandler
}

func startServer(ctx context.Context, cfg config.Config, handler handler) error {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		return err
	}

	grpcOpt := []grpc.ServerOption{}

	server := grpc.NewServer(
		grpcOpt...,
	)

	pb.RegisterAccountServiceServer(server, handler.grpc)
	reflection.Register(server)

	return server.Serve(lis)
}

package main

import (
	"context"
	"net"

	pb "github.com/stephanvebrian/expense-manager/common/proto/transaction-service"
	"github.com/stephanvebrian/expense-manager/transaction-service/internal/config"
	grpcHandler "github.com/stephanvebrian/expense-manager/transaction-service/internal/handler/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type handler struct {
	grpc grpcHandler.IHandler
}

func startServer(ctx context.Context, cfg *config.Config, handler handler) error {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		return err
	}

	grpcOpt := []grpc.ServerOption{}

	server := grpc.NewServer(
		grpcOpt...,
	)

	pb.RegisterTransactionServiceServer(server, handler.grpc)
	reflection.Register(server)

	return server.Serve(lis)
}

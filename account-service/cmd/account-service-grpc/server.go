package main

import (
	"context"
	"net"

	"github.com/stephanvebrian/expense-manager/account-service/internal/config"
	pb "github.com/stephanvebrian/expense-manager/common/proto/account-service"
	"google.golang.org/grpc"
)

func startServer(ctx context.Context, cfg config.Config) error {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		return err
	}

	grpcOpt := []grpc.ServerOption{}

	server := grpc.NewServer(
		grpcOpt...,
	)

	pb.RegisterAccountServiceServer(server, nil)

	return server.Serve(lis)
}

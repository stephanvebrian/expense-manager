package main

import (
	"context"

	"github.com/stephanvebrian/expense-manager/account-service/internal/config"
	grpcHandler "github.com/stephanvebrian/expense-manager/account-service/internal/handler/grpc"
)

func startApp(cfg config.Config) error {
	ctx := context.Background()

	// initialize external client, database, redis, etc

	// initialize repo

	// initialize usecase

	// initializae handler
	handlerGrpc := grpcHandler.New(cfg)

	handler := handler{
		grpc: handlerGrpc,
	}

	return startServer(ctx, cfg, handler)
}

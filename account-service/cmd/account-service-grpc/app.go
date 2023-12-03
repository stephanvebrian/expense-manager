package main

import (
	"context"

	"github.com/stephanvebrian/expense-manager/account-service/internal/config"
)

func startApp(cfg config.Config) error {
	ctx := context.Background()

	// initialize external client, database, redis, etc

	// initialize repo

	// initialize usecase

	return startServer(ctx, cfg)
}

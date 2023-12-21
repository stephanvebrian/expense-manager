package grpc

import (
	"context"

	pb "github.com/stephanvebrian/expense-manager/common/proto/transaction-service"
	"github.com/stephanvebrian/expense-manager/transaction-service/internal/config"
)

type IHandler interface {
	Ping(context.Context, *pb.GetPingRequest) (*pb.GetPingResponse, error)
}

type handler struct {
	cfg *config.Config
}

func New(cfg *config.Config) IHandler {
	return &handler{
		cfg: cfg,
	}
}

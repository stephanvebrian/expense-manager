package grpc

import (
	"context"

	"github.com/stephanvebrian/expense-manager/account-service/internal/config"
	pb "github.com/stephanvebrian/expense-manager/common/proto/account-service"
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

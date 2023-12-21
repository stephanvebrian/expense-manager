package grpc

import (
	"context"

	pb "github.com/stephanvebrian/expense-manager/common/proto/transaction-service"
)

func (h *handler) Ping(context.Context, *pb.GetPingRequest) (*pb.GetPingResponse, error) {
	return &pb.GetPingResponse{
		Message: "Pong!",
	}, nil
}

package grpc

import (
	"context"

	pb "github.com/stephanvebrian/expense-manager/common/proto/account-service"
)

func (h *handler) Login(ctx context.Context, req *pb.PostLoginRequest) (*pb.PostLoginResponse, error) {
	// TODO: implement login handler and logic
	return &pb.PostLoginResponse{
		Token: &pb.TokenResponse{
			AccessToken:  "dummy-token",
			RefreshToken: "dummy-refresh-token",
		},
	}, nil
}

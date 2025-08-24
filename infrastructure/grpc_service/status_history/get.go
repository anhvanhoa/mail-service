package grpcstatushistory

import (
	"context"

	proto_status_history "github.com/anhvanhoa/sf-proto/gen/status_history/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (sh *statusHistoryService) GetStatusHistory(ctx context.Context, req *proto_status_history.GetStatusHistoryRequest) (*proto_status_history.GetStatusHistoryResponse, error) {
	// This method is not implemented in the usecase, so we'll return an error
	return nil, status.Errorf(codes.Unimplemented, "Method GetStatusHistory not implemented")
}

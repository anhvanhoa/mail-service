package grpcstatushistory

import (
	"context"
	proto "mail-service/proto/gen/status_history/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (sh *statusHistoryService) GetStatusHistory(ctx context.Context, req *proto.GetStatusHistoryRequest) (*proto.GetStatusHistoryResponse, error) {
	// This method is not implemented in the usecase, so we'll return an error
	return nil, status.Errorf(codes.Unimplemented, "Method GetStatusHistory not implemented")
}

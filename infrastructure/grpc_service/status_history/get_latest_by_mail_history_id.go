package grpcservice

import (
	"context"
	proto "mail-service/proto/gen/status_history/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (sh *statusHistoryService) GetLatestStatusHistoryByMailHistoryId(ctx context.Context, req *proto.GetLatestStatusHistoryByMailHistoryIdRequest) (*proto.GetLatestStatusHistoryByMailHistoryIdResponse, error) {
	statusHistory, err := sh.getLatestStatusHistoryByMailHistoryIdUsecase.Execute(ctx, req.MailHistoryId)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Status history không tồn tại: %v", err)
	}

	return &proto.GetLatestStatusHistoryByMailHistoryIdResponse{
		Message: "Latest status history retrieved successfully",
		StatusHistory: &proto.StatusHistory{
			Status:        string(statusHistory.Status),
			MailHistoryId: statusHistory.MailHistoryId,
			Message:       statusHistory.Message,
			CreatedAt:     statusHistory.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		},
	}, nil
}

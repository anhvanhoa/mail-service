package grpcstatushistory

import (
	"context"
	"mail-service/domain/entity"
	"time"

	proto_status_history "github.com/anhvanhoa/sf-proto/gen/status_history/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (sh *statusHistoryService) CreateStatusHistory(ctx context.Context, req *proto_status_history.CreateStatusHistoryRequest) (*proto_status_history.CreateStatusHistoryResponse, error) {
	statusHistory := entity.StatusHistory{
		Status:        entity.StatusMail(req.Status),
		MailHistoryId: req.MailHistoryId,
		Message:       req.Message,
		CreatedAt:     time.Now(),
	}

	if req.CreatedAt != "" {
		createdAt, err := time.Parse(time.RFC3339, req.CreatedAt)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "Thời gian tạo không hợp lệ")
		}
		statusHistory.CreatedAt = createdAt
	}

	err := sh.createStatusHistoryUsecase.Execute(ctx, &statusHistory)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Lỗi tạo status history: %v", err)
	}

	return &proto_status_history.CreateStatusHistoryResponse{
		Message: "Status history created successfully",
		StatusHistory: &proto_status_history.StatusHistory{
			Status:        string(statusHistory.Status),
			MailHistoryId: statusHistory.MailHistoryId,
			Message:       statusHistory.Message,
			CreatedAt:     statusHistory.CreatedAt.Format(time.RFC3339),
		},
	}, nil
}

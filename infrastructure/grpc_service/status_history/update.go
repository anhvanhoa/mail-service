package grpcstatushistory

import (
	"context"
	"mail-service/domain/entity"
	proto "mail-service/proto/gen/status_history/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (sh *statusHistoryService) UpdateStatusHistory(ctx context.Context, req *proto.UpdateStatusHistoryRequest) (*proto.UpdateStatusHistoryResponse, error) {
	statusHistory := entity.StatusHistory{
		Status:        entity.StatusMail(req.Status),
		MailHistoryId: req.MailHistoryId,
		Message:       req.Message,
	}

	err := sh.updateStatusHistoryUsecase.Execute(ctx, &statusHistory)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Lỗi cập nhật status history: %v", err)
	}

	return &proto.UpdateStatusHistoryResponse{
		Message: "Status history updated successfully",
		StatusHistory: &proto.StatusHistory{
			Status:        string(statusHistory.Status),
			MailHistoryId: statusHistory.MailHistoryId,
			Message:       statusHistory.Message,
			CreatedAt:     statusHistory.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		},
	}, nil
}

package grpcstatushistory

import (
	"context"
	"mail-service/domain/entity"
	proto "mail-service/proto/gen/status_history/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (sh *statusHistoryService) DeleteStatusHistory(ctx context.Context, req *proto.DeleteStatusHistoryRequest) (*proto.DeleteStatusHistoryResponse, error) {
	err := sh.deleteStatusHistoryUsecase.Execute(ctx, req.MailHistoryId, entity.StatusMail(req.Status))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Lỗi xóa status history: %v", err)
	}

	return &proto.DeleteStatusHistoryResponse{
		Message: "Status history deleted successfully",
	}, nil
}

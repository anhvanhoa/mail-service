package grpcstatushistory

import (
	"context"
	"mail-service/domain/entity"

	proto_status_history "github.com/anhvanhoa/sf-proto/gen/status_history/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (sh *statusHistoryService) DeleteStatusHistory(ctx context.Context, req *proto_status_history.DeleteStatusHistoryRequest) (*proto_status_history.DeleteStatusHistoryResponse, error) {
	err := sh.deleteStatusHistoryUsecase.Execute(ctx, req.MailHistoryId, entity.StatusMail(req.Status))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Lỗi xóa status history: %v", err)
	}

	return &proto_status_history.DeleteStatusHistoryResponse{
		Message: "Status history deleted successfully",
	}, nil
}

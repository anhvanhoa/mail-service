package grpcstatushistory

import (
	"context"

	proto_status_history "github.com/anhvanhoa/sf-proto/gen/status_history/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (sh *statusHistoryService) GetAllStatusHistory(ctx context.Context, req *proto_status_history.GetAllStatusHistoryRequest) (*proto_status_history.GetAllStatusHistoryResponse, error) {
	result, err := sh.getAllStatusHistoryUsecase.Execute(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Lỗi lấy danh sách status history: %v", err)
	}

	// Convert to proto response
	var statusHistories []*proto_status_history.StatusHistory
	for _, sh := range result {
		statusHistories = append(statusHistories, &proto_status_history.StatusHistory{
			Status:        string(sh.Status),
			MailHistoryId: sh.MailHistoryId,
			Message:       sh.Message,
			CreatedAt:     sh.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		})
	}

	return &proto_status_history.GetAllStatusHistoryResponse{
		Message:         "Status histories retrieved successfully",
		Total:           int32(len(statusHistories)),
		Page:            1,
		Limit:           int32(len(statusHistories)),
		StatusHistories: statusHistories,
	}, nil
}

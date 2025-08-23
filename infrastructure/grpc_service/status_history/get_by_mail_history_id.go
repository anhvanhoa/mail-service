package grpcstatushistory

import (
	"context"
	proto "mail-service/proto/gen/status_history/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (sh *statusHistoryService) GetStatusHistoryByMailHistoryId(ctx context.Context, req *proto.GetStatusHistoryByMailHistoryIdRequest) (*proto.GetStatusHistoryByMailHistoryIdResponse, error) {
	result, err := sh.getStatusHistoryByMailHistoryIdUsecase.Execute(ctx, req.MailHistoryId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Lỗi lấy status history theo mail history id: %v", err)
	}

	// Convert to proto response
	var statusHistories []*proto.StatusHistory
	for _, sh := range result {
		statusHistories = append(statusHistories, &proto.StatusHistory{
			Status:        string(sh.Status),
			MailHistoryId: sh.MailHistoryId,
			Message:       sh.Message,
			CreatedAt:     sh.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		})
	}

	return &proto.GetStatusHistoryByMailHistoryIdResponse{
		Message:         "Status histories retrieved successfully",
		Total:           int32(len(statusHistories)),
		Page:            1,
		Limit:           int32(len(statusHistories)),
		StatusHistories: statusHistories,
	}, nil
}

package grpcmailhistory

import (
	"context"
	proto "mail-service/proto/gen/mail_history/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (mh *mailHistoryService) DeleteMailHistory(ctx context.Context, req *proto.DeleteMailHistoryRequest) (*proto.DeleteMailHistoryResponse, error) {
	err := mh.deleteMailHistoryUsecase.Execute(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Lỗi xóa mail history: %v", err)
	}

	return &proto.DeleteMailHistoryResponse{
		Message: "Mail history deleted successfully",
	}, nil
}

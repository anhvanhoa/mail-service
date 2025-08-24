package grpcmailhistory

import (
	"context"

	proto_mail_history "github.com/anhvanhoa/sf-proto/gen/mail_history/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (mh *mailHistoryService) DeleteMailHistory(ctx context.Context, req *proto_mail_history.DeleteMailHistoryRequest) (*proto_mail_history.DeleteMailHistoryResponse, error) {
	err := mh.deleteMailHistoryUsecase.Execute(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Lỗi xóa mail history: %v", err)
	}

	return &proto_mail_history.DeleteMailHistoryResponse{
		Message: "Mail history deleted successfully",
	}, nil
}

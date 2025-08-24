package grpcmailstatus

import (
	"context"
	"mail-service/domain/entity"

	proto_mail_status "github.com/anhvanhoa/sf-proto/gen/mail_status/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (ms *mailStatusService) DeleteMailStatus(ctx context.Context, req *proto_mail_status.DeleteMailStatusRequest) (*proto_mail_status.DeleteMailStatusResponse, error) {
	err := ms.deleteMailStatusUsecase.Execute(ctx, entity.StatusMail(req.Status))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Lỗi xóa mail status: %v", err)
	}

	return &proto_mail_status.DeleteMailStatusResponse{
		Message: "Mail status deleted successfully",
	}, nil
}

package grpcservice

import (
	"context"
	"mail-service/domain/entity"
	proto "mail-service/proto/gen/mail_status/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (ms *mailStatusService) DeleteMailStatus(ctx context.Context, req *proto.DeleteMailStatusRequest) (*proto.DeleteMailStatusResponse, error) {
	err := ms.deleteMailStatusUsecase.Execute(ctx, entity.StatusMail(req.Status))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Lỗi xóa mail status: %v", err)
	}

	return &proto.DeleteMailStatusResponse{
		Message: "Mail status deleted successfully",
	}, nil
}

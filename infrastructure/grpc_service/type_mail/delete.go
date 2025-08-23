package grpctypemail

import (
	"context"
	proto "mail-service/proto/gen/type_mail/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (tm *typeMailService) DeleteTypeMail(ctx context.Context, req *proto.DeleteTypeMailRequest) (*proto.DeleteTypeMailResponse, error) {
	err := tm.deleteTypeMailUsecase.Execute(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Lỗi xóa type mail: %v", err)
	}

	return &proto.DeleteTypeMailResponse{
		Message: "Type mail deleted successfully",
	}, nil
}

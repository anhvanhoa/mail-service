package grpctypemail

import (
	"context"

	proto_type_mail "github.com/anhvanhoa/sf-proto/gen/type_mail/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (tm *typeMailService) DeleteTypeMail(ctx context.Context, req *proto_type_mail.DeleteTypeMailRequest) (*proto_type_mail.DeleteTypeMailResponse, error) {
	err := tm.deleteTypeMailUsecase.Execute(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Lỗi xóa type mail: %v", err)
	}

	return &proto_type_mail.DeleteTypeMailResponse{
		Message: "Type mail deleted successfully",
	}, nil
}

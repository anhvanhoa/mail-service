package grpcservice

import (
	"context"
	proto "mail-service/proto/gen/mail_provider/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (mp *mailProviderService) DeleteMailProvider(ctx context.Context, req *proto.DeleteMailProviderRequest) (*proto.DeleteMailProviderResponse, error) {
	err := mp.deleteMailProviderUsecase.Execute(ctx, req.Email)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Lỗi xóa mail provider: %v", err)
	}

	return &proto.DeleteMailProviderResponse{
		Message: "Mail provider deleted successfully",
	}, nil
}

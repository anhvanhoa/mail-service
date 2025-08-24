package grpcmailprovider

import (
	"context"

	proto_mail_provider "github.com/anhvanhoa/sf-proto/gen/mail_provider/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (mp *mailProviderService) DeleteMailProvider(ctx context.Context, req *proto_mail_provider.DeleteMailProviderRequest) (*proto_mail_provider.DeleteMailProviderResponse, error) {
	err := mp.deleteMailProviderUsecase.Execute(ctx, req.Email)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Lỗi xóa mail provider: %v", err)
	}

	return &proto_mail_provider.DeleteMailProviderResponse{
		Message: "Mail provider deleted successfully",
	}, nil
}

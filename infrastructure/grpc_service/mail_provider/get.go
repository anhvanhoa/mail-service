package grpcservice

import (
	"context"
	proto "mail-service/proto/gen/mail_provider/v1"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (mp *mailProviderService) GetMailProvider(ctx context.Context, req *proto.GetMailProviderRequest) (*proto.GetMailProviderResponse, error) {
	mailProvider, err := mp.getMailProviderUsecase.Execute(ctx, req.Email)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Mail provider không tồn tại: %v", err)
	}

	return &proto.GetMailProviderResponse{
		Message: "Mail provider retrieved successfully",
		MailProvider: &proto.MailProvider{
			Email:      mailProvider.Email,
			Password:   mailProvider.Password,
			UserName:   mailProvider.UserName,
			Port:       int32(mailProvider.Port),
			Host:       mailProvider.Host,
			Encryption: mailProvider.Encryption,
			Name:       mailProvider.Name,
			TypeId:     mailProvider.TypeId,
			CreatedBy:  mailProvider.CreatedBy,
			CreatedAt:  mailProvider.CreatedAt.Format(time.RFC3339),
			UpdatedAt:  mailProvider.UpdatedAt.Format(time.RFC3339),
		},
	}, nil
}

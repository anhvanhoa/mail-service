package grpcmailprovider

import (
	"context"
	"time"

	proto_mail_provider "github.com/anhvanhoa/sf-proto/gen/mail_provider/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (mp *mailProviderService) GetMailProvider(ctx context.Context, req *proto_mail_provider.GetMailProviderRequest) (*proto_mail_provider.GetMailProviderResponse, error) {
	mailProvider, err := mp.getMailProviderUsecase.Execute(ctx, req.Email)
	if mailProvider == nil {
		return nil, status.Errorf(codes.NotFound, "Mail provider không tồn tại")
	}
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Mail provider không tồn tại: %v", err)
	}

	return &proto_mail_provider.GetMailProviderResponse{
		Message: "Mail provider retrieved successfully",
		MailProvider: &proto_mail_provider.MailProvider{
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

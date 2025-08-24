package grpcmailprovider

import (
	"context"
	"mail-service/domain/common"
	"mail-service/domain/entity"
	"time"

	proto_mail_provider "github.com/anhvanhoa/sf-proto/gen/mail_provider/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (mp *mailProviderService) UpdateMailProvider(ctx context.Context, req *proto_mail_provider.UpdateMailProviderRequest) (*proto_mail_provider.UpdateMailProviderResponse, error) {
	now := time.Now()
	mailProvider := entity.MailProvider{
		Email:      req.Email,
		Password:   req.Password,
		UserName:   req.UserName,
		Port:       int(req.Port),
		Host:       req.Host,
		Encryption: req.Encryption,
		Name:       req.Name,
		TypeId:     req.TypeId,
		UpdatedAt:  &now,
		Status:     common.Status(req.Status),
	}

	err := mp.updateMailProviderUsecase.Execute(ctx, &mailProvider)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Lỗi cập nhật mail provider: %v", err)
	}

	return &proto_mail_provider.UpdateMailProviderResponse{
		Message: "Mail provider updated successfully",
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

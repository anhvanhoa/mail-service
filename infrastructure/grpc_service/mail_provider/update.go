package grpcservice

import (
	"context"
	"mail-service/domain/entity"
	proto "mail-service/proto/gen/mail_provider/v1"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (mp *mailProviderService) UpdateMailProvider(ctx context.Context, req *proto.UpdateMailProviderRequest) (*proto.UpdateMailProviderResponse, error) {
	mailProvider := entity.MailProvider{
		Email:      req.Email,
		Password:   req.Password,
		UserName:   req.UserName,
		Port:       int(req.Port),
		Host:       req.Host,
		Encryption: req.Encryption,
		Name:       req.Name,
		TypeId:     req.TypeId,
		UpdatedAt:  time.Now(),
	}

	err := mp.updateMailProviderUsecase.Execute(ctx, &mailProvider)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Lỗi cập nhật mail provider: %v", err)
	}

	return &proto.UpdateMailProviderResponse{
		Message: "Mail provider updated successfully",
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

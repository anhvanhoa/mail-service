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

func (mp *mailProviderService) CreateMailProvider(ctx context.Context, req *proto_mail_provider.CreateMailProviderRequest) (*proto_mail_provider.CreateMailProviderResponse, error) {
	mailProvider := mp.createEntityMailProvider(req)
	if req.CreatedAt != "" {
		createdAt, err := time.Parse(time.RFC3339, req.CreatedAt)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "Thời gian tạo không hợp lệ")
		}
		mailProvider.CreatedAt = createdAt
	}

	err := mp.createMailProviderUsecase.Execute(ctx, mailProvider)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Lỗi tạo mail provider: %v", err)
	}

	return &proto_mail_provider.CreateMailProviderResponse{
		Message: "Mail provider created successfully",
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
		},
	}, nil
}

func (mp *mailProviderService) createEntityMailProvider(req *proto_mail_provider.CreateMailProviderRequest) *entity.MailProvider {
	return &entity.MailProvider{
		Email:      req.Email,
		Password:   req.Password,
		UserName:   req.UserName,
		Port:       int(req.Port),
		Name:       req.Name,
		TypeId:     req.TypeId,
		CreatedBy:  req.CreatedBy,
		CreatedAt:  time.Now(),
		Status:     common.StatusActive,
		Host:       req.Host,
		Encryption: req.Encryption,
	}
}

package grpcservice

import (
	"context"
	"mail-service/domain/entity"
	proto "mail-service/proto/gen/mail_provider/v1"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (mp *mailProviderService) CreateMailProvider(ctx context.Context, req *proto.CreateMailProviderRequest) (*proto.CreateMailProviderResponse, error) {
	mailProvider := entity.MailProvider{
		Email:      req.Email,
		Password:   req.Password,
		UserName:   req.UserName,
		Port:       int(req.Port),
		Host:       req.Host,
		Encryption: req.Encryption,
		Name:       req.Name,
		TypeId:     req.TypeId,
		CreatedBy:  req.CreatedBy,
		CreatedAt:  time.Now(),
	}

	if req.CreatedAt != "" {
		createdAt, err := time.Parse(time.RFC3339, req.CreatedAt)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "Thời gian tạo không hợp lệ")
		}
		mailProvider.CreatedAt = createdAt
	}

	err := mp.createMailProviderUsecase.Execute(ctx, &mailProvider)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Lỗi tạo mail provider: %v", err)
	}

	return &proto.CreateMailProviderResponse{
		Message: "Mail provider created successfully",
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

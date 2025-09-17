package grpcmailprovider

import (
	"context"
	"time"

	proto_mail_provider "github.com/anhvanhoa/sf-proto/gen/mail_provider/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (mp *mailProviderService) GetAllMailProvider(ctx context.Context, req *proto_mail_provider.GetAllMailProviderRequest) (*proto_mail_provider.GetAllMailProviderResponse, error) {
	result, err := mp.getAllMailProviderUsecase.Execute(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Lỗi lấy danh sách mail provider: %v", err)
	}

	var mailProviders []*proto_mail_provider.MailProvider
	for _, mp := range result {
		var updatedAt string
		if mp.UpdatedAt != nil {
			updatedAt = mp.UpdatedAt.Format(time.RFC3339)
		}
		mailProviders = append(mailProviders, &proto_mail_provider.MailProvider{
			Email:      mp.Email,
			Password:   mp.Password,
			UserName:   mp.UserName,
			Port:       int32(mp.Port),
			Host:       mp.Host,
			Encryption: mp.Encryption,
			Name:       mp.Name,
			TypeId:     mp.TypeId,
			CreatedBy:  mp.CreatedBy,
			CreatedAt:  mp.CreatedAt.Format(time.RFC3339),
			UpdatedAt:  updatedAt,
		})
	}

	return &proto_mail_provider.GetAllMailProviderResponse{
		Message:       "Mail providers retrieved successfully",
		Total:         int32(len(mailProviders)),
		Page:          1,
		Limit:         int32(len(mailProviders)),
		MailProviders: mailProviders,
	}, nil
}

package grpcmailprovider

import (
	"context"
	proto "mail-service/proto/gen/mail_provider/v1"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (mp *mailProviderService) GetAllMailProvider(ctx context.Context, req *proto.GetAllMailProviderRequest) (*proto.GetAllMailProviderResponse, error) {
	result, err := mp.getAllMailProviderUsecase.Execute(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Lỗi lấy danh sách mail provider: %v", err)
	}

	// Convert to proto response
	var mailProviders []*proto.MailProvider
	for _, mp := range result {
		mailProviders = append(mailProviders, &proto.MailProvider{
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
			UpdatedAt:  mp.UpdatedAt.Format(time.RFC3339),
		})
	}

	return &proto.GetAllMailProviderResponse{
		Message:       "Mail providers retrieved successfully",
		Total:         int32(len(mailProviders)),
		Page:          1,
		Limit:         int32(len(mailProviders)),
		MailProviders: mailProviders,
	}, nil
}

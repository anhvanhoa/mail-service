package grpctypemail

import (
	"context"
	"mail-service/domain/entity"
	"time"

	proto_type_mail "github.com/anhvanhoa/sf-proto/gen/type_mail/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (tm *typeMailService) CreateTypeMail(ctx context.Context, req *proto_type_mail.CreateTypeMailRequest) (*proto_type_mail.CreateTypeMailResponse, error) {
	typeMail := entity.TypeMail{
		Name:      req.Name,
		CreatedBy: req.CreatedBy,
		CreatedAt: time.Now(),
	}

	if req.CreatedAt != "" {
		createdAt, err := time.Parse(time.RFC3339, req.CreatedAt)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "Thời gian tạo không hợp lệ")
		}
		typeMail.CreatedAt = createdAt
	}

	err := tm.createTypeMailUsecase.Execute(ctx, &typeMail)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Lỗi tạo type mail: %v", err)
	}

	return &proto_type_mail.CreateTypeMailResponse{
		Message: "Type mail created successfully",
		TypeMail: &proto_type_mail.TypeMail{
			Id:        typeMail.ID,
			Name:      typeMail.Name,
			CreatedBy: typeMail.CreatedBy,
			CreatedAt: typeMail.CreatedAt.Format(time.RFC3339),
			UpdatedAt: "",
		},
	}, nil
}

package grpctypemail

import (
	"context"
	"mail-service/domain/entity"
	"time"

	proto_type_mail "github.com/anhvanhoa/sf-proto/gen/type_mail/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (tm *typeMailService) UpdateTypeMail(ctx context.Context, req *proto_type_mail.UpdateTypeMailRequest) (*proto_type_mail.UpdateTypeMailResponse, error) {
	now := time.Now()
	typeMail := entity.TypeMail{
		ID:        req.Id,
		Name:      req.Name,
		UpdatedAt: &now,
	}

	err := tm.updateTypeMailUsecase.Execute(ctx, &typeMail)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Lỗi cập nhật type mail: %v", err)
	}

	updatedAt := ""
	if typeMail.UpdatedAt != nil {
		updatedAt = typeMail.UpdatedAt.Format(time.RFC3339)
	}

	return &proto_type_mail.UpdateTypeMailResponse{
		Message: "Type mail updated successfully",
		TypeMail: &proto_type_mail.TypeMail{
			Id:        typeMail.ID,
			Name:      typeMail.Name,
			CreatedBy: typeMail.CreatedBy,
			CreatedAt: typeMail.CreatedAt.Format(time.RFC3339),
			UpdatedAt: updatedAt,
		},
	}, nil
}

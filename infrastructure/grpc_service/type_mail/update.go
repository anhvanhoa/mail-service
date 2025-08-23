package grpcservice

import (
	"context"
	"mail-service/domain/entity"
	proto "mail-service/proto/gen/type_mail/v1"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (tm *typeMailService) UpdateTypeMail(ctx context.Context, req *proto.UpdateTypeMailRequest) (*proto.UpdateTypeMailResponse, error) {
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

	return &proto.UpdateTypeMailResponse{
		Message: "Type mail updated successfully",
		TypeMail: &proto.TypeMail{
			Id:        typeMail.ID,
			Name:      typeMail.Name,
			CreatedBy: typeMail.CreatedBy,
			CreatedAt: typeMail.CreatedAt.Format(time.RFC3339),
			UpdatedAt: updatedAt,
		},
	}, nil
}

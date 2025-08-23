package grpctypemail

import (
	"context"
	proto "mail-service/proto/gen/type_mail/v1"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (tm *typeMailService) GetTypeMail(ctx context.Context, req *proto.GetTypeMailRequest) (*proto.GetTypeMailResponse, error) {
	typeMail, err := tm.getTypeMailUsecase.Execute(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Type mail không tồn tại: %v", err)
	}

	updatedAt := ""
	if typeMail.UpdatedAt != nil {
		updatedAt = typeMail.UpdatedAt.Format(time.RFC3339)
	}

	return &proto.GetTypeMailResponse{
		Message: "Type mail retrieved successfully",
		TypeMail: &proto.TypeMail{
			Id:        typeMail.ID,
			Name:      typeMail.Name,
			CreatedBy: typeMail.CreatedBy,
			CreatedAt: typeMail.CreatedAt.Format(time.RFC3339),
			UpdatedAt: updatedAt,
		},
	}, nil
}

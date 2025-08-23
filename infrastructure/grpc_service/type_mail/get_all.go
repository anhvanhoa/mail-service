package grpcservice

import (
	"context"
	proto "mail-service/proto/gen/type_mail/v1"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (tm *typeMailService) GetAllTypeMail(ctx context.Context, req *proto.GetAllTypeMailRequest) (*proto.GetAllTypeMailResponse, error) {
	result, err := tm.getAllTypeMailUsecase.Execute(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Lỗi lấy danh sách type mail: %v", err)
	}

	// Convert to proto response
	var typeMails []*proto.TypeMail
	for _, tm := range result {
		updatedAt := ""
		if tm.UpdatedAt != nil {
			updatedAt = tm.UpdatedAt.Format(time.RFC3339)
		}

		typeMails = append(typeMails, &proto.TypeMail{
			Id:        tm.ID,
			Name:      tm.Name,
			CreatedBy: tm.CreatedBy,
			CreatedAt: tm.CreatedAt.Format(time.RFC3339),
			UpdatedAt: updatedAt,
		})
	}

	return &proto.GetAllTypeMailResponse{
		Message:   "Type mails retrieved successfully",
		Total:     int32(len(typeMails)),
		Page:      1,
		Limit:     int32(len(typeMails)),
		TypeMails: typeMails,
	}, nil
}

package grpcmailstatus

import (
	"context"
	"mail-service/domain/entity"
	"time"

	proto_mail_status "github.com/anhvanhoa/sf-proto/gen/mail_status/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (ms *mailStatusService) CreateMailStatus(ctx context.Context, req *proto_mail_status.CreateMailStatusRequest) (*proto_mail_status.CreateMailStatusResponse, error) {
	mailStatus := entity.MailStatus{
		Status:    entity.StatusMail(req.Status),
		Name:      req.Name,
		CreatedAt: time.Now(),
	}

	if req.CreatedAt != "" {
		createdAt, err := time.Parse(time.RFC3339, req.CreatedAt)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "Thời gian tạo không hợp lệ")
		}
		mailStatus.CreatedAt = createdAt
	}

	err := ms.createMailStatusUsecase.Execute(ctx, &mailStatus)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Lỗi tạo mail status: %v", err)
	}

	return &proto_mail_status.CreateMailStatusResponse{
		Message: "Mail status created successfully",
		MailStatus: &proto_mail_status.MailStatus{
			Status:    string(mailStatus.Status),
			Name:      mailStatus.Name,
			CreatedAt: mailStatus.CreatedAt.Format(time.RFC3339),
		},
	}, nil
}

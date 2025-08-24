package grpcmailstatus

import (
	"context"
	"mail-service/domain/entity"

	proto_mail_status "github.com/anhvanhoa/sf-proto/gen/mail_status/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (ms *mailStatusService) UpdateMailStatus(ctx context.Context, req *proto_mail_status.UpdateMailStatusRequest) (*proto_mail_status.UpdateMailStatusResponse, error) {
	mailStatus := entity.MailStatus{
		Status: entity.StatusMail(req.Status),
		Name:   req.Name,
	}

	err := ms.updateMailStatusUsecase.Execute(ctx, &mailStatus)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Lỗi cập nhật mail status: %v", err)
	}

	return &proto_mail_status.UpdateMailStatusResponse{
		Message: "Mail status updated successfully",
		MailStatus: &proto_mail_status.MailStatus{
			Status:    string(mailStatus.Status),
			Name:      mailStatus.Name,
			CreatedAt: mailStatus.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		},
	}, nil
}

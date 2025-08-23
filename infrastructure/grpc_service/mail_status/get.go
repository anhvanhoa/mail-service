package grpcmailstatus

import (
	"context"
	"mail-service/domain/entity"
	proto "mail-service/proto/gen/mail_status/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (ms *mailStatusService) GetMailStatus(ctx context.Context, req *proto.GetMailStatusRequest) (*proto.GetMailStatusResponse, error) {
	mailStatus, err := ms.getMailStatusUsecase.Execute(ctx, entity.StatusMail(req.Status))
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Mail status không tồn tại: %v", err)
	}

	return &proto.GetMailStatusResponse{
		Message: "Mail status retrieved successfully",
		MailStatus: &proto.MailStatus{
			Status:    string(mailStatus.Status),
			Name:      mailStatus.Name,
			CreatedAt: mailStatus.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		},
	}, nil
}

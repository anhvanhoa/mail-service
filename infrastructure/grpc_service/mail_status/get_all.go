package grpcmailstatus

import (
	"context"
	proto "mail-service/proto/gen/mail_status/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (ms *mailStatusService) GetAllMailStatus(ctx context.Context, req *proto.GetAllMailStatusRequest) (*proto.GetAllMailStatusResponse, error) {
	result, err := ms.getAllMailStatusUsecase.Execute(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Lỗi lấy danh sách mail status: %v", err)
	}

	var mailStatuses []*proto.MailStatus
	for _, ms := range result {
		mailStatuses = append(mailStatuses, &proto.MailStatus{
			Status:    string(ms.Status),
			Name:      ms.Name,
			CreatedAt: ms.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		})
	}

	return &proto.GetAllMailStatusResponse{
		Message:      "Mail statuses retrieved successfully",
		Total:        int32(len(mailStatuses)),
		Page:         1,
		Limit:        int32(len(mailStatuses)),
		MailStatuses: mailStatuses,
	}, nil
}

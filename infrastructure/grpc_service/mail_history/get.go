package grpcmailhistory

import (
	"context"
	"time"

	proto_mail_history "github.com/anhvanhoa/sf-proto/gen/mail_history/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (mh *mailHistoryService) GetMailHistory(ctx context.Context, req *proto_mail_history.GetMailHistoryRequest) (*proto_mail_history.GetMailHistoryResponse, error) {
	mailHistory, err := mh.getMailHistoryUsecase.Execute(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Mail history không tồn tại: %v", err)
	}

	updatedAt := ""
	if mailHistory.UpdatedAt != nil {
		updatedAt = mailHistory.UpdatedAt.Format(time.RFC3339)
	}

	return &proto_mail_history.GetMailHistoryResponse{
		Message: "Mail history retrieved successfully",
		MailHistory: &proto_mail_history.MailHistory{
			Id:            mailHistory.ID,
			TemplateId:    mailHistory.TemplateId,
			Subject:       mailHistory.Subject,
			Body:          mailHistory.Body,
			Tos:           mailHistory.Tos,
			Data:          mailHistory.Data,
			EmailProvider: mailHistory.EmailProvider,
			CreatedBy:     mailHistory.CreatedBy,
			CreatedAt:     mailHistory.CreatedAt.Format(time.RFC3339),
			UpdatedAt:     updatedAt,
		},
	}, nil
}

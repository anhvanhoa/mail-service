package grpcmailhistory

import (
	"context"
	"mail-service/domain/entity"
	"time"

	proto_mail_history "github.com/anhvanhoa/sf-proto/gen/mail_history/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (mh *mailHistoryService) UpdateMailHistory(ctx context.Context, req *proto_mail_history.UpdateMailHistoryRequest) (*proto_mail_history.UpdateMailHistoryResponse, error) {
	now := time.Now()
	mailHistory := entity.MailHistory{
		ID:            req.Id,
		TemplateId:    req.TemplateId,
		Subject:       req.Subject,
		Body:          req.Body,
		Tos:           req.Tos,
		Data:          make(map[string]any),
		EmailProvider: req.EmailProvider,
		UpdatedAt:     &now,
	}

	// Convert map[string]string to map[string]any
	for k, v := range req.Data {
		mailHistory.Data[k] = v
	}

	err := mh.updateMailHistoryUsecase.Execute(ctx, &mailHistory)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Lỗi cập nhật mail history: %v", err)
	}

	updatedAt := ""
	if mailHistory.UpdatedAt != nil {
		updatedAt = mailHistory.UpdatedAt.Format(time.RFC3339)
	}

	return &proto_mail_history.UpdateMailHistoryResponse{
		Message: "Mail history updated successfully",
		MailHistory: &proto_mail_history.MailHistory{
			Id:            mailHistory.ID,
			TemplateId:    mailHistory.TemplateId,
			Subject:       mailHistory.Subject,
			Body:          mailHistory.Body,
			Tos:           mailHistory.Tos,
			Data:          req.Data,
			EmailProvider: mailHistory.EmailProvider,
			CreatedBy:     mailHistory.CreatedBy,
			CreatedAt:     mailHistory.CreatedAt.Format(time.RFC3339),
			UpdatedAt:     updatedAt,
		},
	}, nil
}

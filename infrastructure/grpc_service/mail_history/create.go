package grpcmailhistory

import (
	"context"
	"mail-service/domain/entity"
	"time"

	proto_mail_history "github.com/anhvanhoa/sf-proto/gen/mail_history/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (mh *mailHistoryService) CreateMailHistory(ctx context.Context, req *proto_mail_history.CreateMailHistoryRequest) (*proto_mail_history.CreateMailHistoryResponse, error) {
	mailHistory := entity.MailHistory{
		ID:            req.Id,
		TemplateId:    req.TemplateId,
		Subject:       req.Subject,
		Body:          req.Body,
		Tos:           req.Tos,
		Data:          req.Data,
		EmailProvider: req.EmailProvider,
		CreatedBy:     req.CreatedBy,
		CreatedAt:     time.Now(),
	}

	if req.CreatedAt != "" {
		createdAt, err := time.Parse(time.RFC3339, req.CreatedAt)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "Thời gian tạo không hợp lệ")
		}
		mailHistory.CreatedAt = createdAt
	}

	err := mh.createMailHistoryUsecase.Execute(ctx, &mailHistory)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Lỗi tạo mail history: %v", err)
	}

	updatedAt := ""
	if mailHistory.UpdatedAt != nil {
		updatedAt = mailHistory.UpdatedAt.Format(time.RFC3339)
	}

	return &proto_mail_history.CreateMailHistoryResponse{
		Message: "Mail history created successfully",
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

package grpcservice

import (
	"context"
	"mail-service/domain/entity"
	proto "mail-service/proto/gen/mail_history/v1"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (mh *mailHistoryService) UpdateMailHistory(ctx context.Context, req *proto.UpdateMailHistoryRequest) (*proto.UpdateMailHistoryResponse, error) {
	now := time.Now()
	mailHistory := entity.MailHistory{
		ID:            req.Id,
		TemplateId:    req.TemplateId,
		Subject:       req.Subject,
		Body:          req.Body,
		To:            req.To,
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

	return &proto.UpdateMailHistoryResponse{
		Message: "Mail history updated successfully",
		MailHistory: &proto.MailHistory{
			Id:            mailHistory.ID,
			TemplateId:    mailHistory.TemplateId,
			Subject:       mailHistory.Subject,
			Body:          mailHistory.Body,
			To:            mailHistory.To,
			Tos:           mailHistory.Tos,
			Data:          req.Data, // Use original proto map
			EmailProvider: mailHistory.EmailProvider,
			CreatedBy:     mailHistory.CreatedBy,
			CreatedAt:     mailHistory.CreatedAt.Format(time.RFC3339),
			UpdatedAt:     updatedAt,
		},
	}, nil
}

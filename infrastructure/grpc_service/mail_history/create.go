package grpcservice

import (
	"context"
	"mail-service/domain/entity"
	proto "mail-service/proto/gen/mail_history/v1"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (mh *mailHistoryService) CreateMailHistory(ctx context.Context, req *proto.CreateMailHistoryRequest) (*proto.CreateMailHistoryResponse, error) {
	mailHistory := entity.MailHistory{
		TemplateId:    req.TemplateId,
		Subject:       req.Subject,
		Body:          req.Body,
		To:            req.To,
		Tos:           req.Tos,
		Data:          make(map[string]any),
		EmailProvider: req.EmailProvider,
		CreatedBy:     req.CreatedBy,
		CreatedAt:     time.Now(),
	}

	// Convert map[string]string to map[string]any
	for k, v := range req.Data {
		mailHistory.Data[k] = v
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

	return &proto.CreateMailHistoryResponse{
		Message: "Mail history created successfully",
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

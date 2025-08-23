package grpcservice

import (
	"context"
	proto "mail-service/proto/gen/mail_history/v1"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (mh *mailHistoryService) GetMailHistory(ctx context.Context, req *proto.GetMailHistoryRequest) (*proto.GetMailHistoryResponse, error) {
	mailHistory, err := mh.getMailHistoryUsecase.Execute(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Mail history không tồn tại: %v", err)
	}

	updatedAt := ""
	if mailHistory.UpdatedAt != nil {
		updatedAt = mailHistory.UpdatedAt.Format(time.RFC3339)
	}

	// Convert map[string]any to map[string]string for proto
	data := make(map[string]string)
	for k, v := range mailHistory.Data {
		if str, ok := v.(string); ok {
			data[k] = str
		}
	}

	return &proto.GetMailHistoryResponse{
		Message: "Mail history retrieved successfully",
		MailHistory: &proto.MailHistory{
			Id:            mailHistory.ID,
			TemplateId:    mailHistory.TemplateId,
			Subject:       mailHistory.Subject,
			Body:          mailHistory.Body,
			To:            mailHistory.To,
			Tos:           mailHistory.Tos,
			Data:          data,
			EmailProvider: mailHistory.EmailProvider,
			CreatedBy:     mailHistory.CreatedBy,
			CreatedAt:     mailHistory.CreatedAt.Format(time.RFC3339),
			UpdatedAt:     updatedAt,
		},
	}, nil
}

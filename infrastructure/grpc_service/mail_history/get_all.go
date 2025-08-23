package grpcmailhistory

import (
	"context"
	proto "mail-service/proto/gen/mail_history/v1"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (mh *mailHistoryService) GetAllMailHistory(ctx context.Context, req *proto.GetAllMailHistoryRequest) (*proto.GetAllMailHistoryResponse, error) {
	result, err := mh.getAllMailHistoryUsecase.Execute(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Lỗi lấy danh sách mail history: %v", err)
	}

	// Convert to proto response
	var mailHistories []*proto.MailHistory
	for _, mh := range result {
		updatedAt := ""
		if mh.UpdatedAt != nil {
			updatedAt = mh.UpdatedAt.Format(time.RFC3339)
		}

		// Convert map[string]any to map[string]string for proto
		data := make(map[string]string)
		for k, v := range mh.Data {
			if str, ok := v.(string); ok {
				data[k] = str
			}
		}

		mailHistories = append(mailHistories, &proto.MailHistory{
			Id:            mh.ID,
			TemplateId:    mh.TemplateId,
			Subject:       mh.Subject,
			Body:          mh.Body,
			Tos:           mh.Tos,
			Data:          data,
			EmailProvider: mh.EmailProvider,
			CreatedBy:     mh.CreatedBy,
			CreatedAt:     mh.CreatedAt.Format(time.RFC3339),
			UpdatedAt:     updatedAt,
		})
	}

	return &proto.GetAllMailHistoryResponse{
		Message:       "Mail histories retrieved successfully",
		Total:         int32(len(mailHistories)),
		Page:          1,
		Limit:         int32(len(mailHistories)),
		MailHistories: mailHistories,
	}, nil
}

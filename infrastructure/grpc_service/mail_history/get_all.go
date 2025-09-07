package grpcmailhistory

import (
	"context"
	"time"

	proto_mail_history "github.com/anhvanhoa/sf-proto/gen/mail_history/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (mh *mailHistoryService) GetAllMailHistory(ctx context.Context, req *proto_mail_history.GetAllMailHistoryRequest) (*proto_mail_history.GetAllMailHistoryResponse, error) {
	result, err := mh.getAllMailHistoryUsecase.Execute(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Lỗi lấy danh sách mail history: %v", err)
	}

	// Convert to proto response
	var mailHistories []*proto_mail_history.MailHistory
	for _, mh := range result {
		updatedAt := ""
		if mh.UpdatedAt != nil {
			updatedAt = mh.UpdatedAt.Format(time.RFC3339)
		}

		mailHistories = append(mailHistories, &proto_mail_history.MailHistory{
			Id:            mh.ID,
			TemplateId:    mh.TemplateId,
			Subject:       mh.Subject,
			Body:          mh.Body,
			Tos:           mh.Tos,
			Data:          mh.Data,
			EmailProvider: mh.EmailProvider,
			CreatedBy:     mh.CreatedBy,
			CreatedAt:     mh.CreatedAt.Format(time.RFC3339),
			UpdatedAt:     updatedAt,
		})
	}

	return &proto_mail_history.GetAllMailHistoryResponse{
		Message:       "Mail histories retrieved successfully",
		Total:         int32(len(mailHistories)),
		Page:          1,
		Limit:         int32(len(mailHistories)),
		MailHistories: mailHistories,
	}, nil
}

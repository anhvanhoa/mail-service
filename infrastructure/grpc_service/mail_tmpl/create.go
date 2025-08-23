package grpcservice

import (
	"context"
	"mail-service/domain/common"
	"mail-service/domain/entity"
	proto "mail-service/proto/gen/mail_tmpl/v1"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (mtmpl *mailTmplService) CreateMailTmpl(ctx context.Context, req *proto.CreateMailTmplRequest) (*proto.CreateMailTmplResponse, error) {
	mailTmpl := entity.MailTemplate{
		Subject:       req.Subject,
		Body:          req.Body,
		CreatedBy:     req.CreatedBy,
		CreatedAt:     time.Now(),
		Status:        common.Status(req.Status),
		ProviderEmail: req.ProviderEmail,
	}

	if req.CreatedAt != "" {
		createdAt, err := time.Parse(time.RFC3339, req.CreatedAt)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "Thời gian tạo không hợp lệ")
		}
		mailTmpl.CreatedAt = createdAt
	}

	mtmpl.createMailTmplUsecase.Execute(ctx, &mailTmpl)
	return &proto.CreateMailTmplResponse{
		Message: "Mail template created successfully",
		MailTmpl: &proto.MailTmpl{
			Id:        mailTmpl.ID,
			Subject:   mailTmpl.Subject,
			Body:      mailTmpl.Body,
			CreatedBy: mailTmpl.CreatedBy,
			CreatedAt: mailTmpl.CreatedAt.Format(time.RFC3339),
		},
	}, nil
}

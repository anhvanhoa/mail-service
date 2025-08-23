package grpcservice

import (
	"context"
	"mail-service/domain/common"
	"mail-service/domain/entity"
	proto "mail-service/proto/gen/mail_tmpl/v1"
	"time"
)

func (mtmpl *mailTmplService) UpdateMailTmpl(ctx context.Context, req *proto.UpdateMailTmplRequest) (*proto.UpdateMailTmplResponse, error) {
	mailTmpl := entity.MailTemplate{
		Subject:       req.Subject,
		Body:          req.Body,
		CreatedAt:     time.Now(),
		Status:        common.Status(req.Status),
		ProviderEmail: req.ProviderEmail,
	}
	mtmpl.updateMailTmplUsecase.Execute(ctx, &mailTmpl)

	return &proto.UpdateMailTmplResponse{
		Message: "Mail template updated successfully",
	}, nil
}

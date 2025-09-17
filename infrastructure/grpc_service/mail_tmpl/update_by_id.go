package grpcmailtmpl

import (
	"context"
	"mail-service/domain/common"
	"mail-service/domain/entity"

	proto_mail_tmpl "github.com/anhvanhoa/sf-proto/gen/mail_tmpl/v1"
)

func (mtmpl *mailTmplService) UpdateMailTmpl(ctx context.Context, req *proto_mail_tmpl.UpdateMailTmplRequest) (*proto_mail_tmpl.UpdateMailTmplResponse, error) {
	mailTmpl := entity.MailTemplate{
		ID:            req.Id,
		Name:          req.Name,
		Subject:       req.Subject,
		Body:          req.Body,
		Keys:          req.Keys,
		Status:        common.Status(req.Status),
		ProviderEmail: req.ProviderEmail,
	}
	err := mtmpl.updateMailTmplUsecase.Execute(ctx, &mailTmpl)
	if err != nil {
		return nil, err
	}

	return &proto_mail_tmpl.UpdateMailTmplResponse{
		Message: "Mail template updated successfully",
	}, nil
}

package grpcmailtmpl

import (
	"context"
	"mail-service/domain/common"
	"mail-service/domain/entity"
	"time"

	proto_mail_tmpl "github.com/anhvanhoa/sf-proto/gen/mail_tmpl/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (mtmpl *mailTmplService) CreateMailTmpl(ctx context.Context, req *proto_mail_tmpl.CreateMailTmplRequest) (*proto_mail_tmpl.CreateMailTmplResponse, error) {
	mailTmpl := mtmpl.createEntityMailTmpl(req)

	err := mtmpl.createMailTmplUsecase.Execute(ctx, mailTmpl)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Lỗi tạo mail template: %v", err)
	}

	return &proto_mail_tmpl.CreateMailTmplResponse{
		Message:  "Mail template created successfully",
		MailTmpl: mtmpl.createResponseMailTmpl(mailTmpl),
	}, nil
}

func (mtmpl *mailTmplService) createEntityMailTmpl(req *proto_mail_tmpl.CreateMailTmplRequest) *entity.MailTemplate {
	return &entity.MailTemplate{
		ID:            req.Id,
		Name:          req.Name,
		Subject:       req.Subject,
		Keys:          req.Keys,
		Body:          req.Body,
		CreatedBy:     req.CreatedBy,
		CreatedAt:     time.Now(),
		Status:        common.Status(req.Status),
		ProviderEmail: req.ProviderEmail,
	}
}

func (mtmpl *mailTmplService) createResponseMailTmpl(mailTmpl *entity.MailTemplate) *proto_mail_tmpl.MailTmpl {
	var updatedAt string
	if mailTmpl.UpdatedAt != nil {
		updatedAt = mailTmpl.UpdatedAt.Format(time.RFC3339)
	}
	return &proto_mail_tmpl.MailTmpl{
		Id:            mailTmpl.ID,
		Name:          mailTmpl.Name,
		Subject:       mailTmpl.Subject,
		Keys:          mailTmpl.Keys,
		Body:          mailTmpl.Body,
		ProviderEmail: mailTmpl.ProviderEmail,
		CreatedBy:     mailTmpl.CreatedBy,
		UpdatedAt:     updatedAt,
		CreatedAt:     mailTmpl.CreatedAt.Format(time.RFC3339),
	}
}

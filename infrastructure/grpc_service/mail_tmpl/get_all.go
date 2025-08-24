package grpcmailtmpl

import (
	"context"
	"time"

	proto_mail_tmpl "github.com/anhvanhoa/sf-proto/gen/mail_tmpl/v1"
)

func (mtmpl *mailTmplService) GetAllMailTmpl(ctx context.Context, req *proto_mail_tmpl.GetAllMailTmplRequest) (*proto_mail_tmpl.GetAllMailTmplResponse, error) {
	mailTmpls, err := mtmpl.getAllMailTmplUsecase.Execute(ctx)
	if err != nil {
		return nil, err
	}

	var mailTmplsProto []*proto_mail_tmpl.MailTmpl
	for _, mailTmpl := range mailTmpls {
		mailTmplsProto = append(mailTmplsProto, &proto_mail_tmpl.MailTmpl{
			Id:        mailTmpl.ID,
			Subject:   mailTmpl.Subject,
			Body:      mailTmpl.Body,
			CreatedBy: mailTmpl.CreatedBy,
			CreatedAt: mailTmpl.CreatedAt.Format(time.RFC3339),
		})
	}
	return &proto_mail_tmpl.GetAllMailTmplResponse{
		Message:   "Mail templates fetched successfully",
		Total:     int32(len(mailTmpls)),
		MailTmpls: mailTmplsProto,
	}, nil
}

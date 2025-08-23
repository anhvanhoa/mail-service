package grpcservice

import (
	"context"
	proto "mail-service/proto/gen/mail_tmpl/v1"
	"time"
)

func (mtmpl *mailTmplService) GetAllMailTmpl(ctx context.Context, req *proto.GetAllMailTmplRequest) (*proto.GetAllMailTmplResponse, error) {
	mailTmpls, err := mtmpl.getAllMailTmplUsecase.Execute(ctx)
	if err != nil {
		return nil, err
	}

	var mailTmplsProto []*proto.MailTmpl
	for _, mailTmpl := range mailTmpls {
		mailTmplsProto = append(mailTmplsProto, &proto.MailTmpl{
			Id:        mailTmpl.ID,
			Subject:   mailTmpl.Subject,
			Body:      mailTmpl.Body,
			CreatedBy: mailTmpl.CreatedBy,
			CreatedAt: mailTmpl.CreatedAt.Format(time.RFC3339),
		})
	}
	return &proto.GetAllMailTmplResponse{
		Message:   "Mail templates fetched successfully",
		Total:     int32(len(mailTmpls)),
		MailTmpls: mailTmplsProto,
	}, nil
}

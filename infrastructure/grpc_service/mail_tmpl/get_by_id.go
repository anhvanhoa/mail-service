package grpcmailtmpl

import (
	"context"
	"time"

	proto_mail_tmpl "github.com/anhvanhoa/sf-proto/gen/mail_tmpl/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (mtmpl *mailTmplService) GetMailTmpl(ctx context.Context, req *proto_mail_tmpl.GetMailTmplRequest) (*proto_mail_tmpl.GetMailTmplResponse, error) {
	mailTmpl, err := mtmpl.getMailTmplUsecase.Execute(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Lỗi lấy mail template: %v", err)
	}
	return &proto_mail_tmpl.GetMailTmplResponse{
		Message: "Mail template fetched successfully",
		MailTmpl: &proto_mail_tmpl.MailTmpl{
			Id:        mailTmpl.ID,
			Subject:   mailTmpl.Subject,
			Body:      mailTmpl.Body,
			CreatedBy: mailTmpl.CreatedBy,
			CreatedAt: mailTmpl.CreatedAt.Format(time.RFC3339),
		},
	}, nil
}

package grpcservice

import (
	"context"
	proto "mail-service/proto/gen/mail_tmpl/v1"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (mtmpl *mailTmplService) GetMailTmpl(ctx context.Context, req *proto.GetMailTmplRequest) (*proto.GetMailTmplResponse, error) {
	mailTmpl, err := mtmpl.getMailTmplUsecase.Execute(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Lỗi lấy mail template: %v", err)
	}
	return &proto.GetMailTmplResponse{
		Message: "Mail template fetched successfully",
		MailTmpl: &proto.MailTmpl{
			Id:        mailTmpl.ID,
			Subject:   mailTmpl.Subject,
			Body:      mailTmpl.Body,
			CreatedBy: mailTmpl.CreatedBy,
			CreatedAt: mailTmpl.CreatedAt.Format(time.RFC3339),
		},
	}, nil
}

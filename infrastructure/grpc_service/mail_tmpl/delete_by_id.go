package grpcmailtmpl

import (
	"context"

	proto_mail_tmpl "github.com/anhvanhoa/sf-proto/gen/mail_tmpl/v1"
)

func (mtmpl *mailTmplService) DeleteMailTmpl(ctx context.Context, req *proto_mail_tmpl.DeleteMailTmplRequest) (*proto_mail_tmpl.DeleteMailTmplResponse, error) {
	mtmpl.deleteMailTmplUsecase.Execute(ctx, req.Id)
	return &proto_mail_tmpl.DeleteMailTmplResponse{
		Message: "Mail template deleted successfully",
	}, nil
}

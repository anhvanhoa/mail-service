package grpcservice

import (
	"context"
	proto "mail-service/proto/gen/mail_tmpl/v1"
)

func (mtmpl *mailTmplService) DeleteMailTmpl(ctx context.Context, req *proto.DeleteMailTmplRequest) (*proto.DeleteMailTmplResponse, error) {
	mtmpl.deleteMailTmplUsecase.Execute(ctx, req.Id)
	return &proto.DeleteMailTmplResponse{
		Message: "Mail template deleted successfully",
	}, nil
}

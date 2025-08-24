package grpcmailtmpl

import (
	"mail-service/bootstrap"
	usecase "mail-service/domain/usecase/mail_tmpl"
	"mail-service/infrastructure/repo"

	proto_mail_tmpl "github.com/anhvanhoa/sf-proto/gen/mail_tmpl/v1"

	"github.com/go-pg/pg/v10"
)

type mailTmplService struct {
	proto_mail_tmpl.UnimplementedMailTmplServiceServer
	createMailTmplUsecase usecase.CreateMailTmplUsecase
	updateMailTmplUsecase usecase.UpdateByIdMailTmplUsecase
	deleteMailTmplUsecase usecase.DeleteByIdMailTmplUsecase
	getMailTmplUsecase    usecase.GetByIdMailTmplUsecase
	getAllMailTmplUsecase usecase.GetAllMailTmplUsecase
}

func NewMailTmplService(db *pg.DB, env *bootstrap.Env) proto_mail_tmpl.MailTmplServiceServer {
	mailTemplateRepository := repo.NewMailTemplateRepository(db)
	return &mailTmplService{
		createMailTmplUsecase: usecase.NewCreateMailTmplUsecase(mailTemplateRepository),
		updateMailTmplUsecase: usecase.NewUpdateByIdMailTmplUsecase(mailTemplateRepository),
		deleteMailTmplUsecase: usecase.NewDeleteByIdMailTmplUsecase(mailTemplateRepository),
		getMailTmplUsecase:    usecase.NewGetByIdMailTmplUsecase(mailTemplateRepository),
		getAllMailTmplUsecase: usecase.NewGetAllMailTmplUsecase(mailTemplateRepository),
	}
}

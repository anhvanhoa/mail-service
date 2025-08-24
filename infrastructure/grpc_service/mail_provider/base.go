package grpcmailprovider

import (
	"mail-service/bootstrap"
	usecase "mail-service/domain/usecase/mail_provider"
	"mail-service/infrastructure/repo"

	proto_mail_provider "github.com/anhvanhoa/sf-proto/gen/mail_provider/v1"

	"github.com/go-pg/pg/v10"
)

type mailProviderService struct {
	proto_mail_provider.UnsafeMailProviderServiceServer
	createMailProviderUsecase usecase.CreateMailProviderUsecase
	updateMailProviderUsecase usecase.UpdateByEmailMailProviderUsecase
	deleteMailProviderUsecase usecase.DeleteByEmailMailProviderUsecase
	getMailProviderUsecase    usecase.GetByEmailMailProviderUsecase
	getAllMailProviderUsecase usecase.GetAllMailProviderUsecase
}

func NewMailProviderService(db *pg.DB, env *bootstrap.Env) proto_mail_provider.MailProviderServiceServer {
	mailProviderRepository := repo.NewMailProviderRepository(db)
	return &mailProviderService{
		createMailProviderUsecase: usecase.NewCreateMailProviderUsecase(mailProviderRepository),
		updateMailProviderUsecase: usecase.NewUpdateByEmailMailProviderUsecase(mailProviderRepository),
		deleteMailProviderUsecase: usecase.NewDeleteByEmailMailProviderUsecase(mailProviderRepository),
		getMailProviderUsecase:    usecase.NewGetByEmailMailProviderUsecase(mailProviderRepository),
		getAllMailProviderUsecase: usecase.NewGetAllMailProviderUsecase(mailProviderRepository),
	}
}

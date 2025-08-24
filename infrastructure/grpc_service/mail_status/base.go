package grpcmailstatus

import (
	"mail-service/bootstrap"
	usecase "mail-service/domain/usecase/mail_status"
	"mail-service/infrastructure/repo"

	proto_mail_status "github.com/anhvanhoa/sf-proto/gen/mail_status/v1"

	"github.com/go-pg/pg/v10"
)

type mailStatusService struct {
	proto_mail_status.UnsafeMailStatusServiceServer
	createMailStatusUsecase usecase.CreateMailStatusUsecase
	updateMailStatusUsecase usecase.UpdateByStatusMailStatusUsecase
	deleteMailStatusUsecase usecase.DeleteByStatusMailStatusUsecase
	getMailStatusUsecase    usecase.GetByStatusMailStatusUsecase
	getAllMailStatusUsecase usecase.GetAllMailStatusUsecase
}

func NewMailStatusService(db *pg.DB, env *bootstrap.Env) proto_mail_status.MailStatusServiceServer {
	mailStatusRepository := repo.NewMailStatusRepository(db)
	return &mailStatusService{
		createMailStatusUsecase: usecase.NewCreateMailStatusUsecase(mailStatusRepository),
		updateMailStatusUsecase: usecase.NewUpdateByStatusMailStatusUsecase(mailStatusRepository),
		deleteMailStatusUsecase: usecase.NewDeleteByStatusMailStatusUsecase(mailStatusRepository),
		getMailStatusUsecase:    usecase.NewGetByStatusMailStatusUsecase(mailStatusRepository),
		getAllMailStatusUsecase: usecase.NewGetAllMailStatusUsecase(mailStatusRepository),
	}
}

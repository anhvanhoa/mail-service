package grpcmailstatus

import (
	usecase "mail-service/domain/usecase/mail_status"
	"mail-service/infrastructure/repo"

	proto_mail_status "github.com/anhvanhoa/sf-proto/gen/mail_status/v1"

	"github.com/go-pg/pg/v10"
)

type mailStatusService struct {
	proto_mail_status.UnsafeMailStatusServiceServer
	getMailStatusUsecase    usecase.GetByStatusMailStatusUsecase
	getAllMailStatusUsecase usecase.GetAllMailStatusUsecase
}

func NewMailStatusService(db *pg.DB) proto_mail_status.MailStatusServiceServer {
	mailStatusRepository := repo.NewMailStatusRepository(db)
	return &mailStatusService{
		getMailStatusUsecase:    usecase.NewGetByStatusMailStatusUsecase(mailStatusRepository),
		getAllMailStatusUsecase: usecase.NewGetAllMailStatusUsecase(mailStatusRepository),
	}
}

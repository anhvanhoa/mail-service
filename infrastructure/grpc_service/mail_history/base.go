package grpcservice

import (
	"mail-service/bootstrap"
	usecase "mail-service/domain/usecase/mail_history"
	"mail-service/infrastructure/repo"
	proto "mail-service/proto/gen/mail_history/v1"

	"github.com/go-pg/pg/v10"
)

type mailHistoryService struct {
	proto.UnsafeMailHistoryServiceServer
	createMailHistoryUsecase usecase.CreateMailHistoryUsecase
	updateMailHistoryUsecase usecase.UpdateByIdMailHistoryUsecase
	deleteMailHistoryUsecase usecase.DeleteByIdMailHistoryUsecase
	getMailHistoryUsecase    usecase.GetByIdMailHistoryUsecase
	getAllMailHistoryUsecase usecase.GetAllMailHistoryUsecase
}

func NewMailHistoryService(db *pg.DB, env *bootstrap.Env) proto.MailHistoryServiceServer {
	mailHistoryRepository := repo.NewMailHistoryRepository(db)
	return &mailHistoryService{
		createMailHistoryUsecase: usecase.NewCreateMailHistoryUsecase(mailHistoryRepository),
		updateMailHistoryUsecase: usecase.NewUpdateByIdMailHistoryUsecase(mailHistoryRepository),
		deleteMailHistoryUsecase: usecase.NewDeleteByIdMailHistoryUsecase(mailHistoryRepository),
		getMailHistoryUsecase:    usecase.NewGetByIdMailHistoryUsecase(mailHistoryRepository),
		getAllMailHistoryUsecase: usecase.NewGetAllMailHistoryUsecase(mailHistoryRepository),
	}
}

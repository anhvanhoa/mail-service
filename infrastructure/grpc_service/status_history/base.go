package grpcservice

import (
	"mail-service/bootstrap"
	usecase "mail-service/domain/usecase/status_history"
	"mail-service/infrastructure/repo"
	proto "mail-service/proto/gen/status_history/v1"

	"github.com/go-pg/pg/v10"
)

type statusHistoryService struct {
	proto.UnsafeStatusHistoryServiceServer
	createStatusHistoryUsecase                   usecase.CreateStatusHistoryUsecase
	updateStatusHistoryUsecase                   usecase.UpdateStatusHistoryUsecase
	deleteStatusHistoryUsecase                   usecase.DeleteStatusHistoryUsecase
	getAllStatusHistoryUsecase                   usecase.GetAllStatusHistoryUsecase
	getStatusHistoryByMailHistoryIdUsecase       usecase.GetByMailHistoryIdStatusHistoryUsecase
	getLatestStatusHistoryByMailHistoryIdUsecase usecase.GetLatestByMailHistoryIdStatusHistoryUsecase
}

func NewStatusHistoryService(db *pg.DB, env *bootstrap.Env) proto.StatusHistoryServiceServer {
	statusHistoryRepository := repo.NewStatusHistoryRepository(db)
	return &statusHistoryService{
		createStatusHistoryUsecase:                   usecase.NewCreateStatusHistoryUsecase(statusHistoryRepository),
		updateStatusHistoryUsecase:                   usecase.NewUpdateStatusHistoryUsecase(statusHistoryRepository),
		deleteStatusHistoryUsecase:                   usecase.NewDeleteStatusHistoryUsecase(statusHistoryRepository),
		getAllStatusHistoryUsecase:                   usecase.NewGetAllStatusHistoryUsecase(statusHistoryRepository),
		getStatusHistoryByMailHistoryIdUsecase:       usecase.NewGetByMailHistoryIdStatusHistoryUsecase(statusHistoryRepository),
		getLatestStatusHistoryByMailHistoryIdUsecase: usecase.NewGetLatestByMailHistoryIdStatusHistoryUsecase(statusHistoryRepository),
	}
}

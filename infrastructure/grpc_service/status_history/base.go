package grpcstatushistory

import (
	usecase "mail-service/domain/usecase/status_history"
	"mail-service/infrastructure/repo"

	proto_status_history "github.com/anhvanhoa/sf-proto/gen/status_history/v1"

	"github.com/go-pg/pg/v10"
)

type statusHistoryService struct {
	proto_status_history.UnsafeStatusHistoryServiceServer
	createStatusHistoryUsecase                   usecase.CreateStatusHistoryUsecase
	updateStatusHistoryUsecase                   usecase.UpdateStatusHistoryUsecase
	deleteStatusHistoryUsecase                   usecase.DeleteStatusHistoryUsecase
	getAllStatusHistoryUsecase                   usecase.GetAllStatusHistoryUsecase
	getStatusHistoryByMailHistoryIdUsecase       usecase.GetByMailHistoryIdStatusHistoryUsecase
	getLatestStatusHistoryByMailHistoryIdUsecase usecase.GetLatestByMailHistoryIdStatusHistoryUsecase
}

func NewStatusHistoryService(db *pg.DB) proto_status_history.StatusHistoryServiceServer {
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

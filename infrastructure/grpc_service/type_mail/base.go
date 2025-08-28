package grpctypemail

import (
	usecase "mail-service/domain/usecase/type_mail"
	"mail-service/infrastructure/repo"

	proto "github.com/anhvanhoa/sf-proto/gen/type_mail/v1"

	"github.com/go-pg/pg/v10"
)

type typeMailService struct {
	proto.UnsafeTypeMailServiceServer
	createTypeMailUsecase usecase.CreateTypeMailUsecase
	updateTypeMailUsecase usecase.UpdateByIdTypeMailUsecase
	deleteTypeMailUsecase usecase.DeleteByIdTypeMailUsecase
	getTypeMailUsecase    usecase.GetByIdTypeMailUsecase
	getAllTypeMailUsecase usecase.GetAllTypeMailUsecase
}

func NewTypeMailService(db *pg.DB) proto.TypeMailServiceServer {
	typeMailRepository := repo.NewTypeMailRepository(db)
	return &typeMailService{
		createTypeMailUsecase: usecase.NewCreateTypeMailUsecase(typeMailRepository),
		updateTypeMailUsecase: usecase.NewUpdateByIdTypeMailUsecase(typeMailRepository),
		deleteTypeMailUsecase: usecase.NewDeleteByIdTypeMailUsecase(typeMailRepository),
		getTypeMailUsecase:    usecase.NewGetByIdTypeMailUsecase(typeMailRepository),
		getAllTypeMailUsecase: usecase.NewGetAllTypeMailUsecase(typeMailRepository),
	}
}

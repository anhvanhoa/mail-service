package usecase

import (
	"context"
	"mail-service/domain/entity"
	"mail-service/domain/repository"
)

type GetAllTypeMailUsecase interface {
	Execute(ctx context.Context) ([]*entity.TypeMail, error)
}

type getAllTypeMailUsecase struct {
	typeMailRepository repository.TypeMailRepository
}

func NewGetAllTypeMailUsecase(typeMailRepository repository.TypeMailRepository) GetAllTypeMailUsecase {
	return &getAllTypeMailUsecase{
		typeMailRepository: typeMailRepository,
	}
}

func (u *getAllTypeMailUsecase) Execute(ctx context.Context) ([]*entity.TypeMail, error) {
	return u.typeMailRepository.GetAll(ctx)
}

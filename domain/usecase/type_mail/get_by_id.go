package usecase

import (
	"context"
	"mail-service/domain/entity"
	"mail-service/domain/repository"
)

type GetByIdTypeMailUsecase interface {
	Execute(ctx context.Context, id string) (*entity.TypeMail, error)
}

type getByIdTypeMailUsecase struct {
	typeMailRepository repository.TypeMailRepository
}

func NewGetByIdTypeMailUsecase(typeMailRepository repository.TypeMailRepository) GetByIdTypeMailUsecase {
	return &getByIdTypeMailUsecase{
		typeMailRepository: typeMailRepository,
	}
}

func (u *getByIdTypeMailUsecase) Execute(ctx context.Context, id string) (*entity.TypeMail, error) {
	return u.typeMailRepository.GetByID(ctx, id)
}

package usecase

import (
	"context"
	"mail-service/domain/repository"
)

type DeleteByIdTypeMailUsecase interface {
	Execute(ctx context.Context, id string) error
}

type deleteByIdTypeMailUsecase struct {
	typeMailRepository repository.TypeMailRepository
}

func NewDeleteByIdTypeMailUsecase(typeMailRepository repository.TypeMailRepository) DeleteByIdTypeMailUsecase {
	return &deleteByIdTypeMailUsecase{
		typeMailRepository: typeMailRepository,
	}
}

func (u *deleteByIdTypeMailUsecase) Execute(ctx context.Context, id string) error {
	return u.typeMailRepository.Delete(ctx, id)
}

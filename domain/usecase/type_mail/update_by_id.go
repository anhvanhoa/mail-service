package usecase

import (
	"context"
	"mail-service/domain/entity"
	"mail-service/domain/repository"
)

type UpdateByIdTypeMailUsecase interface {
	Execute(ctx context.Context, req *entity.TypeMail) error
}

type updateByIdTypeMailUsecase struct {
	typeMailRepository repository.TypeMailRepository
}

func NewUpdateByIdTypeMailUsecase(typeMailRepository repository.TypeMailRepository) UpdateByIdTypeMailUsecase {
	return &updateByIdTypeMailUsecase{
		typeMailRepository: typeMailRepository,
	}
}

func (u *updateByIdTypeMailUsecase) Execute(ctx context.Context, req *entity.TypeMail) error {
	return u.typeMailRepository.Update(ctx, req)
}

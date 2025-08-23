package usecase

import (
	"context"
	"mail-service/domain/entity"
	"mail-service/domain/repository"
)

type CreateTypeMailUsecase interface {
	Execute(ctx context.Context, req *entity.TypeMail) error
}

type createTypeMailUsecase struct {
	typeMailRepository repository.TypeMailRepository
}

func NewCreateTypeMailUsecase(typeMailRepository repository.TypeMailRepository) CreateTypeMailUsecase {
	return &createTypeMailUsecase{
		typeMailRepository: typeMailRepository,
	}
}

func (u *createTypeMailUsecase) Execute(ctx context.Context, req *entity.TypeMail) error {
	return u.typeMailRepository.Create(ctx, req)
}

package usecase

import (
	"context"
	"mail-service/domain/entity"
	"mail-service/domain/repository"
)

type GetByEmailMailProviderUsecase interface {
	Execute(ctx context.Context, email string) (*entity.MailProvider, error)
}

type getByEmailMailProviderUsecase struct {
	mailProviderRepository repository.MailProviderRepository
}

func NewGetByEmailMailProviderUsecase(mailProviderRepository repository.MailProviderRepository) GetByEmailMailProviderUsecase {
	return &getByEmailMailProviderUsecase{
		mailProviderRepository: mailProviderRepository,
	}
}

func (u *getByEmailMailProviderUsecase) Execute(ctx context.Context, email string) (*entity.MailProvider, error) {
	return u.mailProviderRepository.GetByEmail(ctx, email)
}

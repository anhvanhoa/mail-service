package usecase

import (
	"context"
	"mail-service/domain/repository"
)

type DeleteByEmailMailProviderUsecase interface {
	Execute(ctx context.Context, email string) error
}

type deleteByEmailMailProviderUsecase struct {
	mailProviderRepository repository.MailProviderRepository
}

func NewDeleteByEmailMailProviderUsecase(mailProviderRepository repository.MailProviderRepository) DeleteByEmailMailProviderUsecase {
	return &deleteByEmailMailProviderUsecase{
		mailProviderRepository: mailProviderRepository,
	}
}

func (u *deleteByEmailMailProviderUsecase) Execute(ctx context.Context, email string) error {
	return u.mailProviderRepository.Delete(ctx, email)
}

package usecase

import (
	"context"
	"mail-service/domain/entity"
	"mail-service/domain/repository"
)

type GetAllMailProviderUsecase interface {
	Execute(ctx context.Context) ([]*entity.MailProvider, error)
}

type getAllMailProviderUsecase struct {
	mailProviderRepository repository.MailProviderRepository
}

func NewGetAllMailProviderUsecase(mailProviderRepository repository.MailProviderRepository) GetAllMailProviderUsecase {
	return &getAllMailProviderUsecase{
		mailProviderRepository: mailProviderRepository,
	}
}

func (u *getAllMailProviderUsecase) Execute(ctx context.Context) ([]*entity.MailProvider, error) {
	return u.mailProviderRepository.GetAll(ctx)
}

package usecase

import (
	"context"
	"mail-service/domain/entity"
	"mail-service/domain/repository"
)

type CreateMailProviderUsecase interface {
	Execute(ctx context.Context, req *entity.MailProvider) error
}

type createMailProviderUsecase struct {
	mailProviderRepository repository.MailProviderRepository
}

func NewCreateMailProviderUsecase(mailProviderRepository repository.MailProviderRepository) CreateMailProviderUsecase {
	return &createMailProviderUsecase{
		mailProviderRepository: mailProviderRepository,
	}
}

func (u *createMailProviderUsecase) Execute(ctx context.Context, req *entity.MailProvider) error {
	return u.mailProviderRepository.Create(ctx, req)
}

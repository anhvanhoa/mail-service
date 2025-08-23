package usecase

import (
	"context"
	"mail-service/domain/entity"
	"mail-service/domain/repository"
)

type UpdateByEmailMailProviderUsecase interface {
	Execute(ctx context.Context, req *entity.MailProvider) error
}

type updateByEmailMailProviderUsecase struct {
	mailProviderRepository repository.MailProviderRepository
}

func NewUpdateByEmailMailProviderUsecase(mailProviderRepository repository.MailProviderRepository) UpdateByEmailMailProviderUsecase {
	return &updateByEmailMailProviderUsecase{
		mailProviderRepository: mailProviderRepository,
	}
}

func (u *updateByEmailMailProviderUsecase) Execute(ctx context.Context, req *entity.MailProvider) error {
	return u.mailProviderRepository.Update(ctx, req)
}

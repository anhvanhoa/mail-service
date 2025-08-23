package usecase

import (
	"context"
	"mail-service/domain/entity"
	"mail-service/domain/repository"
)

type CreateMailStatusUsecase interface {
	Execute(ctx context.Context, req *entity.MailStatus) error
}

type createMailStatusUsecase struct {
	mailStatusRepository repository.MailStatusRepository
}

func NewCreateMailStatusUsecase(mailStatusRepository repository.MailStatusRepository) CreateMailStatusUsecase {
	return &createMailStatusUsecase{
		mailStatusRepository: mailStatusRepository,
	}
}

func (u *createMailStatusUsecase) Execute(ctx context.Context, req *entity.MailStatus) error {
	return u.mailStatusRepository.Create(ctx, req)
}

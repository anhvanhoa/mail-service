package usecase

import (
	"context"
	"mail-service/domain/entity"
	"mail-service/domain/repository"
)

type DeleteByStatusMailStatusUsecase interface {
	Execute(ctx context.Context, status entity.StatusMail) error
}

type deleteByStatusMailStatusUsecase struct {
	mailStatusRepository repository.MailStatusRepository
}

func NewDeleteByStatusMailStatusUsecase(mailStatusRepository repository.MailStatusRepository) DeleteByStatusMailStatusUsecase {
	return &deleteByStatusMailStatusUsecase{
		mailStatusRepository: mailStatusRepository,
	}
}

func (u *deleteByStatusMailStatusUsecase) Execute(ctx context.Context, status entity.StatusMail) error {
	return u.mailStatusRepository.Delete(ctx, status)
}

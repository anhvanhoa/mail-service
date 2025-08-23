package usecase

import (
	"context"
	"mail-service/domain/entity"
	"mail-service/domain/repository"
)

type GetByStatusMailStatusUsecase interface {
	Execute(ctx context.Context, status entity.StatusMail) (*entity.MailStatus, error)
}

type getByStatusMailStatusUsecase struct {
	mailStatusRepository repository.MailStatusRepository
}

func NewGetByStatusMailStatusUsecase(mailStatusRepository repository.MailStatusRepository) GetByStatusMailStatusUsecase {
	return &getByStatusMailStatusUsecase{
		mailStatusRepository: mailStatusRepository,
	}
}

func (u *getByStatusMailStatusUsecase) Execute(ctx context.Context, status entity.StatusMail) (*entity.MailStatus, error) {
	return u.mailStatusRepository.GetByStatus(ctx, status)
}

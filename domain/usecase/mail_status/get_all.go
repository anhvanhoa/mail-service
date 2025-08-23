package usecase

import (
	"context"
	"mail-service/domain/entity"
	"mail-service/domain/repository"
)

type GetAllMailStatusUsecase interface {
	Execute(ctx context.Context) ([]*entity.MailStatus, error)
}

type getAllMailStatusUsecase struct {
	mailStatusRepository repository.MailStatusRepository
}

func NewGetAllMailStatusUsecase(mailStatusRepository repository.MailStatusRepository) GetAllMailStatusUsecase {
	return &getAllMailStatusUsecase{
		mailStatusRepository: mailStatusRepository,
	}
}

func (u *getAllMailStatusUsecase) Execute(ctx context.Context) ([]*entity.MailStatus, error) {
	return u.mailStatusRepository.GetAll(ctx)
}

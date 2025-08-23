package usecase

import (
	"context"
	"mail-service/domain/entity"
	"mail-service/domain/repository"
)

type GetAllMailHistoryUsecase interface {
	Execute(ctx context.Context) ([]*entity.MailHistory, error)
}

type getAllMailHistoryUsecase struct {
	mailHistoryRepository repository.MailHistoryRepository
}

func NewGetAllMailHistoryUsecase(mailHistoryRepository repository.MailHistoryRepository) GetAllMailHistoryUsecase {
	return &getAllMailHistoryUsecase{
		mailHistoryRepository: mailHistoryRepository,
	}
}

func (u *getAllMailHistoryUsecase) Execute(ctx context.Context) ([]*entity.MailHistory, error) {
	return u.mailHistoryRepository.GetAll(ctx)
}

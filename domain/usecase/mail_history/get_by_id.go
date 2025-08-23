package usecase

import (
	"context"
	"mail-service/domain/entity"
	"mail-service/domain/repository"
)

type GetByIdMailHistoryUsecase interface {
	Execute(ctx context.Context, id string) (*entity.MailHistory, error)
}

type getByIdMailHistoryUsecase struct {
	mailHistoryRepository repository.MailHistoryRepository
}

func NewGetByIdMailHistoryUsecase(mailHistoryRepository repository.MailHistoryRepository) GetByIdMailHistoryUsecase {
	return &getByIdMailHistoryUsecase{
		mailHistoryRepository: mailHistoryRepository,
	}
}

func (u *getByIdMailHistoryUsecase) Execute(ctx context.Context, id string) (*entity.MailHistory, error) {
	return u.mailHistoryRepository.GetByID(ctx, id)
}

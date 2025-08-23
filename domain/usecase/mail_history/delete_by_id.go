package usecase

import (
	"context"
	"mail-service/domain/repository"
)

type DeleteByIdMailHistoryUsecase interface {
	Execute(ctx context.Context, id string) error
}

type deleteByIdMailHistoryUsecase struct {
	mailHistoryRepository repository.MailHistoryRepository
}

func NewDeleteByIdMailHistoryUsecase(mailHistoryRepository repository.MailHistoryRepository) DeleteByIdMailHistoryUsecase {
	return &deleteByIdMailHistoryUsecase{
		mailHistoryRepository: mailHistoryRepository,
	}
}

func (u *deleteByIdMailHistoryUsecase) Execute(ctx context.Context, id string) error {
	return u.mailHistoryRepository.Delete(ctx, id)
}

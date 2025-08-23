package usecase

import (
	"context"
	"mail-service/domain/entity"
	"mail-service/domain/repository"
)

type DeleteStatusHistoryUsecase interface {
	Execute(ctx context.Context, mailHistoryID string, status entity.StatusMail) error
}

type deleteStatusHistoryUsecase struct {
	statusHistoryRepository repository.StatusHistoryRepository
}

func NewDeleteStatusHistoryUsecase(statusHistoryRepository repository.StatusHistoryRepository) DeleteStatusHistoryUsecase {
	return &deleteStatusHistoryUsecase{
		statusHistoryRepository: statusHistoryRepository,
	}
}

func (u *deleteStatusHistoryUsecase) Execute(ctx context.Context, mailHistoryID string, status entity.StatusMail) error {
	return u.statusHistoryRepository.Delete(ctx, mailHistoryID, status)
}

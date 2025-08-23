package usecase

import (
	"context"
	"mail-service/domain/entity"
	"mail-service/domain/repository"
)

type GetByMailHistoryIdStatusHistoryUsecase interface {
	Execute(ctx context.Context, mailHistoryID string) ([]*entity.StatusHistory, error)
}

type getByMailHistoryIdStatusHistoryUsecase struct {
	statusHistoryRepository repository.StatusHistoryRepository
}

func NewGetByMailHistoryIdStatusHistoryUsecase(statusHistoryRepository repository.StatusHistoryRepository) GetByMailHistoryIdStatusHistoryUsecase {
	return &getByMailHistoryIdStatusHistoryUsecase{
		statusHistoryRepository: statusHistoryRepository,
	}
}

func (u *getByMailHistoryIdStatusHistoryUsecase) Execute(ctx context.Context, mailHistoryID string) ([]*entity.StatusHistory, error) {
	return u.statusHistoryRepository.GetByMailHistoryID(ctx, mailHistoryID)
}

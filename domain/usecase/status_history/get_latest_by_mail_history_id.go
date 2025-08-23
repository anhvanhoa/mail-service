package usecase

import (
	"context"
	"mail-service/domain/entity"
	"mail-service/domain/repository"
)

type GetLatestByMailHistoryIdStatusHistoryUsecase interface {
	Execute(ctx context.Context, mailHistoryID string) (*entity.StatusHistory, error)
}

type getLatestByMailHistoryIdStatusHistoryUsecase struct {
	statusHistoryRepository repository.StatusHistoryRepository
}

func NewGetLatestByMailHistoryIdStatusHistoryUsecase(statusHistoryRepository repository.StatusHistoryRepository) GetLatestByMailHistoryIdStatusHistoryUsecase {
	return &getLatestByMailHistoryIdStatusHistoryUsecase{
		statusHistoryRepository: statusHistoryRepository,
	}
}

func (u *getLatestByMailHistoryIdStatusHistoryUsecase) Execute(ctx context.Context, mailHistoryID string) (*entity.StatusHistory, error) {
	return u.statusHistoryRepository.GetLatestByMailHistoryID(ctx, mailHistoryID)
}

package usecase

import (
	"context"
	"mail-service/domain/entity"
	"mail-service/domain/repository"
)

type CreateStatusHistoryUsecase interface {
	Execute(ctx context.Context, req *entity.StatusHistory) error
}

type createStatusHistoryUsecase struct {
	statusHistoryRepository repository.StatusHistoryRepository
}

func NewCreateStatusHistoryUsecase(statusHistoryRepository repository.StatusHistoryRepository) CreateStatusHistoryUsecase {
	return &createStatusHistoryUsecase{
		statusHistoryRepository: statusHistoryRepository,
	}
}

func (u *createStatusHistoryUsecase) Execute(ctx context.Context, req *entity.StatusHistory) error {
	return u.statusHistoryRepository.Create(ctx, req)
}

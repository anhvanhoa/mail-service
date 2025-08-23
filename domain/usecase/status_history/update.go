package usecase

import (
	"context"
	"mail-service/domain/entity"
	"mail-service/domain/repository"
)

type UpdateStatusHistoryUsecase interface {
	Execute(ctx context.Context, req *entity.StatusHistory) error
}

type updateStatusHistoryUsecase struct {
	statusHistoryRepository repository.StatusHistoryRepository
}

func NewUpdateStatusHistoryUsecase(statusHistoryRepository repository.StatusHistoryRepository) UpdateStatusHistoryUsecase {
	return &updateStatusHistoryUsecase{
		statusHistoryRepository: statusHistoryRepository,
	}
}

func (u *updateStatusHistoryUsecase) Execute(ctx context.Context, req *entity.StatusHistory) error {
	return u.statusHistoryRepository.Update(ctx, req)
}

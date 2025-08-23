package usecase

import (
	"context"
	"mail-service/domain/entity"
	"mail-service/domain/repository"
)

type GetAllStatusHistoryUsecase interface {
	Execute(ctx context.Context) ([]*entity.StatusHistory, error)
}

type getAllStatusHistoryUsecase struct {
	statusHistoryRepository repository.StatusHistoryRepository
}

func NewGetAllStatusHistoryUsecase(statusHistoryRepository repository.StatusHistoryRepository) GetAllStatusHistoryUsecase {
	return &getAllStatusHistoryUsecase{
		statusHistoryRepository: statusHistoryRepository,
	}
}

func (u *getAllStatusHistoryUsecase) Execute(ctx context.Context) ([]*entity.StatusHistory, error) {
	return u.statusHistoryRepository.GetAll(ctx)
}

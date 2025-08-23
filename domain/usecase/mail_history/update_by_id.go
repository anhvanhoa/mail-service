package usecase

import (
	"context"
	"mail-service/domain/entity"
	"mail-service/domain/repository"
)

type UpdateByIdMailHistoryUsecase interface {
	Execute(ctx context.Context, req *entity.MailHistory) error
}

type updateByIdMailHistoryUsecase struct {
	mailHistoryRepository repository.MailHistoryRepository
}

func NewUpdateByIdMailHistoryUsecase(mailHistoryRepository repository.MailHistoryRepository) UpdateByIdMailHistoryUsecase {
	return &updateByIdMailHistoryUsecase{
		mailHistoryRepository: mailHistoryRepository,
	}
}

func (u *updateByIdMailHistoryUsecase) Execute(ctx context.Context, req *entity.MailHistory) error {
	return u.mailHistoryRepository.Update(ctx, req)
}

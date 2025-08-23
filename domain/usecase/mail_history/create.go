package usecase

import (
	"context"
	"mail-service/domain/entity"
	"mail-service/domain/repository"
)

type CreateMailHistoryUsecase interface {
	Execute(ctx context.Context, req *entity.MailHistory) error
}

type createMailHistoryUsecase struct {
	mailHistoryRepository repository.MailHistoryRepository
}

func NewCreateMailHistoryUsecase(mailHistoryRepository repository.MailHistoryRepository) CreateMailHistoryUsecase {
	return &createMailHistoryUsecase{
		mailHistoryRepository: mailHistoryRepository,
	}
}

func (u *createMailHistoryUsecase) Execute(ctx context.Context, req *entity.MailHistory) error {
	return u.mailHistoryRepository.Create(ctx, req)
}

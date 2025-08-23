package usecase

import (
	"context"
	"mail-service/domain/entity"
	"mail-service/domain/repository"
)

type UpdateByStatusMailStatusUsecase interface {
	Execute(ctx context.Context, req *entity.MailStatus) error
}

type updateByStatusMailStatusUsecase struct {
	mailStatusRepository repository.MailStatusRepository
}

func NewUpdateByStatusMailStatusUsecase(mailStatusRepository repository.MailStatusRepository) UpdateByStatusMailStatusUsecase {
	return &updateByStatusMailStatusUsecase{
		mailStatusRepository: mailStatusRepository,
	}
}

func (u *updateByStatusMailStatusUsecase) Execute(ctx context.Context, req *entity.MailStatus) error {
	return u.mailStatusRepository.Update(ctx, req)
}

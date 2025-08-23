package repository

import (
	"context"

	"mail-service/domain/entity"
)

type MailStatusRepository interface {
	Create(ctx context.Context, status *entity.MailStatus) error
	GetByStatus(ctx context.Context, status entity.StatusMail) (*entity.MailStatus, error)
	GetByName(ctx context.Context, name string) (*entity.MailStatus, error)
	GetAll(ctx context.Context) ([]*entity.MailStatus, error)
	Update(ctx context.Context, status *entity.MailStatus) error
	Delete(ctx context.Context, status entity.StatusMail) error
}

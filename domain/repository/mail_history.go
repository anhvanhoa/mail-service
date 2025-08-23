package repository

import (
	"context"

	"mail-service/domain/entity"
)

type MailHistoryRepository interface {
	Create(ctx context.Context, history *entity.MailHistory) error
	GetByID(ctx context.Context, id string) (*entity.MailHistory, error)
	GetByTemplateID(ctx context.Context, templateID string) ([]*entity.MailHistory, error)
	GetByEmailProvider(ctx context.Context, emailProvider string) ([]*entity.MailHistory, error)
	GetByCreatedBy(ctx context.Context, createdBy string) ([]*entity.MailHistory, error)
	GetByTo(ctx context.Context, to string) ([]*entity.MailHistory, error)
	GetAll(ctx context.Context) ([]*entity.MailHistory, error)
	Update(ctx context.Context, history *entity.MailHistory) error
	Delete(ctx context.Context, id string) error
}

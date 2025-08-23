package repository

import (
	"context"

	"mail-service/domain/entity"
)

type StatusHistoryRepository interface {
	Create(ctx context.Context, history *entity.StatusHistory) error
	GetByMailHistoryID(ctx context.Context, mailHistoryID string) ([]*entity.StatusHistory, error)
	GetByStatus(ctx context.Context, status entity.StatusMail) ([]*entity.StatusHistory, error)
	GetAll(ctx context.Context) ([]*entity.StatusHistory, error)
	Update(ctx context.Context, history *entity.StatusHistory) error
	Delete(ctx context.Context, mailHistoryID string, status entity.StatusMail) error
	GetLatestByMailHistoryID(ctx context.Context, mailHistoryID string) (*entity.StatusHistory, error)
}

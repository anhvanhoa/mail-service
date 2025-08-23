package repository

import (
	"context"

	"mail-service/domain/entity"
)

type TypeMailRepository interface {
	Create(ctx context.Context, typeMail *entity.TypeMail) error
	GetByID(ctx context.Context, id string) (*entity.TypeMail, error)
	GetByName(ctx context.Context, name string) (*entity.TypeMail, error)
	GetByCreatedBy(ctx context.Context, createdBy string) ([]*entity.TypeMail, error)
	GetAll(ctx context.Context) ([]*entity.TypeMail, error)
	Update(ctx context.Context, typeMail *entity.TypeMail) error
	Delete(ctx context.Context, id string) error
}

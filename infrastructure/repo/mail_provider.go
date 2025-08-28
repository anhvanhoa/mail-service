package repo

import (
	"context"

	"mail-service/domain/entity"
	repository "mail-service/domain/repository"

	"github.com/go-pg/pg/v10"
)

type mailProviderRepository struct {
	db *pg.DB
}

func NewMailProviderRepository(db *pg.DB) repository.MailProviderRepository {
	return &mailProviderRepository{
		db: db,
	}
}

func (r *mailProviderRepository) Create(ctx context.Context, provider *entity.MailProvider) error {
	_, err := r.db.Model(provider).Insert()
	return err
}

func (r *mailProviderRepository) GetByEmail(ctx context.Context, email string) (*entity.MailProvider, error) {
	provider := &entity.MailProvider{}
	err := r.db.Model(provider).Where("email = ?", email).Select()
	if err != nil {
		return nil, err
	}
	return provider, nil
}

func (r *mailProviderRepository) GetAll(ctx context.Context) ([]*entity.MailProvider, error) {
	var providers []*entity.MailProvider
	err := r.db.Model(&providers).Select()
	return providers, err
}

func (r *mailProviderRepository) Update(ctx context.Context, provider *entity.MailProvider) error {
	_, err := r.db.Model(provider).Where("email = ?", provider.Email).Update()
	return err
}

func (r *mailProviderRepository) Delete(ctx context.Context, email string) error {
	_, err := r.db.Model(&entity.MailProvider{}).Where("email = ?", email).Delete()
	return err
}

func (r *mailProviderRepository) GetByTypeId(ctx context.Context, typeId string) ([]*entity.MailProvider, error) {
	var providers []*entity.MailProvider
	err := r.db.Model(&providers).Where("type_id = ?", typeId).Select()
	return providers, err
}

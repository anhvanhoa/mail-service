package repo

import (
	"context"

	"mail-service/domain/entity"
	repository "mail-service/domain/repository"

	"github.com/go-pg/pg/v10"
)

type mailHistoryRepository struct {
	db *pg.DB
}

func NewMailHistoryRepository(db *pg.DB) repository.MailHistoryRepository {
	return &mailHistoryRepository{
		db: db,
	}
}

func (r *mailHistoryRepository) Create(ctx context.Context, history *entity.MailHistory) error {
	db := getTx(ctx, r.db)
	_, err := db.Model(history).Insert()
	return err
}

func (r *mailHistoryRepository) GetByID(ctx context.Context, id string) (*entity.MailHistory, error) {
	db := getTx(ctx, r.db)
	history := &entity.MailHistory{}
	err := db.Model(history).Where("id = ?", id).Select()
	if err != nil {
		return nil, err
	}
	return history, nil
}

func (r *mailHistoryRepository) GetByTemplateID(ctx context.Context, templateID string) ([]*entity.MailHistory, error) {
	db := getTx(ctx, r.db)
	var histories []*entity.MailHistory
	err := db.Model(&histories).Where("template_id = ?", templateID).Select()
	return histories, err
}

func (r *mailHistoryRepository) GetByEmailProvider(ctx context.Context, emailProvider string) ([]*entity.MailHistory, error) {
	db := getTx(ctx, r.db)
	var histories []*entity.MailHistory
	err := db.Model(&histories).Where("email_provider = ?", emailProvider).Select()
	return histories, err
}

func (r *mailHistoryRepository) GetByCreatedBy(ctx context.Context, createdBy string) ([]*entity.MailHistory, error) {
	db := getTx(ctx, r.db)
	var histories []*entity.MailHistory
	err := db.Model(&histories).Where("created_by = ?", createdBy).Select()
	return histories, err
}

func (r *mailHistoryRepository) GetByTo(ctx context.Context, to string) ([]*entity.MailHistory, error) {
	db := getTx(ctx, r.db)
	var histories []*entity.MailHistory
	err := db.Model(&histories).Where("to = ?", to).Select()
	return histories, err
}

func (r *mailHistoryRepository) GetAll(ctx context.Context) ([]*entity.MailHistory, error) {
	db := getTx(ctx, r.db)
	var histories []*entity.MailHistory
	err := db.Model(&histories).Select()
	return histories, err
}

func (r *mailHistoryRepository) Update(ctx context.Context, history *entity.MailHistory) error {
	db := getTx(ctx, r.db)
	_, err := db.Model(history).Where("id = ?", history.ID).Update()
	return err
}

func (r *mailHistoryRepository) Delete(ctx context.Context, id string) error {
	db := getTx(ctx, r.db)
	_, err := db.Model(&entity.MailHistory{}).Where("id = ?", id).Delete()
	return err
}

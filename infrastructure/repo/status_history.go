package repo

import (
	"context"

	"mail-service/domain/entity"
	repository "mail-service/domain/repository"

	"github.com/go-pg/pg/v10"
)

type statusHistoryRepository struct {
	db *pg.DB
}

func NewStatusHistoryRepository(db *pg.DB) repository.StatusHistoryRepository {
	return &statusHistoryRepository{
		db: db,
	}
}

func (r *statusHistoryRepository) Create(ctx context.Context, history *entity.StatusHistory) error {
	db := getTx(ctx, r.db)
	_, err := db.Model(history).Insert()
	return err
}

func (r *statusHistoryRepository) GetByMailHistoryID(ctx context.Context, mailHistoryID string) ([]*entity.StatusHistory, error) {
	db := getTx(ctx, r.db)
	var histories []*entity.StatusHistory
	err := db.Model(&histories).Where("mail_history_id = ?", mailHistoryID).Order("created_at ASC").Select()
	return histories, err
}

func (r *statusHistoryRepository) GetByStatus(ctx context.Context, status entity.StatusMail) ([]*entity.StatusHistory, error) {
	db := getTx(ctx, r.db)
	var histories []*entity.StatusHistory
	err := db.Model(&histories).Where("status = ?", status).Select()
	return histories, err
}

func (r *statusHistoryRepository) GetAll(ctx context.Context) ([]*entity.StatusHistory, error) {
	db := getTx(ctx, r.db)
	var histories []*entity.StatusHistory
	err := db.Model(&histories).Select()
	return histories, err
}

func (r *statusHistoryRepository) Update(ctx context.Context, history *entity.StatusHistory) error {
	db := getTx(ctx, r.db)
	_, err := db.Model(history).Where("mail_history_id = ? AND status = ?", history.MailHistoryId, history.Status).Update()
	return err
}

func (r *statusHistoryRepository) Delete(ctx context.Context, mailHistoryID string, status entity.StatusMail) error {
	db := getTx(ctx, r.db)
	_, err := db.Model(&entity.StatusHistory{}).Where("mail_history_id = ? AND status = ?", mailHistoryID, status).Delete()
	return err
}

func (r *statusHistoryRepository) GetLatestByMailHistoryID(ctx context.Context, mailHistoryID string) (*entity.StatusHistory, error) {
	db := getTx(ctx, r.db)
	history := &entity.StatusHistory{}
	err := db.Model(history).Where("mail_history_id = ?", mailHistoryID).Order("created_at DESC").Limit(1).Select()
	if err != nil {
		return nil, err
	}
	return history, nil
}

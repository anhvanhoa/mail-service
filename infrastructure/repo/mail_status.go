package repo

import (
	"context"

	"mail-service/domain/entity"
	repository "mail-service/domain/repository"

	"github.com/go-pg/pg/v10"
)

type mailStatusRepository struct {
	db *pg.DB
}

func NewMailStatusRepository(db *pg.DB) repository.MailStatusRepository {
	return &mailStatusRepository{
		db: db,
	}
}

func (r *mailStatusRepository) Create(ctx context.Context, status *entity.MailStatus) error {
	db := getTx(ctx, r.db)
	_, err := db.Model(status).Insert()
	return err
}

func (r *mailStatusRepository) GetByStatus(ctx context.Context, status entity.StatusMail) (*entity.MailStatus, error) {
	db := getTx(ctx, r.db)
	mailStatus := &entity.MailStatus{}
	err := db.Model(mailStatus).Where("status = ?", status).Select()
	if err != nil {
		return nil, err
	}
	return mailStatus, nil
}

func (r *mailStatusRepository) GetByName(ctx context.Context, name string) (*entity.MailStatus, error) {
	db := getTx(ctx, r.db)
	mailStatus := &entity.MailStatus{}
	err := db.Model(mailStatus).Where("name = ?", name).Select()
	if err != nil {
		return nil, err
	}
	return mailStatus, nil
}

func (r *mailStatusRepository) GetAll(ctx context.Context) ([]*entity.MailStatus, error) {
	db := getTx(ctx, r.db)
	var statuses []*entity.MailStatus
	err := db.Model(&statuses).Select()
	return statuses, err
}

func (r *mailStatusRepository) Update(ctx context.Context, status *entity.MailStatus) error {
	db := getTx(ctx, r.db)
	_, err := db.Model(status).Where("status = ?", status.Status).Update()
	return err
}

func (r *mailStatusRepository) Delete(ctx context.Context, status entity.StatusMail) error {
	db := getTx(ctx, r.db)
	_, err := db.Model(&entity.MailStatus{}).Where("status = ?", status).Delete()
	return err
}

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
	_, err := r.db.Model(status).Insert()
	return err
}

func (r *mailStatusRepository) GetByStatus(ctx context.Context, status entity.StatusMail) (*entity.MailStatus, error) {
	mailStatus := &entity.MailStatus{}
	err := r.db.Model(mailStatus).Where("status = ?", status).Select()
	if err != nil {
		return nil, err
	}
	return mailStatus, nil
}

func (r *mailStatusRepository) GetByName(ctx context.Context, name string) (*entity.MailStatus, error) {
	mailStatus := &entity.MailStatus{}
	err := r.db.Model(mailStatus).Where("name = ?", name).Select()
	if err != nil {
		return nil, err
	}
	return mailStatus, nil
}

func (r *mailStatusRepository) GetAll(ctx context.Context) ([]*entity.MailStatus, error) {
	var statuses []*entity.MailStatus
	err := r.db.Model(&statuses).Select()
	return statuses, err
}

func (r *mailStatusRepository) Update(ctx context.Context, status *entity.MailStatus) error {
	_, err := r.db.Model(status).Where("status = ?", status.Status).Update()
	return err
}

func (r *mailStatusRepository) Delete(ctx context.Context, status entity.StatusMail) error {
	_, err := r.db.Model(&entity.MailStatus{}).Where("status = ?", status).Delete()
	return err
}

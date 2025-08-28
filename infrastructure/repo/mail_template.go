package repo

import (
	"context"

	"mail-service/domain/common"
	"mail-service/domain/entity"
	repository "mail-service/domain/repository"

	"github.com/go-pg/pg/v10"
)

type mailTemplateRepository struct {
	db *pg.DB
}

func NewMailTemplateRepository(db *pg.DB) repository.MailTemplateRepository {
	return &mailTemplateRepository{
		db: db,
	}
}

func (r *mailTemplateRepository) Create(ctx context.Context, template *entity.MailTemplate) error {
	_, err := r.db.Model(template).Insert()
	return err
}

func (r *mailTemplateRepository) GetByID(ctx context.Context, id string) (*entity.MailTemplate, error) {
	template := &entity.MailTemplate{}
	err := r.db.Model(template).Where("id = ?", id).Select()
	if err != nil {
		return nil, err
	}
	return template, nil
}

func (r *mailTemplateRepository) GetBySubject(ctx context.Context, subject string) (*entity.MailTemplate, error) {
	template := &entity.MailTemplate{}
	err := r.db.Model(template).Where("subject = ?", subject).Select()
	if err != nil {
		return nil, err
	}
	return template, nil
}

func (r *mailTemplateRepository) GetAll(ctx context.Context) ([]*entity.MailTemplate, error) {
	var templates []*entity.MailTemplate
	err := r.db.Model(&templates).Select()
	return templates, err
}

func (r *mailTemplateRepository) GetByStatus(ctx context.Context, status common.Status) ([]*entity.MailTemplate, error) {
	var templates []*entity.MailTemplate
	err := r.db.Model(&templates).Where("status = ?", status).Select()
	return templates, err
}

func (r *mailTemplateRepository) GetByProviderEmail(ctx context.Context, providerEmail string) ([]*entity.MailTemplate, error) {
	var templates []*entity.MailTemplate
	err := r.db.Model(&templates).Where("provider_email = ?", providerEmail).Select()
	return templates, err
}

func (r *mailTemplateRepository) Update(ctx context.Context, template *entity.MailTemplate) error {
	_, err := r.db.Model(template).Where("id = ?", template.ID).Update()
	return err
}

func (r *mailTemplateRepository) Delete(ctx context.Context, id string) error {
	_, err := r.db.Model(&entity.MailTemplate{}).Where("id = ?", id).Delete()
	return err
}

func (r *mailTemplateRepository) UpdateStatus(ctx context.Context, id string, status common.Status) error {
	_, err := r.db.Model(&entity.MailTemplate{}).Set("status = ?", status).Where("id = ?", id).Update()
	return err
}

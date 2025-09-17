package repo

import (
	"context"

	"mail-service/domain/entity"
	repository "mail-service/domain/repository"

	"github.com/go-pg/pg/v10"
)

type typeMailRepository struct {
	db *pg.DB
}

func NewTypeMailRepository(db *pg.DB) repository.TypeMailRepository {
	return &typeMailRepository{
		db: db,
	}
}

func (r *typeMailRepository) Create(ctx context.Context, typeMail *entity.TypeMail) error {
	_, err := r.db.Model(typeMail).Insert()
	return err
}

func (r *typeMailRepository) GetByID(ctx context.Context, id string) (*entity.TypeMail, error) {
	typeMail := &entity.TypeMail{}
	err := r.db.Model(typeMail).Where("id = ?", id).Select()
	if err != nil {
		return nil, err
	}
	return typeMail, nil
}

func (r *typeMailRepository) GetByName(ctx context.Context, name string) (*entity.TypeMail, error) {
	typeMail := &entity.TypeMail{}
	err := r.db.Model(typeMail).Where("name = ?", name).Select()
	if err != nil {
		return nil, err
	}
	return typeMail, nil
}

func (r *typeMailRepository) GetByCreatedBy(ctx context.Context, createdBy string) ([]*entity.TypeMail, error) {
	var typeMails []*entity.TypeMail
	err := r.db.Model(&typeMails).Where("created_by = ?", createdBy).Select()
	return typeMails, err
}

func (r *typeMailRepository) GetAll(ctx context.Context) ([]*entity.TypeMail, error) {
	var typeMails []*entity.TypeMail
	err := r.db.Model(&typeMails).Select()
	return typeMails, err
}

func (r *typeMailRepository) Update(ctx context.Context, typeMail *entity.TypeMail) error {
	_, err := r.db.Model(typeMail).Where("id = ?", typeMail.ID).UpdateNotZero()
	return err
}

func (r *typeMailRepository) Delete(ctx context.Context, id string) error {
	_, err := r.db.Model(&entity.TypeMail{}).Where("id = ?", id).Delete()
	return err
}

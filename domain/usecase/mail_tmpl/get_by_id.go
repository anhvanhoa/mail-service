package usecase

import (
	"context"
	"mail-service/domain/entity"
	"mail-service/domain/repository"
)

type GetByIdMailTmplUsecase interface {
	Execute(ctx context.Context, id string) (*entity.MailTemplate, error)
}

type getByIdMailTmplUsecase struct {
	mailTemplateRepository repository.MailTemplateRepository
}

func NewGetByIdMailTmplUsecase(mailTemplateRepository repository.MailTemplateRepository) GetByIdMailTmplUsecase {
	return &getByIdMailTmplUsecase{
		mailTemplateRepository: mailTemplateRepository,
	}
}

func (u *getByIdMailTmplUsecase) Execute(ctx context.Context, id string) (*entity.MailTemplate, error) {
	return u.mailTemplateRepository.GetByID(ctx, id)
}

package usecase

import (
	"context"
	"mail-service/domain/entity"
	"mail-service/domain/repository"
)

type GetAllMailTmplUsecase interface {
	Execute(ctx context.Context) ([]*entity.MailTemplate, error)
}

type getAllMailTmplUsecase struct {
	mailTemplateRepository repository.MailTemplateRepository
}

func NewGetAllMailTmplUsecase(mailTemplateRepository repository.MailTemplateRepository) GetAllMailTmplUsecase {
	return &getAllMailTmplUsecase{
		mailTemplateRepository: mailTemplateRepository,
	}
}

func (u *getAllMailTmplUsecase) Execute(ctx context.Context) ([]*entity.MailTemplate, error) {
	return u.mailTemplateRepository.GetAll(ctx)
}

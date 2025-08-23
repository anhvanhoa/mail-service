package usecase

import (
	"context"
	"mail-service/domain/entity"
	"mail-service/domain/repository"
)

type CreateMailTmplUsecase interface {
	Execute(ctx context.Context, req *entity.MailTemplate) error
}

type createMailTmplUsecase struct {
	mailTemplateRepository repository.MailTemplateRepository
}

func NewCreateMailTmplUsecase(mailTemplateRepository repository.MailTemplateRepository) CreateMailTmplUsecase {
	return &createMailTmplUsecase{
		mailTemplateRepository: mailTemplateRepository,
	}
}

func (u *createMailTmplUsecase) Execute(ctx context.Context, req *entity.MailTemplate) error {
	return u.mailTemplateRepository.Create(ctx, req)
}

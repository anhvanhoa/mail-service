package usecase

import (
	"context"
	"mail-service/domain/repository"
)

type DeleteByIdMailTmplUsecase interface {
	Execute(ctx context.Context, id string) error
}

type deleteByIdMailTmplUsecase struct {
	mailTemplateRepository repository.MailTemplateRepository
}

func NewDeleteByIdMailTmplUsecase(mailTemplateRepository repository.MailTemplateRepository) DeleteByIdMailTmplUsecase {
	return &deleteByIdMailTmplUsecase{
		mailTemplateRepository: mailTemplateRepository,
	}
}

func (u *deleteByIdMailTmplUsecase) Execute(ctx context.Context, id string) error {
	return u.mailTemplateRepository.Delete(ctx, id)
}

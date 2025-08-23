package usecase

import (
	"context"
	"mail-service/domain/entity"
	"mail-service/domain/repository"
)

type UpdateByIdMailTmplUsecase interface {
	Execute(ctx context.Context, req *entity.MailTemplate) error
}

type updateByIdMailTmplUsecase struct {
	mailTemplateRepository repository.MailTemplateRepository
}

func NewUpdateByIdMailTmplUsecase(mailTemplateRepository repository.MailTemplateRepository) UpdateByIdMailTmplUsecase {
	return &updateByIdMailTmplUsecase{
		mailTemplateRepository: mailTemplateRepository,
	}
}

func (u *updateByIdMailTmplUsecase) Execute(ctx context.Context, req *entity.MailTemplate) error {
	return u.mailTemplateRepository.Update(ctx, req)
}

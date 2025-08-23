package entity

import (
	"mail-service/domain/common"
	"time"
)

type MailProvider struct {
	tableName  struct{}      `pg:"mail_providers,alias:mp"`
	Email      string        `pg:"email,pk"`
	Password   string        `pg:"password"`
	UserName   string        `pg:"user_name"`
	Port       int           `pg:"port"`
	Host       string        `pg:"host"`
	Encryption string        `pg:"encryption"`
	Name       string        `pg:"name"`
	TypeId     string        `pg:"type_id"`
	CreatedBy  string        `pg:"created_by"`
	Status     common.Status `pg:"status"`
	CreatedAt  time.Time     `pg:"created_at"`
	UpdatedAt  *time.Time    `pg:"updated_at"`
}

func (mp *MailProvider) GetNameTable() any {
	return mp.tableName
}

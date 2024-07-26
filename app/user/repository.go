package user

import (
	"context"
	"database/sql"
	"go-rental/domain"
)

type Repository struct {
}

func (rpo Repository) Create(ctx context.Context, tx *sql.Tx, user domain.UserEntity) domain.UserEntity {
	//TODO implement me
	panic("implement me")
}

func (rpo Repository) FindByEmail(ctx context.Context, tx *sql.Tx, email string) (domain.UserEntity, error) {
	//TODO implement me
	panic("implement me")
}

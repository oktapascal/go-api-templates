package user

import (
	"context"
	"database/sql"
	"go-rental/domain"
)

type Repository struct {
}

func (r Repository) Create(ctx context.Context, tx *sql.Tx, user *domain.User) *domain.User {
	//TODO implement me
	panic("implement me")
}

func (r Repository) FindByEmail(ctx context.Context, tx *sql.Tx, email string) (*domain.User, error) {
	//TODO implement me
	panic("implement me")
}

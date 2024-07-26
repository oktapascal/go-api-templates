package user

import (
	"context"
	"database/sql"
	"go-rental/domain"
)

type Service struct {
	rpo domain.UserRepository
	db  *sql.DB
}

func (svc *Service) Save(ctx context.Context, request domain.User) domain.User {
	//TODO implement me
	panic("implement me")
}

func (svc *Service) GetByEmail(ctx context.Context, email string) domain.User {
	//TODO implement me
	panic("implement me")
}

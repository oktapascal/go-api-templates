package user

import (
	"context"
	"database/sql"
	"go-rental/domain"
	"go-rental/utils"
)

type Service struct {
	rpo domain.UserRepository
	db  *sql.DB
}

func (svc *Service) SaveUserWithoutSSO(ctx context.Context, request *domain.RegisterWithoutSSORequest) *domain.UserResponse {
	tx, err := svc.db.Begin()
	if err != nil {
		panic(err)
	}
	defer utils.CommitRollback(tx)

	user := &domain.User{
		IdNumber:    request.IdNumber,
		Email:       request.Email,
		Password:    &request.Password,
		PhoneNumber: request.PhoneNumber,
		Address:     request.Address,
		FirstName:   request.FirstName,
		LastName:    request.LastName,
	}
	user.Password = utils.Hash(user.Password)
	//TODO implement me
	panic("implement me")
}

func (svc *Service) SaveUserWithSSO(ctx context.Context, request *domain.RegisterWithSSORequest) *domain.UserResponse {
	//TODO implement me
	panic("implement me")
}

func (svc *Service) GetByEmail(ctx context.Context, email string) *domain.UserResponse {
	//TODO implement me
	panic("implement me")
}

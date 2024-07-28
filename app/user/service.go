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

	hash, errHash := utils.Hash(*user.Password)
	if errHash != nil {
		user.Password = &hash
	}

	return &domain.UserResponse{
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}
}

func (svc *Service) SaveUserWithSSO(ctx context.Context, request *domain.RegisterWithSSORequest) *domain.UserResponse {
	tx, err := svc.db.Begin()
	if err != nil {
		panic(err)
	}

	defer utils.CommitRollback(tx)

	user := &domain.User{
		IdNumber:    request.IdNumber,
		Email:       request.Email,
		PhoneNumber: request.PhoneNumber,
		Address:     request.Address,
		FirstName:   request.FirstName,
		LastName:    request.LastName,
		Provider:    &request.Provider,
		ProviderId:  &request.ProviderId,
	}

	return &domain.UserResponse{
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}
}

func (svc *Service) GetByEmail(ctx context.Context, email string) *domain.UserResponse {
	//TODO implement me
	panic("implement me")
}

func NewService(rpo domain.UserRepository, db *sql.DB) *Service {
	return &Service{
		rpo: rpo,
		db:  db,
	}
}

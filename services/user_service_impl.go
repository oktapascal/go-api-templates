package services

import (
	"context"
	"go-rental/requests"
	"go-rental/responses"
)

//type UserServiceImpl struct {
//	UserRepository repositories.UserRepository
//	DB             *sql.DB
//	Validate       *validator.Validate
//}

type UserServiceImpl struct{}

func NewUserService() UserService {
	return &UserServiceImpl{}
}

//func NewUserServiceImpl(userRepository repositories.UserRepository, DB *sql.DB, validate *validator.Validate) UserService {
//	return &UserServiceImpl{UserRepository: userRepository, DB: DB, Validate: validate}
//}

func (service *UserServiceImpl) Save(ctx context.Context, request requests.UserCreateRequest) responses.UserResponse {
	//TODO implement me
	panic("implement me")
}

func (service *UserServiceImpl) GetByEmail(ctx context.Context, email string) responses.UserResponse {
	//TODO implement me
	panic("implement me")
}

package services

import (
	"context"
	"github.com/go-playground/validator/v10"
	"go-rental/requests"
	"go-rental/responses"
	"log"
)

//type UserServiceImpl struct {
//	UserRepository repositories.UserRepository
//	DB             *sql.DB
//	Validate       *validator.Validate
//}

type UserServiceImpl struct {
	Validate *validator.Validate
}

func NewUserService(validate *validator.Validate) UserService {
	return &UserServiceImpl{
		Validate: validate,
	}
}

//func NewUserServiceImpl(userRepository repositories.UserRepository, DB *sql.DB, validate *validator.Validate) UserService {
//	return &UserServiceImpl{UserRepository: userRepository, DB: DB, Validate: validate}
//}

func (service *UserServiceImpl) Save(ctx context.Context, request requests.UserCreateRequest) responses.UserResponse {
	err := service.Validate.Struct(request)
	if err != nil {
		log.Panic(err)
	}
	//TODO implement me
	panic("implement me")
}

func (service *UserServiceImpl) GetByEmail(ctx context.Context, email string) responses.UserResponse {
	//TODO implement me
	panic("implement me")
}

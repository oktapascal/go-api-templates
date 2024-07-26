package services

import (
	"context"
	"go-rental/app/http/requests"
	"go-rental/responses"
)

type UserService interface {
	Save(ctx context.Context, request requests.UserCreateRequest) responses.UserResponse
	GetByEmail(ctx context.Context, email string) responses.UserResponse
}

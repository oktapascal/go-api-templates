package domain

import (
	"context"
	"database/sql"
	"net/http"
)

type (
	User struct {
		IdNumber    string `validate:"required,min=1,max=16" json:"id_number"`
		Email       string `validate:"required,email" json:"email"`
		PhoneNumber string `validate:"required,min=10,max=13" json:"phone_number"`
		Address     string `validate:"required" json:"address"`
		FirstName   string `validate:"required,min=1,max=50" json:"first_name"`
		LastName    string `validate:"required,min=1,max=50" json:"last_name"`
		Provider    string `validate:"required" json:"provider"`
		ProviderId  int8   `validate:"required" json:"provider_id"`
		PhotoIdCard string `json:"photo_id_card"`
	}

	UserRepository interface {
		Create(ctx context.Context, tx *sql.Tx, user *User) *User
		FindByEmail(ctx context.Context, tx *sql.Tx, email string) (*User, error)
	}

	UserService interface {
		Save(ctx context.Context, request *User) *User
		GetByEmail(ctx context.Context, email string) *User
	}

	UserHandler interface {
		Store() http.HandlerFunc
		GetByEmail() http.HandlerFunc
	}
)

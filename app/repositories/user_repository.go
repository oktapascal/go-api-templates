package repositories

import (
	"context"
	"database/sql"
	"go-rental/app/models"
)

type UserRepository interface {
	Create(ctx context.Context, tx *sql.Tx, user models.User) models.User
	FindByEmail(ctx context.Context, tx *sql.Tx, email string) (models.User, error)
}

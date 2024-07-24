package repositories

import (
	"context"
	"database/sql"
	"go-rental/entities"
)

type UserRepository interface {
	Create(ctx context.Context, tx *sql.Tx, user entities.User) entities.User
	FindByEmail(ctx context.Context, tx *sql.Tx, email string) (entities.User, error)
}

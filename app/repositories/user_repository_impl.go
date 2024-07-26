package repositories

import (
	"context"
	"database/sql"
	"errors"
	"go-rental/app/models"
	"log"
)

type UserRepositoryImpl struct {
}

func NewUserRepositoryImpl() UserRepository {
	return &UserRepositoryImpl{}
}

func (repo *UserRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, user models.User) models.User {
	query := "insert into users(id_number, email, phone_number, address, first_name, last_name, provider, provider_id) values (?,?,?,?,?,?,?,?)"

	_, err := tx.ExecContext(ctx, query, user.Email, user.PhoneNumber, user.Address, user.FirstName, user.LastName,
		user.Provider, user.ProviderId)
	if err != nil {
		log.Panic(err)
	}

	return user
}

func (repo *UserRepositoryImpl) FindByEmail(ctx context.Context, tx *sql.Tx, email string) (models.User, error) {
	query := "select * from users where email = ?"

	rows, err := tx.QueryContext(ctx, query, email)
	if err != nil {
		log.Panic(err)
	}

	defer func() {
		err := rows.Close()
		if err != nil {
			log.Panic(err)
		}
	}()

	user := models.User{}

	if rows.Next() {
		err := rows.Scan(&user.IdNumber, &user.Email, &user.PhoneNumber, &user.Address, &user.FirstName,
			&user.LastName, &user.Provider, &user.ProviderId, &user.PhotoIdCard)
		if err != nil {
			log.Panic(err)
		}

		return user, nil
	} else {
		return user, errors.New("not found")
	}
}

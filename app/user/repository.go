package user

import (
	"context"
	"database/sql"
	"go-rental/domain"
)

type Repository struct {
}

func (rpo *Repository) Create(ctx context.Context, tx *sql.Tx, user *domain.User) *domain.User {
	var query = "insert into users (id_number,email,phone_number,address,first_name,last_name,provider,provider_id) values (?,?,?,?,?,?,?,?)"
	var _, err = tx.ExecContext(ctx, query, user.IdNumber, user.Email, user.PhoneNumber, user.Address, user.FirstName, user.LastName, user.Provider, user.ProviderId)
	if err != nil {
		panic(err)
	}

	return user
}

func (rpo *Repository) FindByEmail(ctx context.Context, tx *sql.Tx, email string) (*domain.User, error) {
	//var query = "select id_number, email, phone_number, address, first_name, last_name, role, provider, provider_id, photo_id_card from users where email = ?"
	//TODO implement me
	panic("implement me")
}

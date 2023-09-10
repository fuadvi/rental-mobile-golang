package users

import (
	"Rental_Mobil/helpers"
	"Rental_Mobil/model/domain"
	"context"
	"database/sql"
)

type UserRepositoryImpl struct {
}

func NewUserRepositoryImpl() *UserRepositoryImpl {
	return &UserRepositoryImpl{}
}

func (repository UserRepositoryImpl) GetAll(ctx context.Context, tx *sql.Tx) []domain.User {
	//TODO implement me
	panic("implement me")
}

func (repository UserRepositoryImpl) Get(ctx context.Context, tx *sql.Tx, leaseTypeId int) (domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (repository UserRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, user domain.User) error {
	querySql := "INSERT INTO users (name,email, password) VALUES (?,?,?)"
	_, err := tx.ExecContext(ctx, querySql, user.Name, user.Email, user.Password)
	helpers.PanicIfError(err)
	return nil
}

func (repository UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, leaseTypeId int, user domain.User) error {
	querySQL := "UPDATE users SET name =?, email=? WHERE id =?"
	_, err := tx.ExecContext(ctx, querySQL, user.Name, user.Email, leaseTypeId)
	helpers.PanicIfError(err)
	return nil
}

func (repository UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, leaseTypeId int) error {
	querySQL := "DELETE FROM users WHERE id =?"
	_, err := tx.ExecContext(ctx, querySQL, leaseTypeId)
	helpers.PanicIfError(err)

	return nil
}

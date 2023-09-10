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
	querySQL := "SELECT id, name,email, password FROM users"
	row, err := tx.QueryContext(ctx, querySQL)
	helpers.PanicIfError(err)

	var users []domain.User
	for row.Next() {
		var user domain.User
		err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Password)
		helpers.PanicIfError(err)
		users = append(users, user)
	}

	return users
}

func (repository UserRepositoryImpl) Get(ctx context.Context, tx *sql.Tx, userId int) (domain.User, error) {
	querySQL := "SELECT id, name,email, password FROM users WHERE id = ?"
	row, err := tx.QueryContext(ctx, querySQL, userId)
	helpers.PanicIfError(err)

	var user domain.User
	if row.Next() {
		err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Password)
		helpers.PanicIfError(err)
	}

	return user, nil
}

func (repository UserRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, user domain.User) error {
	querySql := "INSERT INTO users (name,email, password) VALUES (?,?,?)"
	_, err := tx.ExecContext(ctx, querySql, user.Name, user.Email, user.Password)
	helpers.PanicIfError(err)
	return nil
}

func (repository UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, userId int, user domain.User) error {
	querySQL := "UPDATE users SET name =?, email=? WHERE id =?"
	_, err := tx.ExecContext(ctx, querySQL, user.Name, user.Email, userId)
	helpers.PanicIfError(err)
	return nil
}

func (repository UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, leaseTypeId int) error {
	querySQL := "DELETE FROM users WHERE id =?"
	_, err := tx.ExecContext(ctx, querySQL, leaseTypeId)
	helpers.PanicIfError(err)

	return nil
}

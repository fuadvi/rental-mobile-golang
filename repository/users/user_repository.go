package users

import (
	"Rental_Mobil/model/domain"
	"context"
	"database/sql"
)

type UserRepository interface {
	GetAll(ctx context.Context, tx *sql.Tx) []domain.User
	Get(ctx context.Context, tx *sql.Tx, userId int) (domain.User, error)
	Save(ctx context.Context, tx *sql.Tx, user domain.User) error
	Update(ctx context.Context, tx *sql.Tx, userId int, user domain.User) error
	Delete(ctx context.Context, tx *sql.Tx, userId int) error
}

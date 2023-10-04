package cars

import (
	"Rental_Mobil/model/domain"
	"Rental_Mobil/model/dto"
	"context"
	"database/sql"
)

type CarRepository interface {
	GetAll(ctx context.Context, tx *sql.Tx) []domain.Car
	Get(ctx context.Context, tx *sql.Tx, carId int) (domain.Car, error)
	Create(ctx context.Context, tx *sql.Tx, request dto.CarRequestDto) (int, error)
	Update(ctx context.Context, tx *sql.Tx, request dto.CarRequestDto, carId int)
	Delete(ctx context.Context, tx *sql.Tx, carId int) error
}

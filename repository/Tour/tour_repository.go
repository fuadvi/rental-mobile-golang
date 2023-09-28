package Tour

import (
	"Rental_Mobil/model/domain"
	"Rental_Mobil/model/dto"
	"context"
	"database/sql"
)

type TourRepository interface {
	GetAll(ctx context.Context, tx *sql.Tx) []domain.Tour
	Get(ctx context.Context, tx *sql.Tx, tourId int) (domain.Tour, error)
	Create(ctx context.Context, tx *sql.Tx, request dto.TourRequestDto) error
	Update(ctx context.Context, tx *sql.Tx, request dto.TourRequestDto, tourId int) error
	Delete(ctx context.Context, tx *sql.Tx, tourId int) error
}

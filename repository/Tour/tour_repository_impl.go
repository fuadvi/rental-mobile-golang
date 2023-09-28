package Tour

import (
	"Rental_Mobil/helpers"
	"Rental_Mobil/model/domain"
	"Rental_Mobil/model/dto"
	"context"
	"database/sql"
	"errors"
)

type TourRepositoryImpl struct {
}

func NewTourRepositoryImpl() *TourRepositoryImpl {
	return &TourRepositoryImpl{}
}

func (repository TourRepositoryImpl) GetAll(ctx context.Context, tx *sql.Tx) []domain.Tour {
	querySQL := "SELECT id,title,price,duration,description FROM tours"

	rows, err := tx.QueryContext(ctx, querySQL)
	helpers.PanicIfError(err)

	var tours []domain.Tour
	for rows.Next() {
		var tour domain.Tour
		err = rows.Scan(&tour.Id, &tour.Title, &tour.Price, &tour.Duration, &tour.Description)
		helpers.PanicIfError(err)
		tours = append(tours, tour)
	}

	return tours
}

func (repository TourRepositoryImpl) Get(ctx context.Context, tx *sql.Tx, tourId int) (domain.Tour, error) {
	querySQL := "SELECT id,title,price,duration,description FROM tours where id =?"

	row, err := tx.QueryContext(ctx, querySQL, tourId)
	helpers.PanicIfError(err)
	var tour domain.Tour

	if row.Next() {
		err = row.Scan(&tour.Id, &tour.Title, &tour.Price, &tour.Duration, &tour.Description)
		helpers.PanicIfError(err)
	} else {
		return tour, errors.New("Not Found")
	}

	return tour, nil
}

func (repository TourRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, request dto.TourRequestDto) error {
	querySQL := "INSERT INTO tours (title,price,duration,description) VALUES (?,?,?,?)"

	_, err := tx.ExecContext(ctx, querySQL, request.Title, request.Price, request.Duration, request.Description)
	helpers.PanicIfError(err)

	return nil
}

func (repository TourRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, request dto.TourRequestDto, tourId int) error {
	querySQL := "UPDATE tours SET title=?,price=?,duration=?,description=? where id = ?"

	_, err := tx.ExecContext(ctx, querySQL, request.Title, request.Price, request.Duration, request.Description, tourId)
	helpers.PanicIfError(err)

	return nil
}

func (repository TourRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, tourId int) error {
	querySQL := "DELETE FROM tours where id = ?"

	_, err := tx.ExecContext(ctx, querySQL, tourId)
	helpers.PanicIfError(err)

	return nil
}

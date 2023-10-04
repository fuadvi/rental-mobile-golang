package cars

import (
	"Rental_Mobil/helpers"
	"Rental_Mobil/model/domain"
	"Rental_Mobil/model/dto"
	"context"
	"database/sql"
	"errors"
)

type CarRepositoryImpl struct {
}

func NewCarRepositoryImpl() *CarRepositoryImpl {
	return &CarRepositoryImpl{}
}

func (repository CarRepositoryImpl) GetAll(ctx context.Context, tx *sql.Tx) []domain.Car {
	querySql := "SELECT id, title, price, duration, image_url, description, passenger, luggage, car_type, is_driver FROM cars"
	rows, err := tx.QueryContext(ctx, querySql)
	helpers.PanicIfError(err)

	var cars []domain.Car
	for rows.Next() {
		var car domain.Car

		err := rows.Scan(&car.ID, &car.TITLE, &car.PRICE, &car.DURATION, &car.IMGURL, &car.PASSENGER, &car.LUGGAGE, &car.CARTYPE, &car.ISDRIVER)
		helpers.PanicIfError(err)

		cars = append(cars, car)
	}

	return cars
}

func (repository CarRepositoryImpl) Get(ctx context.Context, tx *sql.Tx, carId int) (domain.Car, error) {
	querySql := "SELECT id, title, price, duration, image_url, description, passenger, luggage, car_type, is_driver FROM cars WHERE id = ?"
	row, err := tx.QueryContext(ctx, querySql, carId)
	helpers.PanicIfError(err)

	var car domain.Car
	if row.Next() {
		err := row.Scan(&car.ID, &car.TITLE, &car.PRICE, &car.DURATION, &car.IMGURL, &car.PASSENGER, &car.LUGGAGE, &car.CARTYPE, &car.ISDRIVER)
		helpers.PanicIfError(err)
	} else {
		return car, errors.New("Not Found")
	}

	return car, nil
}

func (repository CarRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, request dto.CarRequestDto) (int, error) {
	querySql := "INSERT INTO cars (title, price, duration, image_url, description, passenger, luggage, car_type, is_driver) VALUES (?,?,?,?,?,?,?,?,?)"

	result, err := tx.ExecContext(ctx, querySql, request.TITLE, request.PRICE, request.DURATION, request.IMGURL, request.DESCRIPTION, request.PASSENGER, request.LUGGAGE, request.CARTYPE, request.ISDRIVER)
	helpers.PanicIfError(err)
	carId, err := result.LastInsertId()
	helpers.PanicIfError(err)

	return int(carId), nil
}

func (repository CarRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, request dto.CarRequestDto, carId int) {
	querySql := "update cars set title = ?, price = ?, duration = ?, image_url = ?, description = ?, passenger = ?, luggage = ?, car_type = ?, is_drive = ? where id = ?"
	_, err := tx.ExecContext(ctx, querySql, request.TITLE, request.PRICE, request.DURATION, request.IMGURL, request.DESCRIPTION, request.PASSENGER, request.LUGGAGE, request.CARTYPE, request.ISDRIVER)
	helpers.PanicIfError(err)
}

func (repository CarRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, carId int) error {
	querySql := "DELETE FROM cars WHERE id = ?"
	_, err := tx.ExecContext(ctx, querySql, carId)
	helpers.PanicIfError(err)

	return nil
}

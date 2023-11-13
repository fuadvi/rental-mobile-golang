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

func (repository CarRepositoryImpl) GetAll(ctx context.Context, tx *sql.Tx) []dto.CarResponseDto {
	querySql := "SELECT c.id, c.title, c.price, c.duration, c.image_url, c.description, c.passenger, c.luggage, c.car_type, c.is_driver, lt.title FROM cars c " +
		"join car_lease_type clt on c.id = clt.car_id " +
		"join lease_types lt on clt.lease_type_id = lt.id"

	rows, err := tx.QueryContext(ctx, querySql)
	helpers.PanicIfError(err)

	var cars []dto.CarResponseDto
	for rows.Next() {
		var car dto.CarResponseDto

		err := rows.Scan(&car.ID, &car.TITLE, &car.PRICE, &car.DURATION, &car.IMGURL, &car.PASSENGER, &car.LUGGAGE, &car.CARTYPE, &car.ISDRIVER, &car.LEASETYPE)
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
		err := row.Scan(&car.ID, &car.TITLE, &car.PRICE, &car.DURATION, &car.IMGURL, &car.DESCRIPTION, &car.PASSENGER, &car.LUGGAGE, &car.CARTYPE, &car.ISDRIVER)
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

func (repository CarRepositoryImpl) CarLeaseTypeCreate(ctx context.Context, tx *sql.Tx, leaseTypeId int, carId int) error {
	querySql := "INSERT INTO car_lease_type (lease_type_id, car_id) VALUES (?,?)"
	_, err := tx.ExecContext(ctx, querySql, leaseTypeId, carId)
	helpers.PanicIfError(err)
	return nil
}

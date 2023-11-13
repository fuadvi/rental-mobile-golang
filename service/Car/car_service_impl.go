package Car

import (
	"Rental_Mobil/exception"
	"Rental_Mobil/helpers"
	"Rental_Mobil/model/dto"
	"Rental_Mobil/repository/cars"
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
)

type CarServiceImpl struct {
	carRepo  cars.CarRepository
	db       *sql.DB
	validate *validator.Validate
}

func NewCarServiceImpl(carRepo cars.CarRepository, db *sql.DB, validate *validator.Validate) *CarServiceImpl {
	return &CarServiceImpl{carRepo: carRepo, db: db, validate: validate}
}

func (service CarServiceImpl) GetAll(ctx context.Context) []dto.CarResponseDto {
	tx, err := service.db.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	return service.carRepo.GetAll(ctx, tx)
}

func (service CarServiceImpl) Get(ctx context.Context, carId int) dto.CarResponseDto {
	tx, err := service.db.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	car, err := service.carRepo.Get(ctx, tx, carId)

	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return car.ToCarResponse()
}

func (service CarServiceImpl) Create(ctx context.Context, request dto.CarRequestDto) {
	err := service.validate.Struct(request)
	helpers.PanicIfError(err)

	tx, err := service.db.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	carId, err := service.carRepo.Create(ctx, tx, request)
	helpers.PanicIfError(err)

	// insert data ke table car lease type
	err = service.carRepo.CarLeaseTypeCreate(ctx, tx, request.LEASETYPEID, carId)

}

func (service CarServiceImpl) Update(ctx context.Context, request dto.CarRequestDto, carId int) {
	err := service.validate.Struct(request)
	helpers.PanicIfError(err)

	tx, err := service.db.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	service.carRepo.Update(ctx, tx, request, carId)
}

func (service CarServiceImpl) Delete(ctx context.Context, carId int) {
	tx, err := service.db.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	err = service.carRepo.Delete(ctx, tx, carId)
	helpers.PanicIfError(err)
}

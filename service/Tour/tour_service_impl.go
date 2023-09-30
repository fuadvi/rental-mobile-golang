package Tour

import (
	"Rental_Mobil/exception"
	"Rental_Mobil/helpers"
	"Rental_Mobil/model/dto"
	"Rental_Mobil/repository/Tour"
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
)

type TourServiceImp struct {
	TourRepo Tour.TourRepository
	DB       *sql.DB
	validate *validator.Validate
}

func (service TourServiceImp) GetAll(ctx context.Context) []dto.TourResponseDto {
	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)
	helpers.CommitOrRollback(tx)

	tours := service.TourRepo.GetAll(context.Background(), tx)

	var response []dto.TourResponseDto

	for _, tour := range tours {
		response = append(response, tour.ToResponse())
	}

	return response

}

func (service TourServiceImp) Get(ctx context.Context, tourId int) dto.TourResponseDto {
	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)
	helpers.CommitOrRollback(tx)

	tour, err := service.TourRepo.Get(context.Background(), tx, tourId)

	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return tour.ToResponse()
}

func (service TourServiceImp) Create(ctx context.Context, request dto.TourRequestDto) error {
	err := service.validate.Struct(request)
	helpers.PanicIfError(err)

	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)
	helpers.CommitOrRollback(tx)

	err = service.TourRepo.Create(context.Background(), tx, request)
	helpers.PanicIfError(err)

	return nil
}

func (service TourServiceImp) Update(ctx context.Context, request dto.TourRequestDto, tourId int) error {
	err := service.validate.Struct(request)
	helpers.PanicIfError(err)

	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)
	helpers.CommitOrRollback(tx)

	err = service.TourRepo.Update(context.Background(), tx, request, tourId)
	helpers.PanicIfError(err)

	return nil
}

func (service TourServiceImp) Delete(ctx context.Context, tourId int) error {
	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)
	helpers.CommitOrRollback(tx)

	err = service.TourRepo.Delete(context.Background(), tx, tourId)
	helpers.PanicIfError(err)

	return nil
}

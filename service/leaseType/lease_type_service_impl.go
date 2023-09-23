package leaseType

import (
	"Rental_Mobil/helpers"
	"Rental_Mobil/model/dto"
	"Rental_Mobil/repository/lease_types"
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
)

type leaseTypeServiceImpl struct {
	LeasTypeRepo lease_types.LeaseTypeRepository
	DB           *sql.DB
	validate     *validator.Validate
}

func (service leaseTypeServiceImpl) GetList(ctx context.Context) []dto.LeaseTypeResponseDto {
	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	leaseTypes := service.LeasTypeRepo.GetAll(ctx, tx)

	var response []dto.LeaseTypeResponseDto
	for _, leaseType := range leaseTypes {
		response = append(response, leaseType.ToResponse())
	}

	return response
}

func (service leaseTypeServiceImpl) GET(ctx context.Context, leaseTypeId int) dto.LeaseTypeResponseDto {
	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	return service.LeasTypeRepo.Get(ctx, tx, leaseTypeId).ToResponse()
}

func (service leaseTypeServiceImpl) Create(ctx context.Context, request dto.LeaseTypeRequest) error {
	err := service.validate.Struct(request)
	helpers.PanicIfError(err)

	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	err = service.LeasTypeRepo.Create(ctx, tx, request)
	helpers.PanicIfError(err)

	return nil
}

func (service leaseTypeServiceImpl) Update(ctx context.Context, request dto.LeaseTypeRequest, leaseTypeId int) error {
	err := service.validate.Struct(request)
	helpers.PanicIfError(err)

	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	err = service.LeasTypeRepo.Update(ctx, tx, request, leaseTypeId)
	helpers.PanicIfError(err)
	return nil
}

func (service leaseTypeServiceImpl) Delete(ctx context.Context, leaseTypeId int) error {
	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	service.LeasTypeRepo.Delete(ctx, tx, leaseTypeId)
	return nil
}

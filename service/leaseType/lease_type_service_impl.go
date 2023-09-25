package leaseType

import (
	"Rental_Mobil/exception"
	"Rental_Mobil/helpers"
	"Rental_Mobil/model/dto"
	"Rental_Mobil/repository/lease_types"
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"log"
)

type LeaseTypeServiceImpl struct {
	LeasTypeRepo lease_types.LeaseTypeRepository
	DB           *sql.DB
	validate     *validator.Validate
}

func NewLeaseTypeServiceImpl(leasTypeRepo lease_types.LeaseTypeRepository, DB *sql.DB, validate *validator.Validate) *LeaseTypeServiceImpl {
	return &LeaseTypeServiceImpl{LeasTypeRepo: leasTypeRepo, DB: DB, validate: validate}
}

func (service LeaseTypeServiceImpl) GetList(ctx context.Context) []dto.LeaseTypeResponseDto {
	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	leaseTypes := service.LeasTypeRepo.GetAll(ctx, tx)
	log.Println(leaseTypes)

	var response []dto.LeaseTypeResponseDto
	for _, leaseType := range leaseTypes {
		response = append(response, leaseType.ToResponse())
	}

	return response
}

func (service LeaseTypeServiceImpl) GET(ctx context.Context, leaseTypeId int) dto.LeaseTypeResponseDto {
	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	leaseType, err := service.LeasTypeRepo.Get(ctx, tx, leaseTypeId)

	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return leaseType.ToResponse()
}

func (service LeaseTypeServiceImpl) Create(ctx context.Context, request dto.LeaseTypeRequest) error {
	err := service.validate.Struct(request)
	helpers.PanicIfError(err)

	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	err = service.LeasTypeRepo.Create(ctx, tx, request)
	helpers.PanicIfError(err)

	return nil
}

func (service LeaseTypeServiceImpl) Update(ctx context.Context, request dto.LeaseTypeRequest, leaseTypeId int) error {
	err := service.validate.Struct(request)
	helpers.PanicIfError(err)

	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	err = service.LeasTypeRepo.Update(ctx, tx, request, leaseTypeId)
	helpers.PanicIfError(err)
	return nil
}

func (service LeaseTypeServiceImpl) Delete(ctx context.Context, leaseTypeId int) error {
	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	service.LeasTypeRepo.Delete(ctx, tx, leaseTypeId)
	return nil
}

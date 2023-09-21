package lease_types

import (
	"Rental_Mobil/model/domain"
	"Rental_Mobil/model/dto"
	"context"
	"database/sql"
)

type LeaseTypeRepositoryImpl struct {
}

func (repository LeaseTypeRepositoryImpl) Get(ctx context.Context, tx *sql.Tx, leaseTypeId int) domain.LeaseType {
	//TODO implement me
	panic("implement me")
}

func (repository LeaseTypeRepositoryImpl) GetAll(ctx context.Context, tx *sql.Tx) []domain.LeaseType {
	//TODO implement me
	panic("implement me")
}

func (repository LeaseTypeRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, request dto.LeaseTypeRequest) error {
	//TODO implement me
	panic("implement me")
}

func (repository LeaseTypeRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, request dto.LeaseTypeRequest, leaseTypeId int) error {
	//TODO implement me
	panic("implement me")
}

func (repository LeaseTypeRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, leaseTypeId int) error {
	//TODO implement me
	panic("implement me")
}

package lease_types

import (
	"Rental_Mobil/model/domain"
	"Rental_Mobil/model/dto"
	"context"
	"database/sql"
)

type LeaseTypeRepository interface {
	Get(ctx context.Context, tx *sql.Tx, leaseTypeId int) (domain.LeaseType, error)
	GetAll(ctx context.Context, tx *sql.Tx) []domain.LeaseType
	Create(ctx context.Context, tx *sql.Tx, request dto.LeaseTypeRequest) error
	Update(ctx context.Context, tx *sql.Tx, request dto.LeaseTypeRequest, leaseTypeId int) error
	Delete(ctx context.Context, tx *sql.Tx, leaseTypeId int)
}

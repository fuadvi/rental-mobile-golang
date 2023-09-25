package lease_types

import (
	"Rental_Mobil/helpers"
	"Rental_Mobil/model/domain"
	"Rental_Mobil/model/dto"
	"context"
	"database/sql"
	"errors"
)

type LeaseTypeRepositoryImpl struct {
}

func NewLeaseTypeRepositoryImpl() *LeaseTypeRepositoryImpl {
	return &LeaseTypeRepositoryImpl{}
}

func (repository LeaseTypeRepositoryImpl) Get(ctx context.Context, tx *sql.Tx, leaseTypeId int) (domain.LeaseType, error) {
	querySQL := "SELECT id, title, description FROM lease_types where id = ?"

	rows, err := tx.QueryContext(ctx, querySQL, leaseTypeId)
	helpers.PanicIfError(err)

	var leaseType domain.LeaseType

	if rows.Next() {
		err := rows.Scan(&leaseType.Id, &leaseType.Title, &leaseType.Description)
		helpers.PanicIfError(err)
	} else {
		return leaseType, errors.New("Not Found")
	}

	return leaseType, nil
}

func (repository LeaseTypeRepositoryImpl) GetAll(ctx context.Context, tx *sql.Tx) []domain.LeaseType {
	querySQL := "SELECT id, title, description FROM lease_types"

	rows, err := tx.QueryContext(ctx, querySQL)
	helpers.PanicIfError(err)

	var leaseTypes []domain.LeaseType

	for rows.Next() {
		var leaseType domain.LeaseType
		err := rows.Scan(&leaseType.Id, &leaseType.Title, &leaseType.Description)
		helpers.PanicIfError(err)

		leaseTypes = append(leaseTypes, leaseType)
	}

	return leaseTypes
}

func (repository LeaseTypeRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, request dto.LeaseTypeRequest) error {
	querySQL := "INSERT INTO lease_types (title,description) values (?,?)"
	_, err := tx.ExecContext(ctx, querySQL, request.Title, request.Description)
	helpers.PanicIfError(err)

	return nil
}

func (repository LeaseTypeRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, request dto.LeaseTypeRequest, leaseTypeId int) error {
	querySQL := "UPDATE lease_types set title =?, description =? WHERE id =?"
	_, err := tx.QueryContext(ctx, querySQL, request.Title, request.Description, leaseTypeId)
	helpers.PanicIfError(err)

	return nil
}

func (repository LeaseTypeRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, leaseTypeId int) {
	querySQL := "DELETE FROM lease_types WHERE id =?"
	_, err := tx.ExecContext(ctx, querySQL, leaseTypeId)
	helpers.PanicIfError(err)
}

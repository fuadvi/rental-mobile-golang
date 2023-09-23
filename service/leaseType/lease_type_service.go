package leaseType

import (
	"Rental_Mobil/model/dto"
	"context"
)

type leaseTypeService interface {
	GetList(ctx context.Context) []dto.LeaseTypeResponseDto
	GET(ctx context.Context, leaseTypeId int) dto.LeaseTypeResponseDto
	Create(ctx context.Context, request dto.LeaseTypeRequest) error
	Update(ctx context.Context, request dto.LeaseTypeRequest, leaseTypeId int) error
	Delete(ctx context.Context, leaseTypeId int) error
}

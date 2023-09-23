package domain

import "Rental_Mobil/model/dto"

type LeaseType struct {
	Id          int
	Title       string
	Description string
}

func (leaseType LeaseType) ToResponse() dto.LeaseTypeResponseDto {
	return dto.LeaseTypeResponseDto{
		Id:          leaseType.Id,
		Title:       leaseType.Title,
		Description: leaseType.Description,
	}
}

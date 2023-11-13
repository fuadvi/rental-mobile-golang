package Car

import (
	"Rental_Mobil/model/dto"
	"context"
)

type CarService interface {
	GetAll(ctx context.Context) []dto.CarResponseDto
	Get(ctx context.Context, carId int) dto.CarResponseDto
	Create(ctx context.Context, request dto.CarRequestDto)
	Update(ctx context.Context, request dto.CarRequestDto, carId int)
	Delete(ctx context.Context, carId int)
}

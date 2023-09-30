package Tour

import (
	"Rental_Mobil/model/dto"
	"context"
)

type TourService interface {
	GetAll(ctx context.Context) []dto.TourResponseDto
	Get(ctx context.Context, tourId int) dto.TourResponseDto
	Create(ctx context.Context, request dto.TourRequestDto) error
	Update(ctx context.Context, request dto.TourRequestDto, tourId int) error
	Delete(ctx context.Context, tourId int) error
}

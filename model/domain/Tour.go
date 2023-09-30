package domain

import "Rental_Mobil/model/dto"

type Tour struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Price       int    `json:"price"`
	Duration    string `json:"duration"`
	Description string `json:"description"`
}

func (tour Tour) ToResponse() dto.TourResponseDto {
	return dto.TourResponseDto{
		Id:          tour.Id,
		Title:       tour.Title,
		Price:       tour.Price,
		Duration:    tour.Duration,
		Description: tour.Description,
	}
}

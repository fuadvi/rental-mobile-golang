package domain

import "Rental_Mobil/model/dto"

type Car struct {
	ID          int
	TITLE       string
	PRICE       int
	DURATION    string
	IMGURL      string
	DESCRIPTION string
	PASSENGER   int8
	LUGGAGE     int8
	CARTYPE     string
	ISDRIVER    bool
}

func (car Car) ToCarResponse() dto.CarResponseDto {
	return dto.CarResponseDto{
		ID:          car.ID,
		TITLE:       car.TITLE,
		PRICE:       car.PRICE,
		DURATION:    car.DURATION,
		IMGURL:      car.IMGURL,
		DESCRIPTION: car.DESCRIPTION,
		PASSENGER:   car.PASSENGER,
		LUGGAGE:     car.LUGGAGE,
		CARTYPE:     car.CARTYPE,
		ISDRIVER:    car.ISDRIVER,
	}
}

package dto

type TourRequestDto struct {
	Title       string `json:"title" validate:"required"`
	Price       int    `json:"price" validate:"required"`
	Duration    string `json:"duration" validate:"required"`
	Description string `json:"description" validate:"required,number"`
}

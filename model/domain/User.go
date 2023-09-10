package domain

import "Rental_Mobil/model/dto"

type User struct {
	Id       int
	Name     string
	Email    string
	Password string
}

func (user User) ToDtoResponse() dto.UserResponse {
	return dto.UserResponse{
		ID:    user.Id,
		Name:  user.Name,
		Email: user.Email,
	}
}

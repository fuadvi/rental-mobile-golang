package user

import (
	"Rental_Mobil/model/dto"
	"Rental_Mobil/model/web"
	"context"
)

type UserService interface {
	GetAll(ctx context.Context) []dto.UserResponse
	Get(ctx context.Context, userId int) dto.UserResponse
	GetByEmail(ctx context.Context, request web.LoginRequest) dto.UserResponse
	Create(ctx context.Context, request web.UserRequest) error
	Update(ctx context.Context, request web.UserRequest, userId int) error
	Delete(ctx context.Context, userId int) error
}

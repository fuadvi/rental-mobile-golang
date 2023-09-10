package user

import (
	"Rental_Mobil/helpers"
	"Rental_Mobil/model/domain"
	"Rental_Mobil/model/dto"
	"Rental_Mobil/model/web"
	"Rental_Mobil/repository/users"
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
)

type UserServiceImpl struct {
	UserRepository users.UserRepository
	DB             *sql.DB
	validate       *validator.Validate
}

func NewUserServiceImpl(userRepository users.UserRepository, DB *sql.DB, validate *validator.Validate) *UserServiceImpl {
	return &UserServiceImpl{UserRepository: userRepository, DB: DB, validate: validate}
}

func (service UserServiceImpl) GetAll(ctx context.Context) []dto.UserResponse {
	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	users := service.UserRepository.GetAll(ctx, tx)

	var userResponse []dto.UserResponse
	for _, user := range users {
		userResponse = append(userResponse, dto.UserResponse{
			ID:    user.Id,
			Name:  user.Name,
			Email: user.Email,
		})
	}

	return userResponse
}

func (service UserServiceImpl) Get(ctx context.Context, userId int) dto.UserResponse {
	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	user, err := service.UserRepository.Get(ctx, tx, userId)
	helpers.PanicIfError(err)

	return user.ToDtoResponse()
}

func (service UserServiceImpl) Create(ctx context.Context, request web.UserRequest) error {
	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)

	defer helpers.CommitOrRollback(tx)

	password, err := helpers.HashPassword(request.Password)
	helpers.PanicIfError(err)

	user := domain.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: password,
	}

	err = service.UserRepository.Save(ctx, tx, user)
	helpers.PanicIfError(err)

	return nil
}

func (service UserServiceImpl) Update(ctx context.Context, request web.UserRequest, userId int) error {
	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	user := domain.User{
		Id:    userId,
		Name:  request.Name,
		Email: request.Email,
	}

	err = service.UserRepository.Update(ctx, tx, userId, user)
	helpers.PanicIfError(err)

	return nil
}

func (service UserServiceImpl) Delete(ctx context.Context, userId int) error {
	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	err = service.UserRepository.Delete(ctx, tx, userId)
	helpers.PanicIfError(err)

	return nil
}

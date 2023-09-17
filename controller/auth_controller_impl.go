package controller

import (
	"Rental_Mobil/helpers"
	"Rental_Mobil/model/web"
	"Rental_Mobil/service/user"
	"context"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type AuthControllerImpl struct {
	AuthService user.UserService
}

func NewAuthControllerImpl(authService user.UserService) *AuthControllerImpl {
	return &AuthControllerImpl{AuthService: authService}
}

func (controller AuthControllerImpl) Login(write http.ResponseWriter, request *http.Request, params httprouter.Params) {
	loginRequest := web.LoginRequest{}
	decode := json.NewDecoder(request.Body)
	err := decode.Decode(&loginRequest)
	helpers.PanicIfError(err)

	user := controller.AuthService.GetByEmail(context.Background(), loginRequest)

	token := helpers.CreateToken(user)

	write.Header().Add("Content-Type", "application-json")
	response := web.FormatResponse{
		Code:   200,
		Status: "Ok",
		Data: map[string]interface{}{
			"user":  user,
			"token": token,
		},
	}

	encode := json.NewEncoder(write)
	err = encode.Encode(&response)
	helpers.PanicIfError(err)
}

func (controller AuthControllerImpl) Register(write http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userRequest := web.UserRequest{}
	decode := json.NewDecoder(request.Body)
	err := decode.Decode(&userRequest)
	helpers.PanicIfError(err)

	err = controller.AuthService.Create(context.Background(), userRequest)
	helpers.PanicIfError(err)

	write.Header().Add("Content-Type", "application-json")
	responseUser := web.FormatResponse{
		Code:   200,
		Status: "ok",
		Data:   userRequest,
	}

	encode := json.NewEncoder(write)
	err = encode.Encode(responseUser)
	helpers.PanicIfError(err)
}

package controller

import (
	"Rental_Mobil/helpers"
	"Rental_Mobil/model/dto"
	"Rental_Mobil/model/web"
	"Rental_Mobil/service/Car"
	"context"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type CarControllerImpl struct {
	carService Car.CarService
}

func NewCarControllerImpl(carService Car.CarService) *CarControllerImpl {
	return &CarControllerImpl{carService: carService}
}

func (controller CarControllerImpl) GetAll(write http.ResponseWriter, request *http.Request, params httprouter.Params) {

	cars := controller.carService.GetAll(context.Background())

	response := web.FormatResponse{
		Code:   http.StatusOK,
		Status: "ok",
		Data:   cars,
	}

	write.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(write)
	err := encoder.Encode(&response)
	helpers.PanicIfError(err)
}

func (controller CarControllerImpl) Create(write http.ResponseWriter, request *http.Request, params httprouter.Params) {
	requestDto := dto.CarRequestDto{}
	decode := json.NewDecoder(request.Body)
	err := decode.Decode(&requestDto)
	helpers.PanicIfError(err)

	controller.carService.Create(context.Background(), requestDto)

	response := web.FormatResponse{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	write.Header().Add("Content-Type", "application/json")
	encode := json.NewEncoder(write)
	err = encode.Encode(&response)
	helpers.PanicIfError(err)
}

func (controller CarControllerImpl) Get(write http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id := params.ByName("id")
	carId, err := strconv.Atoi(id)
	helpers.PanicIfError(err)

	car := controller.carService.Get(context.Background(), carId)

	response := web.FormatResponse{
		Code:   200,
		Status: "Ok",
		Data:   car,
	}

	write.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(write)
	err = encoder.Encode(&response)
	helpers.PanicIfError(err)
}

func (controller CarControllerImpl) Update(write http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id := params.ByName("id")
	cardId, err := strconv.Atoi(id)
	helpers.PanicIfError(err)

	var requestData dto.CarRequestDto
	decode := json.NewDecoder(request.Body)
	err = decode.Decode(&requestData)
	helpers.PanicIfError(err)

	controller.carService.Update(context.Background(), requestData, cardId)

	response := web.FormatResponse{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	write.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(write)
	err = encoder.Encode(&response)
	helpers.PanicIfError(err)

}

func (controller CarControllerImpl) Delete(write http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id := params.ByName("id")
	cardId, err := strconv.Atoi(id)
	helpers.PanicIfError(err)

	controller.carService.Delete(context.Background(), cardId)

	response := web.FormatResponse{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	write.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(write)
	err = encoder.Encode(&response)
	helpers.PanicIfError(err)
}

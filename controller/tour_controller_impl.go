package controller

import (
	"Rental_Mobil/helpers"
	"Rental_Mobil/model/dto"
	"Rental_Mobil/model/web"
	"Rental_Mobil/service/Tour"
	"context"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type TourControllerImpl struct {
	tourService Tour.TourService
}

func NewTourControllerImpl(tourService Tour.TourService) *TourControllerImpl {
	return &TourControllerImpl{tourService: tourService}
}

func (controller TourControllerImpl) GetAll(write http.ResponseWriter, request *http.Request, params httprouter.Params) {
	tours := controller.tourService.GetAll(context.Background())

	response := web.FormatResponse{
		Code:   http.StatusOK,
		Status: "ok",
		Data:   tours,
	}

	write.Header().Add("Content-Type", "application/json")
	encode := json.NewEncoder(write)
	err := encode.Encode(&response)
	helpers.PanicIfError(err)
}

func (controller TourControllerImpl) Get(write http.ResponseWriter, request *http.Request, params httprouter.Params) {
	tourId := params.ByName("id")
	id, err := strconv.Atoi(tourId)
	helpers.PanicIfError(err)

	tour := controller.tourService.Get(context.Background(), id)

	response := web.FormatResponse{
		Code:   http.StatusOK,
		Status: "ok",
		Data:   tour,
	}

	write.Header().Add("Content-Type", "application/json")
	encode := json.NewEncoder(write)
	err = encode.Encode(&response)
	helpers.PanicIfError(err)
}

func (controller TourControllerImpl) Create(write http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var tourRequest dto.TourRequestDto
	decode := json.NewDecoder(request.Body)
	err := decode.Decode(&tourRequest)
	helpers.PanicIfError(err)

	err = controller.tourService.Create(context.Background(), tourRequest)
	helpers.PanicIfError(err)

	response := web.FormatResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}

	write.Header().Add("Content-Type", "application/json")
	encode := json.NewEncoder(write)
	err = encode.Encode(&response)
	helpers.PanicIfError(err)
}

func (controller TourControllerImpl) Update(write http.ResponseWriter, request *http.Request, params httprouter.Params) {
	tourId := params.ByName("id")
	id, err := strconv.Atoi(tourId)
	helpers.PanicIfError(err)

	var tourRequest dto.TourRequestDto
	decode := json.NewDecoder(request.Body)
	err = decode.Decode(&tourRequest)
	helpers.PanicIfError(err)

	err = controller.tourService.Update(context.Background(), tourRequest, id)
	helpers.PanicIfError(err)

	response := web.FormatResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}

	write.Header().Add("Content-Type", "application/json")
	encode := json.NewEncoder(write)
	err = encode.Encode(&response)
	helpers.PanicIfError(err)
}

func (controller TourControllerImpl) Delete(write http.ResponseWriter, request *http.Request, params httprouter.Params) {
	tourId := params.ByName("id")
	id, err := strconv.Atoi(tourId)
	helpers.PanicIfError(err)

	err = controller.tourService.Delete(context.Background(), id)
	helpers.PanicIfError(err)

	response := web.FormatResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}

	write.Header().Add("Content-Type", "application/json")
	encode := json.NewEncoder(write)
	err = encode.Encode(&response)
	helpers.PanicIfError(err)
}

package controller

import (
	"Rental_Mobil/helpers"
	"Rental_Mobil/model/dto"
	"Rental_Mobil/model/web"
	"Rental_Mobil/service/leaseType"
	"context"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"strconv"
)

type LeaseTypeControllerImpl struct {
	leaseTypeService leaseType.LeaseTypeService
}

func NewLeaseTypeControllerImpl(leaseTypeService leaseType.LeaseTypeService) *LeaseTypeControllerImpl {
	return &LeaseTypeControllerImpl{leaseTypeService: leaseTypeService}
}

func (controller LeaseTypeControllerImpl) ListLeaseType(write http.ResponseWriter, request *http.Request, params httprouter.Params) {

	listTypes := controller.leaseTypeService.GetList(context.Background())
	log.Println(listTypes)

	write.Header().Add("Content-Type", "application-json")
	response := web.FormatResponse{
		Code:   http.StatusOK,
		Status: "ok",
		Data:   listTypes,
	}

	write.Header().Add("Content-Type", "application-json")
	encode := json.NewEncoder(write)
	err := encode.Encode(response)
	helpers.PanicIfError(err)
}

func (controller LeaseTypeControllerImpl) GetLeaseType(write http.ResponseWriter, request *http.Request, params httprouter.Params) {
	leaseTypeId := params.ByName("id")
	id, err := strconv.Atoi(leaseTypeId)
	helpers.PanicIfError(err)
	leaseType := controller.leaseTypeService.GET(context.Background(), id)

	response := web.FormatResponse{
		Code:   http.StatusOK,
		Status: "ok",
		Data:   leaseType,
	}

	write.Header().Add("Content-Type", "application-json")
	encode := json.NewEncoder(write)
	err = encode.Encode(&response)
	helpers.PanicIfError(err)
}

func (controller LeaseTypeControllerImpl) CreateLeaseType(write http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var leaseTypeRequest dto.LeaseTypeRequest
	decode := json.NewDecoder(request.Body)
	err := decode.Decode(&leaseTypeRequest)
	helpers.PanicIfError(err)

	err = controller.leaseTypeService.Create(context.Background(), leaseTypeRequest)
	helpers.PanicIfError(err)

	response := web.FormatResponse{
		Code:   http.StatusOK,
		Status: "oke",
		Data:   "success created lease type",
	}

	encode := json.NewEncoder(write)
	err = encode.Encode(&response)
	helpers.PanicIfError(err)
}

func (controller LeaseTypeControllerImpl) UpdateLeaseType(write http.ResponseWriter, request *http.Request, params httprouter.Params) {
	//TODO implement me
	panic("implement me")
}

func (controller LeaseTypeControllerImpl) DeleteLeaseType(write http.ResponseWriter, request *http.Request, params httprouter.Params) {
	//TODO implement me
	panic("implement me")
}

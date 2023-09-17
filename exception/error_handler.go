package exception

import (
	"Rental_Mobil/helpers"
	"Rental_Mobil/model/web"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func ErrorHandler(write http.ResponseWriter, request *http.Request, err interface{}) {

	if notFoundError(write, request, err) {
		return
	}

	if validatorError(write, request, err) {
		return
	}

	internalServerError(write, request, err)
}

func validatorError(write http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		write.Header().Set("Content-type", "application/json")
		write.WriteHeader(http.StatusBadRequest)

		webResponse := web.FormatResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   exception.Error(),
		}

		write.Header().Add("Content-Type", "application/json")
		encoder := json.NewEncoder(write)
		err := encoder.Encode(webResponse)
		helpers.PanicIfError(err)
		return true
	} else {
		return false
	}
}

func notFoundError(write http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)
	fmt.Println(ok)
	if ok {
		write.Header().Set("Content-type", "application/json")
		write.WriteHeader(http.StatusNotFound)

		webResponse := web.FormatResponse{
			Code:   http.StatusNotFound,
			Status: "Not Found",
			Data:   exception.Error,
		}

		write.Header().Add("Content-Type", "application/json")
		encoder := json.NewEncoder(write)
		err := encoder.Encode(webResponse)
		helpers.PanicIfError(err)
		return true
	} else {
		return false
	}
}

func internalServerError(write http.ResponseWriter, request *http.Request, err interface{}) {
	write.Header().Set("Content-type", "application/json")
	write.WriteHeader(http.StatusInternalServerError)

	webResponse := web.FormatResponse{
		Code:   http.StatusInternalServerError,
		Status: "Internal Status Error",
		Data:   err,
	}

	decode := json.NewDecoder(request.Body)
	err = decode.Decode(&webResponse)
	if err != nil {
		panic(err)
	}

}

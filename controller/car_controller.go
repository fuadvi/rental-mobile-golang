package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type CarController interface {
	GetAll(write http.ResponseWriter, request *http.Request, params httprouter.Params)
	Create(write http.ResponseWriter, request *http.Request, params httprouter.Params)
	Get(write http.ResponseWriter, request *http.Request, params httprouter.Params)
	Update(write http.ResponseWriter, request *http.Request, params httprouter.Params)
	Delete(write http.ResponseWriter, request *http.Request, params httprouter.Params)
}

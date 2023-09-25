package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type LeaseTypeController interface {
	ListLeaseType(write http.ResponseWriter, request *http.Request, params httprouter.Params)
	GetLeaseType(write http.ResponseWriter, request *http.Request, params httprouter.Params)
	CreateLeaseType(write http.ResponseWriter, request *http.Request, params httprouter.Params)
	UpdateLeaseType(write http.ResponseWriter, request *http.Request, params httprouter.Params)
	DeleteLeaseType(write http.ResponseWriter, request *http.Request, params httprouter.Params)
}

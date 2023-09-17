package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type AuthController interface {
	login(write http.ResponseWriter, request *http.Request, params httprouter.Params)
	register(write http.ResponseWriter, request *http.Request, params httprouter.Params)
}

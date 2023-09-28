package main

import (
	"Rental_Mobil/app"
	"Rental_Mobil/controller"
	"Rental_Mobil/exception"
	"Rental_Mobil/helpers"
	"Rental_Mobil/middleware"
	"Rental_Mobil/repository/lease_types"
	"Rental_Mobil/repository/users"
	"Rental_Mobil/service/leaseType"
	"Rental_Mobil/service/user"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	router := httprouter.New()
	db := app.NewDB()
	validate := validator.New()
	userRepo := users.NewUserRepositoryImpl()
	userService := user.NewUserServiceImpl(userRepo, db, validate)
	authController := controller.NewAuthControllerImpl(userService)

	router.POST("/api/register", middleware.JWTMiddleware(authController.Register))
	router.POST("/api/login", authController.Login)

	// lease  type
	leaseTypeRepo := lease_types.NewLeaseTypeRepositoryImpl()
	leaseTypeService := leaseType.NewLeaseTypeServiceImpl(leaseTypeRepo, db, validate)
	leaseTypeController := controller.NewLeaseTypeControllerImpl(leaseTypeService)
	router.GET("/api/lease-type", middleware.JWTMiddleware(leaseTypeController.ListLeaseType))
	router.POST("/api/lease-type", middleware.JWTMiddleware(leaseTypeController.CreateLeaseType))
	router.GET("/api/lease-type/:id", middleware.JWTMiddleware(leaseTypeController.GetLeaseType))
	router.PUT("/api/lease-type/:id", middleware.JWTMiddleware(leaseTypeController.UpdateLeaseType))
	router.DELETE("/api/lease-type/:id", middleware.JWTMiddleware(leaseTypeController.DeleteLeaseType))

	router.PanicHandler = exception.ErrorHandler
	server := &http.Server{
		Addr:    "localhost:5050",
		Handler: router,
	}

	err := server.ListenAndServe()
	helpers.PanicIfError(err)
}

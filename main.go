package main

import (
	"log"

	"github.com/elpahlevi/go-rest-api-validation/config"
	"github.com/elpahlevi/go-rest-api-validation/controller"
	"github.com/elpahlevi/go-rest-api-validation/service"
	"github.com/go-playground/validator/v10"
)

func main() {
	inputValidation := config.NewInputValidation(validator.New())
	server := config.NewHTTPServer()

	authService := service.NewAuthService(inputValidation)
	authController := controller.NewAuthController(authService)

	apiRouter := server.Group("/api/v1")
	authController.Router(apiRouter)

	log.Fatal(server.Listen(":8080"))
}

package main

import (
	"log"

	"github.com/elpahlevi/go-rest-api-validation/config"
	"github.com/elpahlevi/go-rest-api-validation/controller"
	"github.com/elpahlevi/go-rest-api-validation/service"
)

func main() {
	server := config.NewHTTPServer()
	inputValidation := config.NewInputValidation()

	authService := service.NewAuthService(inputValidation)
	authController := controller.NewAuthController(authService)

	apiRouter := server.Group("/api/v1")
	authController.Router(apiRouter)

	log.Fatal(server.Listen(":8080"))
}

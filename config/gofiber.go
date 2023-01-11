package config

import (
	"encoding/json"

	"github.com/elpahlevi/go-rest-api-validation/exception"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func NewHTTPServer() *fiber.App {
	config := fiber.Config{
		JSONEncoder:  json.Marshal,
		JSONDecoder:  json.Unmarshal,
		ErrorHandler: exception.NewHTTPErrorHandler,
	}
	app := fiber.New(config)
	app.Use(logger.New())
	app.Use(recover.New())

	return app
}

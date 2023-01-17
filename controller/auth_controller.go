package controller

import (
	"context"
	"net/http"

	"github.com/elpahlevi/go-rest-api-validation/exception"
	"github.com/elpahlevi/go-rest-api-validation/model"
	"github.com/gofiber/fiber/v2"
)

// Agar method yang ada di AuthService bisa digunakan, kita bisa membuat interface untuk menjembataninya
type AuthService interface {
	Register(ctx context.Context, data *model.RegisterInput) model.ServiceResponse
}

type AuthControllerImpl struct {
	AuthService AuthService
}

func NewAuthController(authService AuthService) *AuthControllerImpl {
	return &AuthControllerImpl{AuthService: authService}
}

func (controller *AuthControllerImpl) Router(router fiber.Router) {
	router.Post("/auth/register", controller.Register)
}

func (controller *AuthControllerImpl) Register(ctx *fiber.Ctx) error {
	data := new(model.RegisterInput)
	err := ctx.BodyParser(data)
	if err != nil {
		panic(exception.NewHTTPError(400, err))
	}

	response := controller.AuthService.Register(ctx.Context(), data)
	return ctx.Status(fiber.StatusCreated).JSON(model.WebResponse{
		Code:    fiber.StatusCreated,
		Status:  http.StatusText(fiber.StatusCreated),
		Message: response.Message,
	})
}

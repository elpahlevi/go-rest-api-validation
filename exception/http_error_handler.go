package exception

import (
	"encoding/json"
	"net/http"

	"github.com/elpahlevi/go-rest-api-validation/model"
	"github.com/gofiber/fiber/v2"
)

func NewHTTPErrorHandler(ctx *fiber.Ctx, err error) error {
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	switch e := err.(type) {
	case *HTTPError:
		response := model.WebResponse{
			Code:    e.Code,
			Status:  http.StatusText(e.Code),
			Message: e.Error(),
		}
		return ctx.Status(e.Code).JSON(response)
	case *HTTPInputValidationError:
		response := model.WebErrorInputResponse{
			Code:       fiber.StatusBadRequest,
			Status:     http.StatusText(fiber.StatusBadRequest),
			ErrorField: json.RawMessage(e.Error()),
		}
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	case *fiber.Error:
		response := model.WebResponse{
			Code:    e.Code,
			Status:  http.StatusText(e.Code),
			Message: e.Error(),
		}
		return ctx.Status(e.Code).JSON(response)
	default:
		response := model.WebResponse{
			Code:    fiber.StatusInternalServerError,
			Status:  http.StatusText(fiber.StatusInternalServerError),
			Message: e.Error(),
		}
		return ctx.Status(response.Code).JSON(response)
	}
}

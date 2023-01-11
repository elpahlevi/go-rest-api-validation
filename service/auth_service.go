package service

import (
	"context"

	"github.com/elpahlevi/go-rest-api-validation/config"
	"github.com/elpahlevi/go-rest-api-validation/exception"
	"github.com/elpahlevi/go-rest-api-validation/model"
)

type AuthServiceImpl struct {
	InputValidation *config.InputValidation
}

func NewAuthService(inputValidation *config.InputValidation) *AuthServiceImpl {
	return &AuthServiceImpl{InputValidation: inputValidation}
}

func (service *AuthServiceImpl) Save(ctx context.Context, data *model.RegisterInput) model.ServiceResponse {
	err := service.InputValidation.Validate(data)
	if err != nil {
		panic(exception.NewHTTPInputValidationError(err))
	}

	return model.ServiceResponse{Message: "User Registered"}
}

package test

import (
	"fmt"
	"testing"

	"github.com/elpahlevi/go-rest-api-validation/config"
	"github.com/elpahlevi/go-rest-api-validation/model"
)

func TestInputValidationSuccess(t *testing.T) {
	inputValidation := config.NewInputValidation()
	data := model.RegisterInput{
		Name:     "Mikey",
		Email:    "mikey@mail.com",
		Phone:    "08654322123",
		Password: "cihuy1234",
	}
	err := inputValidation.Validate(data)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Result: Data is valid")
}

func TestInputValidationFailedOne(t *testing.T) {
	inputValidation := config.NewInputValidation()
	data := model.RegisterInput{
		Name:     "Mikey",
		Email:    "mikeymail.com",
		Phone:    "08654322123",
		Password: "cihuy1234",
	}
	err := inputValidation.Validate(data)
	if err != nil {
		t.Fatal("error:", err)
	}
	fmt.Println("Result: Data is valid")
}

func TestInputValidationFailedMoreThanOne(t *testing.T) {
	inputValidation := config.NewInputValidation()
	data := model.RegisterInput{
		Name:     "Mikey",
		Email:    "mikeymail.com",
		Phone:    "08654322123",
		Password: "",
	}
	err := inputValidation.Validate(data)
	if err != nil {
		t.Fatal("error:", err)
	}
	fmt.Println("Result: Data is valid")
}

func TestInputValidationFailedPhoneNumber(t *testing.T) {
	inputValidation := config.NewInputValidation()
	data := model.RegisterInput{
		Name:     "Mikey",
		Email:    "mikey@mail.com",
		Phone:    "5555551234",
		Password: "mike1234",
	}
	err := inputValidation.Validate(data)
	if err != nil {
		t.Fatal("error:", err)
	}
	fmt.Println("Result: Data is valid")
}

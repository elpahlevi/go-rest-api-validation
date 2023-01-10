package model

type RegisterInput struct {
	Name     string `json:"name" validate:"required,min=2"`
	Email    string `json:"email" validate:"required,email"`
	Phone    string `json:"phone" validate:"required,phone"`
	Password string `json:"password" validate:"required,min=8"`
}

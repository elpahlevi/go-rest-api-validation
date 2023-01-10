package model

type ErrorInputResponse struct {
	FieldName string `json:"fieldName"`
	Message   string `json:"message"`
}

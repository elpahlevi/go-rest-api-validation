package config

import (
	"encoding/json"
	"errors"
	"regexp"
	"strings"

	"github.com/elpahlevi/go-rest-api-validation/model"
	"github.com/go-playground/validator/v10"
)

// InputValidation berisi satu field yang akan memanggil struct dari package golang validator
type InputValidation struct {
	Validator *validator.Validate
}

// NewInputValidation merupakan fungsi yang akan mengembalikan struct InputValidation
func NewInputValidation() *InputValidation {
	return &InputValidation{Validator: validator.New()}
}

// Method Validate() adalah logic dari validasi input yang akan kita buat
// Selain itu, kita dapat membuat logic validasi input secara kustom
func (iv *InputValidation) Validate(data interface{}) error {
	// Buat sebuah slice dari struct untuk menampung error input
	// Error dalam input bisa hanya satu atau lebih dari satu
	var errFields []model.ErrorInputResponse

	// Menambah kriteria validasi baru agar dapat memvalidasi nomor telepon
	iv.Validator.RegisterValidation("phone", phoneValidation)
	// Kita akan memanggil method Struct() dari package golang validator untuk memvalidasi input yang diterima
	err := iv.Validator.Struct(data)
	// Jika input yang diterima tidak sesuai dengan kriteria, maka akan menghasilkan error.
	// Kita dapat membuat pesan error secara custom berdasarkan kriteria yang sudah diberikan
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var errField model.ErrorInputResponse
			switch err.Tag() {
			case "email":
				errField.FieldName = strings.ToLower(err.Field())
				errField.Message = "Email format is invalid"
			case "min":
				errField.FieldName = strings.ToLower(err.Field())
				errField.Message = err.Field() + " must be minimum " + err.Param() + " characters"
			case "max":
				errField.FieldName = strings.ToLower(err.Field())
				errField.Message = err.Field() + " maximum allowed is" + err.Param() + " characters"
			case "required":
				errField.FieldName = strings.ToLower(err.Field())
				errField.Message = err.Field() + " cannot be blank"
			case "phone":
				errField.FieldName = strings.ToLower(err.Field())
				errField.Message = "Phone number format is invalid"
			}
			errFields = append(errFields, errField)
		}
	}
	// Kalau tidak ada error, kita bisa mengembalikan nilainya menjadi nil
	if len(errFields) == 0 {
		return nil
	}
	// Jadikan slice dari errFields menjadi JSON Array of Objects
	// Hasilnya bisa digunakan oleh Frontend Developer untuk membuat pesan error input.
	marshaledErr, _ := json.Marshal(errFields)
	return errors.New(string(marshaledErr))
}

func phoneValidation(fl validator.FieldLevel) bool {
	phoneNumber := fl.Field().String()
	if len(phoneNumber) < 10 || len(phoneNumber) > 13 {
		return false
	}
	phoneRegex, _ := regexp.Compile(`^(0|\\+62|062|62)[0-9]+$`)
	return phoneRegex.MatchString(phoneNumber)
}

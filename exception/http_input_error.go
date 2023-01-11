package exception

type HTTPInputValidationError struct {
	Err error
}

func NewHTTPInputValidationError(err error) error {
	return &HTTPInputValidationError{Err: err}
}

func (e *HTTPInputValidationError) Error() string {
	return e.Err.Error()
}

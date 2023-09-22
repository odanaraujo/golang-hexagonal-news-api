package exception

import "net/http"

type Exception struct {
	Message string   `json:"message"`
	Err     string   `json:"error"`
	Code    int      `json:"code"`
	Causes  []Causes `json:"causes"`
}

type Causes struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (r *Exception) Error() string {
	return r.Message
}

func NewException(message, err string, code int, causes []Causes) *Exception {
	return &Exception{
		Message: message,
		Err:     err,
		Code:    code,
		Causes:  causes,
	}
}

func NewBadRequestError(message string) *Exception {
	return &Exception{
		Message: message,
		Err:     "bad request",
		Code:    http.StatusBadRequest,
	}
}

func NewBadRequestValidationError(message string, causes []Causes) *Exception {
	return &Exception{
		Message: message,
		Err:     "bad request",
		Code:    http.StatusBadRequest,
		Causes:  causes,
	}
}

func NewInternalServerError(message string) *Exception {
	return &Exception{
		Message: message,
		Err:     "internal server error",
		Code:    http.StatusInternalServerError,
	}
}

func NewNotFoundError(message string) *Exception {
	return &Exception{
		Message: message,
		Err:     "not found",
		Code:    http.StatusNotFound,
	}
}

func NewForbiddenError(message string) *Exception {
	return &Exception{
		Message: message,
		Err:     "forbidden",
		Code:    http.StatusForbidden,
	}
}
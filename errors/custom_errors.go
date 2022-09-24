package errors

import (
	"errors"
	"net/http"
)

var BadRequestError error
var InternalServerError error

var AppErrors map[error]int

func init() {
	AppErrors = make(map[error]int)

	BadRequestError = errors.New("bad request")
	AppErrors[BadRequestError] = http.StatusBadRequest

	InternalServerError = errors.New("something is wrong, please try again later")
	AppErrors[InternalServerError] = http.StatusInternalServerError

}

func GetErrorCode(err error) int {
	code, ok := AppErrors[err]
	if !ok {
		return http.StatusInternalServerError
	} else {
		return code
	}
}

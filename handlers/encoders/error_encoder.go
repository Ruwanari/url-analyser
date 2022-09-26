package encoders

import (
	"context"
	"html/template"
	"log"
	"net/http"
	"path"
	"web-page-analyser/entities/response_schemas"
	app_errors "web-page-analyser/errors"
)

const ErrorEncodersLogPrefix = "delivery-data-cacher-api.http-encoders.error-encoder"

type CustomError struct {
	Message string
	Code    int
}

func (e *CustomError) Error() string {
	return e.Message
}

//CustomErrorEncoder encodes and renders error responses to error template.
func CustomErrorEncoder(ctx context.Context, err error, w http.ResponseWriter) {

	errs := response_schemas.ErrorResponse{
		Code:    app_errors.GetErrorCode(err),
		Message: err.Error(),
	}

	var errorResponse response_schemas.ErrorResponses
	errorResponse.Errors = append(errorResponse.Errors, errs)

	fp := path.Join("static", "error_response_template.gohtml")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		log.Printf("Error parsing template ctx : %v, err : %v ", ctx, err.Error())
		return
	}

	err = tmpl.Execute(w, response_schemas.ErrorResponseWrapper(errorResponse))
	if err != nil {
		log.Printf("Error executing template ctx : %v, err : %v ", ctx, err.Error())
		return
	}
}

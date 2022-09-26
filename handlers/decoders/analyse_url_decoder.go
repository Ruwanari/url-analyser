package decoders

import (
	"context"
	transportHttp "github.com/go-kit/kit/transport/http"
	"log"
	"net/http"
	"web-page-analyser/entities/request_schemas"
	app_errors "web-page-analyser/errors"
)

//AnalyseUrlDecoder decodes the http request to a readable struct while checking if the request is valid.
func AnalyseUrlDecoder() transportHttp.DecodeRequestFunc {
	return func(ctx context.Context, r *http.Request) (request interface{}, err error) {
		req := request_schemas.AnalyseUrlRequest{}
		err = r.ParseForm()
		if err != nil {
			log.Printf("Cannot parse form ctx: %v err : %v", ctx, err)
			return nil, app_errors.BadRequestError
		}
		req.Url = r.FormValue("url")
		if req.Url == "" {
			log.Printf("Empty url : %v ctx: %v", app_errors.BadRequestError, ctx)
			return nil, app_errors.BadRequestError
		}
		return req, err
	}
}

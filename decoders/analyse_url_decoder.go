package decoders

import (
	"context"
	transportHttp "github.com/go-kit/kit/transport/http"
	"log"
	"net/http"
	"web-page-analyser/request_schemas"
)

func AnalyseUrlDecoder() transportHttp.DecodeRequestFunc {
	return func(ctx context.Context, r *http.Request) (request interface{}, err error) {
		req := request_schemas.AnalyseUrlRequest{}
		err = r.ParseForm()
		if err != nil {
			log.Printf("Cannot parse form")
		}
		req.Url = r.FormValue("url")
		return req, err
	}
}

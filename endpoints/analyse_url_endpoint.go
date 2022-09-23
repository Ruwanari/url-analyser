package endpoints

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-kit/kit/endpoint"
	"log"
	"web-page-analyser/request_schemas"
	"web-page-analyser/usecases"
)

func AnalyseUrl() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		analyseUrlRequest, ok := request.(request_schemas.AnalyseUrlRequest)
		if !ok {
			err = errors.New("bad request")
			return
		}
		marshalledRequest, _ := json.Marshal(request)
		log.Printf("Processed request %s", marshalledRequest)

		response, err = usecases.AnalyseUrlUsecase(analyseUrlRequest.Url)
		if err != nil {
			log.Printf("Could not analyse url : %v , error : %v", analyseUrlRequest.Url, err)
			return
		}
		log.Printf("Response %s", response)
		return

	}
}

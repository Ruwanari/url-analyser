package endpoints

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"log"
	"web-page-analyser/entities/request_schemas"
	app_errors "web-page-analyser/errors"
	"web-page-analyser/usecases"
)

/*AnalyseUrl endpoint accepts a url as a request parameter and analyses the content of the web page directed
by the url.*/
func AnalyseUrl() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		analyseUrlRequest, ok := request.(request_schemas.AnalyseUrlRequest)
		//add ctx
		if !ok {
			log.Printf("Error in Request")
			return nil, app_errors.BadRequestError
		}

		response, err = usecases.AnalyseUrlUsecase(analyseUrlRequest.Url)
		if err != nil {
			log.Printf("Could not analyse url : %v , error : %v", analyseUrlRequest.Url, err)
			return nil, app_errors.InternalServerError
		}
		return

	}
}

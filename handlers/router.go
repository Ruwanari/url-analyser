package handlers

import (
	httpTransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
	"web-page-analyser/decoders"
	"web-page-analyser/encoders"
	"web-page-analyser/endpoints"
)

func GetRoutes() *mux.Router {

	r := mux.NewRouter()

	r.Handle("/v1.0/analyse-url", httpTransport.
		NewServer(
			endpoints.AnalyseUrl(),
			decoders.AnalyseUrlDecoder(),
			encoders.EncodeSuccessPayloadResponse,
		)).Methods(http.MethodPost)

	return r
}

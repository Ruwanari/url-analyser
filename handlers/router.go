package handlers

import (
	httpTransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
	"web-page-analyser/handlers/decoders"
	encoders2 "web-page-analyser/handlers/encoders"
	"web-page-analyser/handlers/endpoints"
)

var serverOptions = []httpTransport.ServerOption{
	httpTransport.ServerErrorEncoder(encoders2.CustomErrorEncoder),
}

func GetRoutes() *mux.Router {

	r := mux.NewRouter()

	r.Handle("/v1.0/analyse-url", httpTransport.
		NewServer(
			endpoints.AnalyseUrl(),
			decoders.AnalyseUrlDecoder(),
			encoders2.EncodeSuccessPayloadResponse,
			serverOptions...,
		)).Methods(http.MethodPost)

	return r
}

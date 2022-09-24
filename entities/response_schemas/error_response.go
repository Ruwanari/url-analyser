package response_schemas

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ErrorResponses struct {
	Errors []ErrorResponse `json:"errors"`
}

type ErrorResponseWrapper struct {
	Errors []ErrorResponse `json:"errors"`
}

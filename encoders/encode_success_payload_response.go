package encoders

import (
	"context"
	"encoding/json"
	"net/http"
)

func EncodeSuccessPayloadResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(response)
}

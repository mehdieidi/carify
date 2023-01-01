package transport

import (
	"context"
	"encoding/json"
	"net/http"

	"back/protocol/derror"
)

type Response struct {
	Data any `json:"data"`
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, res any) error {
	setContentType(w, JSONContentType)
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(res)
}

type errorResponse struct {
	Error string `json:"error"`
}

func EncodeError(_ context.Context, w http.ResponseWriter, err error) {
	setContentType(w, JSONContentType)

	f, ok := errs[err]
	if !ok {
		f = errs[derror.ErrInternalServer]
	}
	f(w)

	json.NewEncoder(w).Encode(errorResponse{Error: err.Error()})
}

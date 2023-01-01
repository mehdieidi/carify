package transport

import (
	"context"
	"encoding/json"
	"net/http"

	"back/protocol/derror"
)

func DecodeJSONRequest[T any](_ context.Context, r *http.Request) (any, error) {
	var req T

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, derror.ErrInvalidRequest
	}

	return req, nil
}

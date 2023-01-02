package transport

import (
	"net/http"

	"back/protocol/derror"
)

var errs = map[error]func(http.ResponseWriter){
	// status internal server error.
	derror.ErrInternalServer: status(http.StatusInternalServerError),

	// status unauthorized.

	// status forbidden.

	// status conflict.

	// status bad request.
	derror.ErrInvalidRequest: status(http.StatusBadRequest),

	// status not found.

}

func status(code int) func(http.ResponseWriter) {
	return func(w http.ResponseWriter) {
		w.WriteHeader(code)
	}
}

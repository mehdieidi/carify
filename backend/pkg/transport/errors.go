package transport

import (
	"net/http"

	"back/protocol/derror"

	kitjwt "bitbucket.imenaria.org/tool/toolkit/auth/jwt"
)

var errs = map[error]func(http.ResponseWriter){
	// status internal server error.
	derror.ErrInternalServer: status(http.StatusInternalServerError),

	// status unauthorized.
	kitjwt.ErrTokenContextMissing: status(http.StatusUnauthorized),
	kitjwt.ErrTokenExpired:        status(http.StatusUnauthorized),
	kitjwt.ErrTokenInvalid:        status(http.StatusUnauthorized),
	kitjwt.ErrTokenMalformed:      status(http.StatusUnauthorized),
	kitjwt.ErrTokenNotActive:      status(http.StatusUnauthorized),

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

package derror

import "errors"

var (
	ErrUnknownDoctorServiceType = errors.New("unknown doctor service type")
)

var (
	ErrInvalidRequest = errors.New("invalid request")
)

var (
	ErrUnexpected     = errors.New("unexpected error")
	ErrInternalServer = errors.New("internal server error")
)

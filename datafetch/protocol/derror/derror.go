package derror

import "errors"

var (
	ErrUnknownCar = errors.New("unknown car")
	ErrUnexpected = errors.New("unexpected error")
)

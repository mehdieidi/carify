package transport

import "context"

// ErrorHandler receives a transport error to be processed for diagnostic purposes.
type ErrorHandler interface {
	Handle(ctx context.Context, err error)
}

// ErrorHandlerFunc type is an adapter to allow the use of
// ordinary
type ErrorHandlerFunc func(ctx context.Context, err error)

// Handle calls f(ctx, err).
func (f ErrorHandlerFunc) Handle(ctx context.Context, err error) {
	f(ctx, err)
}

// NopErrHandler returns a error handler that does nothing.
func NopErrorHandler() ErrorHandlerFunc {
	return func(ctx context.Context, err error) {}
}

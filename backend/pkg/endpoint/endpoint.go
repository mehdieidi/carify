package endpoint

import "context"

// Endpoint is the fundamental building block of servers and clients.
type Endpoint func(ctx context.Context, req any) (res any, err error)

// Nop is an endpoint that does nothing and returns a nil error.
func Nop(ctx context.Context, req any) (any, error) { return struct{}{}, nil }

// Middleware is a chainable behavior modifier for endpoints.
type Middleware func(Endpoint) Endpoint

// Chain is a helper function for composing middlewares. Requests will
// traverse them in the order they're declared. That is, the first middleware
// is treated as the outermost middleware.
func Chain(outer Middleware, others ...Middleware) Middleware {
	return func(next Endpoint) Endpoint {
		for i := len(others) - 1; i >= 0; i-- {
			next = others[i](next)
		}
		return outer(next)
	}
}

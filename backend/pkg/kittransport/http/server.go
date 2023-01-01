package http

import (
	"context"
	"net/http"

	"back/pkg/endpoint"
	transport "back/pkg/kittransport"
)

type Server struct {
	e            endpoint.Endpoint
	dec          RequestDecoderFunc
	enc          ResponseEncoderFunc
	before       []RequestFunc
	after        []ServerResponseFunc
	encodeError  ErrorEncoderFunc
	errorHandler transport.ErrorHandler
}

// NewServer constructs a new server, which implements http.Handler and wraps
// the provided endpoint.
func NewServer(
	e endpoint.Endpoint,
	dec RequestDecoderFunc,
	enc ResponseEncoderFunc,
	encerr ErrorEncoderFunc,
	options ...ServerOption,
) *Server {
	s := &Server{
		e:            e,
		dec:          dec,
		enc:          enc,
		encodeError:  encerr,
		errorHandler: transport.NopErrorHandler(),
	}
	for _, option := range options {
		option(s)
	}
	return s
}

// ServerOption sets an optional parameter for servers.
type ServerOption func(*Server)

// ServerBefore functions are executed on the HTTP request object before the
// request is decoded.
func ServerBefore(before ...RequestFunc) ServerOption {
	return func(s *Server) { s.before = append(s.before, before...) }
}

// ServerAfter functions are executed on the HTTP response writer after the
// endpoint is invoked, but before interface{}thing is written to the client.
func ServerAfter(after ...ServerResponseFunc) ServerOption {
	return func(s *Server) { s.after = append(s.after, after...) }
}

// ServerErrorHandler is used to handle non-terminal errors. By default, non-terminal errors
// are ignored.
func ServerErrorHandler(errorHandler transport.ErrorHandler) ServerOption {
	return func(s *Server) { s.errorHandler = errorHandler }
}

// ServeHTTP implements http.Handler.
func (s Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	for _, f := range s.before {
		ctx = f(ctx, r)
	}

	request, err := s.dec(ctx, r)
	if err != nil {
		s.errorHandler.Handle(ctx, err)
		s.encodeError(ctx, w, err)
		return
	}

	response, err := s.e(ctx, request)
	if err != nil {
		s.errorHandler.Handle(ctx, err)
		s.encodeError(ctx, w, err)
		return
	}

	for _, f := range s.after {
		ctx = f(ctx, w)
	}

	if err := s.enc(ctx, w, response); err != nil {
		s.errorHandler.Handle(ctx, err)
		s.encodeError(ctx, w, err)
		return
	}
}

// NopRequestDecoder is a DecodeRequestFunc that can be used for requests that do not
// need to be decoded, and returns simply nil, nil.
func NopRequestDecoder(ctx context.Context, r *http.Request) (any, error) {
	return nil, nil
}

// If an error value implements Headerer,
// the Headers can be used when encoding the error.
type Headerer interface {
	Headers() http.Header
}

// If an error value implements StatusCoder,
// the StatusCode can be used when encoding the error.
type StatusCoder interface {
	StatusCode() int
}

package http

import (
	"context"
	"net/http"
)

// RequestEncoderFunc encodes the passed request object into the HTTP request
// object.
type RequestEncoderFunc func(context.Context, *http.Request, any) error

// RequestDecoderFunc extracts a request object from an HTTP
// request object.
type RequestDecoderFunc func(context.Context, *http.Request) (any, error)

// RequestCreatorFunc creates an outgoing HTTP request based on the passed
// request object.
type RequestCreatorFunc func(context.Context, any) (*http.Request, error)

// ResponseEncoderFunc encodes the passed response object to the HTTP response
// writer.
type ResponseEncoderFunc func(context.Context, http.ResponseWriter, any) error

// ResponseDecoderFunc extracts a response object from an HTTP
// response object.
type ResponseDecoderFunc func(context.Context, *http.Response) (any, error)

// ErrorEncoderFunc is for encoding an error to the ResponseWriter.
type ErrorEncoderFunc func(context.Context, http.ResponseWriter, error)

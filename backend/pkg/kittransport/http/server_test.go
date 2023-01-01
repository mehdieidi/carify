package http_test

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"back/pkg/endpoint"
	httptransport "back/pkg/kittransport/http"
)

func TestServerBadDecode(t *testing.T) {
	handler := httptransport.NewServer(
		func(context.Context, any) (any, error) { return struct{}{}, nil },
		func(context.Context, *http.Request) (any, error) { return struct{}{}, errors.New("bigbang") },
		func(context.Context, http.ResponseWriter, any) error { return nil },
		func(_ context.Context, w http.ResponseWriter, err error) {
			w.WriteHeader(http.StatusInternalServerError)
		},
	)
	server := httptest.NewServer(handler)
	defer server.Close()
	resp, _ := http.Get(server.URL)
	if want, have := http.StatusInternalServerError, resp.StatusCode; want != have {
		t.Errorf("want %d, have %d", want, have)
	}
}

func TestServerBadEndpoint(t *testing.T) {
	handler := httptransport.NewServer(
		func(context.Context, any) (any, error) { return struct{}{}, errors.New("dang") },
		func(context.Context, *http.Request) (any, error) { return struct{}{}, nil },
		func(context.Context, http.ResponseWriter, any) error { return nil },
		func(_ context.Context, w http.ResponseWriter, err error) {
			w.WriteHeader(http.StatusInternalServerError)
		},
	)
	server := httptest.NewServer(handler)
	defer server.Close()
	resp, _ := http.Get(server.URL)
	if want, have := http.StatusInternalServerError, resp.StatusCode; want != have {
		t.Errorf("want %d, have %d", want, have)
	}
}

func TestServerBadEncode(t *testing.T) {
	handler := httptransport.NewServer(
		func(context.Context, any) (any, error) { return struct{}{}, nil },
		func(context.Context, *http.Request) (any, error) { return struct{}{}, nil },
		func(context.Context, http.ResponseWriter, any) error { return errors.New("dang") },
		func(_ context.Context, w http.ResponseWriter, err error) {
			w.WriteHeader(http.StatusInternalServerError)
		},
	)
	server := httptest.NewServer(handler)
	defer server.Close()
	resp, _ := http.Get(server.URL)
	if want, have := http.StatusInternalServerError, resp.StatusCode; want != have {
		t.Errorf("want %d, have %d", want, have)
	}
}

func TestServerErrorEncoder(t *testing.T) {
	errTeapot := errors.New("teapot")
	code := func(err error) int {
		if errors.Is(err, errTeapot) {
			return http.StatusTeapot
		}
		return http.StatusInternalServerError
	}
	handler := httptransport.NewServer(
		func(context.Context, any) (any, error) { return struct{}{}, errTeapot },
		func(context.Context, *http.Request) (any, error) { return struct{}{}, nil },
		func(context.Context, http.ResponseWriter, any) error { return nil },
		func(_ context.Context, w http.ResponseWriter, err error) { w.WriteHeader(code(err)) },
	)
	server := httptest.NewServer(handler)
	defer server.Close()
	resp, _ := http.Get(server.URL)
	if want, have := http.StatusTeapot, resp.StatusCode; want != have {
		t.Errorf("want %d, have %d", want, have)
	}
}

func TestMultipleServerBefore(t *testing.T) {
	var (
		headerKey    = "X-Henlo-Lizer"
		headerVal    = "Helllo you stinky lizard"
		statusCode   = http.StatusTeapot
		responseBody = "go eat a fly ugly\n"
		done         = make(chan struct{})
	)
	handler := httptransport.NewServer(
		endpoint.Nop,
		func(context.Context, *http.Request) (any, error) {
			return struct{}{}, nil
		},
		func(_ context.Context, w http.ResponseWriter, _ any) error {
			w.Header().Set(headerKey, headerVal)
			w.WriteHeader(statusCode)
			w.Write([]byte(responseBody))
			return nil
		},
		func(context.Context, http.ResponseWriter, error) {},
		httptransport.ServerBefore(func(ctx context.Context, r *http.Request) context.Context {
			ctx = context.WithValue(ctx, "one", 1)

			return ctx
		}),
		httptransport.ServerBefore(func(ctx context.Context, r *http.Request) context.Context {
			if _, ok := ctx.Value("one").(int); !ok {
				t.Error("Value was not set properly when multiple ServerBefores are used")
			}

			close(done)
			return ctx
		}),
	)

	server := httptest.NewServer(handler)
	defer server.Close()
	go http.Get(server.URL)

	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatal("timeout waiting for finalizer")
	}
}

func TestMultipleServerAfter(t *testing.T) {
	var (
		headerKey    = "X-Henlo-Lizer"
		headerVal    = "Helllo you stinky lizard"
		statusCode   = http.StatusTeapot
		responseBody = "go eat a fly ugly\n"
		done         = make(chan struct{})
	)
	handler := httptransport.NewServer(
		endpoint.Nop,
		func(context.Context, *http.Request) (any, error) {
			return struct{}{}, nil
		},
		func(_ context.Context, w http.ResponseWriter, _ any) error {
			w.Header().Set(headerKey, headerVal)
			w.WriteHeader(statusCode)
			w.Write([]byte(responseBody))
			return nil
		},
		func(context.Context, http.ResponseWriter, error) {},
		httptransport.ServerAfter(func(ctx context.Context, w http.ResponseWriter) context.Context {
			ctx = context.WithValue(ctx, "one", 1)

			return ctx
		}),
		httptransport.ServerAfter(func(ctx context.Context, w http.ResponseWriter) context.Context {
			if _, ok := ctx.Value("one").(int); !ok {
				t.Error("Value was not set properly when multiple ServerAfters are used")
			}

			close(done)
			return ctx
		}),
	)

	server := httptest.NewServer(handler)
	defer server.Close()
	go http.Get(server.URL)

	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatal("timeout waiting for finalizer")
	}
}

type enhancedError struct{}

func (e enhancedError) Error() string                { return "enhanced error" }
func (e enhancedError) StatusCode() int              { return http.StatusTeapot }
func (e enhancedError) MarshalJSON() ([]byte, error) { return []byte(`{"err":"enhanced"}`), nil }
func (e enhancedError) Headers() http.Header         { return http.Header{"X-Enhanced": []string{"1"}} }

func TestEnhancedError(t *testing.T) {
	handler := httptransport.NewServer(
		func(context.Context, any) (any, error) { return nil, enhancedError{} },
		func(context.Context, *http.Request) (any, error) { return struct{}{}, nil },
		func(_ context.Context, w http.ResponseWriter, _ any) error { return nil },
		func(_ context.Context, w http.ResponseWriter, err error) {
			contentType, body := "text/plain; charset=utf-8", []byte(err.Error())
			if marshaler, ok := err.(json.Marshaler); ok {
				if jsonBody, marshalErr := marshaler.MarshalJSON(); marshalErr == nil {
					contentType, body = "application/json; charset=utf-8", jsonBody
				}
			}

			w.Header().Set("Content-Type", contentType)
			enhancedErr := err.(enhancedError)
			for key, values := range enhancedErr.Headers() {
				for _, val := range values {
					w.Header().Add(key, val)
				}
			}
			w.WriteHeader(enhancedErr.StatusCode())
			w.Write(body)
		},
	)

	server := httptest.NewServer(handler)
	defer server.Close()

	resp, err := http.Get(server.URL)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if want, have := http.StatusTeapot, resp.StatusCode; want != have {
		t.Errorf("StatusCode: want %d, have %d", want, have)
	}
	if want, have := "1", resp.Header.Get("X-Enhanced"); want != have {
		t.Errorf("X-Enhanced: want %q, have %q", want, have)
	}
	buf, _ := ioutil.ReadAll(resp.Body)
	if want, have := `{"err":"enhanced"}`, strings.TrimSpace(string(buf)); want != have {
		t.Errorf("Body: want %s, have %s", want, have)
	}
}

func TestNoOpRequestDecoder(t *testing.T) {
	resw := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Error("Failed to create request")
	}
	handler := httptransport.NewServer(
		func(ctx context.Context, request any) (any, error) {
			if request != nil {
				t.Error("Expected nil request in endpoint when using NopRequestDecoder")
			}
			return nil, nil
		},
		httptransport.NopRequestDecoder,
		func(_ context.Context, w http.ResponseWriter, _ any) error {
			w.WriteHeader(http.StatusOK)
			return nil
		},
		func(context.Context, http.ResponseWriter, error) {},
	)
	handler.ServeHTTP(resw, req)
	if resw.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, resw.Code)
	}
}

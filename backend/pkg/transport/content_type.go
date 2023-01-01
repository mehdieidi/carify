package transport

import "net/http"

type contentType uint8

const (
	JSONContentType contentType = iota
)

func (c contentType) String() string {
	switch c {
	case JSONContentType:
		return "application/json; charset=utf-8"
	default:
		return ""
	}
}

func setContentType(w http.ResponseWriter, c contentType) {
	w.Header().Set("Content-Type", c.String())
}

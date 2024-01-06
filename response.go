package httx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Response is a prepared response from httx to serve
type Response struct {
	Headers  map[string]string
	Cookies  []*http.Cookie
	Redirect string
	File     string
	Body     string
}

// fileResponse creates a new response for a static file
func fileResponse(path string) *Response {
	return &Response{File: path}
}

// loadResponse from an io reader, assumed json encoding
func loadResponse(r io.Reader) (res *Response) {
	var buf bytes.Buffer
	tee := io.TeeReader(r, &buf)
	if err := json.NewDecoder(tee).Decode(&res); err != nil {
		res = &Response{Body: buf.String()}
	}
	return res
}

// ServeHTTP fulfilling the http.Handler interface in net/http
func (res *Response) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// If static file, serve statically using go
	if res.File != "" {
		http.ServeFile(w, r, res.File)
		return
	}
	// Set outbound headers on response writer
	for key, val := range res.Headers {
		w.Header().Set(key, val)
	}
	// Set cookies that were to the request
	for _, cookie := range res.Cookies {
		http.SetCookie(w, cookie)
	}
	// Redirect if necessary
	if res.Redirect != "" {
		http.Redirect(w, r, res.Redirect, http.StatusFound)
		return
	}
	// Otherwise print the body given to request
	fmt.Fprint(w, res.Body)
}

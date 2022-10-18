package httpclient

import (
	"net/http"
	"strings"
)

type (
	Transport interface {
		RoundTrip(r *http.Request) (*http.Response, error)
	}
	implTransport struct {
		header    http.Header
		transport http.RoundTripper
	}
)

func NewTransport(opts Options) Transport {
	return &implTransport{
		header:    opts.Header,
		transport: opts.Transport,
	}
}

func (i *implTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	for k, v := range i.header {
		r.Header.Set(k, strings.Join(v, "; "))
	}
	return i.transport.RoundTrip(r)
}

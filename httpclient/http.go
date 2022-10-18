package httpclient

import (
	"net/http"
	"time"
)

type (
	Options struct {
		Timeout   time.Duration
		Header    http.Header
		Transport http.RoundTripper
	}
)

func New(opts Options) *http.Client {
	if opts.Timeout == 0 {
		opts.Timeout = 5 * time.Second
	}
	if opts.Transport == nil {
		opts.Transport = http.DefaultTransport
	}
	return &http.Client{
		Timeout:   opts.Timeout,
		Transport: NewTransport(opts),
	}
}

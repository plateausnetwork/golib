package httphelper

import (
	"context"
	"net/http"
	"strings"
	"time"
)

var defaultHeaders = map[string]string{
	"Content-Type": "application/json",
}

type (
	Options struct {
		BaseURL    string
		HttpClient *http.Client
		Timeout    time.Duration
		Header     map[string]string
	}
	implClient struct {
		baseURL    string
		header     map[string]string
		httpClient *http.Client
	}
	Client interface {
		Get(ctx context.Context, endpoint string, requestBody interface{}, dest interface{}) *ResponseFail
		Post(ctx context.Context, endpoint string, requestBody interface{}, dest interface{}) *ResponseFail
		Patch(ctx context.Context, endpoint string, requestBody interface{}, dest interface{}) *ResponseFail
		Put(ctx context.Context, endpoint string, requestBody interface{}, dest interface{}) *ResponseFail
		Delete(ctx context.Context, endpoint string, requestBody interface{}, dest interface{}) *ResponseFail
		SetAuthBasicToHeader(user, password string)
	}
	ResponseFail struct {
		Err        error
		StatusCode int
		Body       []byte
		Data       map[string]interface{}
		Header     http.Header
	}
)

func New(opts Options) Client {
	if opts.BaseURL != "" && !strings.HasSuffix(opts.BaseURL, "/") {
		opts.BaseURL += "/"
	}
	if opts.Header == nil {
		opts.Header = defaultHeaders
	}
	httpClient := opts.HttpClient
	if opts.HttpClient == nil {
		httpClient = http.DefaultClient
	}
	if opts.Timeout > 0 {
		httpClient.Timeout = opts.Timeout
	}
	return &implClient{
		httpClient: httpClient,
		baseURL:    opts.BaseURL,
		header:     opts.Header,
	}
}

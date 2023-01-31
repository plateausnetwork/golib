package httphelper

import (
	"context"
	"encoding/base64"
	"net/http"
	"strings"
)

func (i *implClient) do(method string, request Request) (*http.Response, error) {
	bodyReader, err := GetBodyReader(request.Body)
	if err != nil {
		return nil, err
	}
	ctx := request.Context
	if ctx == nil {
		ctx = context.Background()
	}
	endpoint := strings.TrimPrefix(request.Endpoint, "/")
	req, err := http.NewRequestWithContext(ctx, method, i.baseURL+endpoint, bodyReader)
	if err != nil {
		return nil, err
	}
	for key, value := range i.header {
		req.Header.Set(key, value)
	}
	for key, value := range request.Header {
		req.Header.Set(key, value)
	}
	return i.httpClient.Do(req)
}

func (i *implClient) Get(request Request) *ResponseFail {
	response, err := i.do(http.MethodGet, request)
	if err != nil {
		return &ResponseFail{
			Err:        err,
			StatusCode: response.StatusCode,
			Header:     response.Header,
		}
	}
	return DecodeResponse(response, request.Destination)
}

func (i *implClient) Post(request Request) *ResponseFail {
	response, err := i.do(http.MethodPost, request)
	if err != nil {
		return &ResponseFail{
			Err:        err,
			StatusCode: response.StatusCode,
			Header:     response.Header,
		}
	}
	return DecodeResponse(response, request.Destination)
}

func (i *implClient) Patch(request Request) *ResponseFail {
	response, err := i.do(http.MethodPatch, request)
	if err != nil {
		return &ResponseFail{
			Err:        err,
			StatusCode: response.StatusCode,
			Header:     response.Header,
		}
	}
	return DecodeResponse(response, request.Destination)
}

func (i *implClient) Put(request Request) *ResponseFail {
	response, err := i.do(http.MethodPut, request)
	if err != nil {
		return &ResponseFail{
			Err:        err,
			StatusCode: response.StatusCode,
			Header:     response.Header,
		}
	}
	return DecodeResponse(response, request.Destination)
}

func (i *implClient) Delete(request Request) *ResponseFail {
	response, err := i.do(http.MethodDelete, request)
	if err != nil {
		return &ResponseFail{
			Err:        err,
			StatusCode: response.StatusCode,
			Header:     response.Header,
		}
	}
	return DecodeResponse(response, request.Destination)
}

func (i implClient) SetAuthBasicToHeader(user, password string) {
	auth := []byte(user + ":" + password)
	i.header["Authorization"] = "Basic " + base64.StdEncoding.EncodeToString(auth)
}

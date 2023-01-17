package httphelper

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

func getBodyReader(body interface{}) (io.Reader, error) {
	if body == nil {
		return nil, nil
	}
	if bodyReader, ok := body.(strings.Reader); ok {
		return &bodyReader, nil
	}
	if bodyReader, ok := body.(io.Reader); ok {
		return bodyReader, nil
	}
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(bodyBytes), nil
}

func (i *implClient) do(method string, request Request) (*http.Response, error) {
	bodyReader, err := getBodyReader(request.Body)
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

func (i *implClient) Decode(response *http.Response, dest interface{}) *ResponseFail {
	defer response.Body.Close()
	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return &ResponseFail{
			Err:        err,
			StatusCode: response.StatusCode,
			Header:     response.Header,
		}
	}
	if response.StatusCode < 200 || response.StatusCode > 299 {
		var decodeData map[string]interface{}
		if err := json.Unmarshal(bodyBytes, &decodeData); err != nil {
			return &ResponseFail{
				Err:        err,
				StatusCode: response.StatusCode,
				Header:     response.Header,
				Body:       bodyBytes,
			}
		}
		return &ResponseFail{
			StatusCode: response.StatusCode,
			Data:       decodeData,
			Header:     response.Header,
			Body:       bodyBytes,
		}
	}
	if dest != nil {
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return &ResponseFail{
				Err:        err,
				StatusCode: response.StatusCode,
				Header:     response.Header,
				Body:       bodyBytes,
			}
		}
	}
	return nil
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
	return i.Decode(response, request.Destination)
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
	return i.Decode(response, request.Destination)
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
	return i.Decode(response, request.Destination)
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
	return i.Decode(response, request.Destination)
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
	return i.Decode(response, request.Destination)
}

func (i implClient) SetAuthBasicToHeader(user, password string) {
	auth := []byte(user + ":" + password)
	i.header["Authorization"] = "Basic " + base64.StdEncoding.EncodeToString(auth)
}

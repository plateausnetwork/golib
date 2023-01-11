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

func (i *implClient) do(ctx context.Context, method, endpoint string, requestBody interface{}) (*http.Response, error) {
	bodyReader, err := getBodyReader(requestBody)
	if err != nil {
		return nil, err
	}
	endpoint = strings.TrimPrefix(endpoint, "/")
	request, err := http.NewRequestWithContext(ctx, method, i.baseURL+endpoint, bodyReader)
	if err != nil {
		return nil, err
	}
	for key, value := range i.header {
		request.Header.Set(key, value)
	}
	return i.httpClient.Do(request)
}

func (i *implClient) decode(response *http.Response, dest interface{}) *ResponseFail {
	defer response.Body.Close()
	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return &ResponseFail{
			Err:        err,
			StatusCode: response.StatusCode,
			Header:     response.Header,
		}
	}
	decoder := json.NewDecoder(response.Body)
	if response.StatusCode < 200 || response.StatusCode > 299 {
		decodeData := make(map[string]interface{})
		if err := decoder.Decode(&decodeData); err != nil {
			return &ResponseFail{
				Err:        err,
				StatusCode: response.StatusCode,
				Header:     response.Header,
				Body:       bodyBytes,
			}
		}
		if err != nil {
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
		if err := decoder.Decode(dest); err != nil {
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

func (i *implClient) Get(ctx context.Context, endpoint string, requestBody interface{}, dest interface{}) *ResponseFail {
	response, err := i.do(ctx, http.MethodGet, endpoint, requestBody)
	if err != nil {
		return &ResponseFail{
			Err:        err,
			StatusCode: response.StatusCode,
			Header:     response.Header,
		}
	}
	return i.decode(response, dest)
}

func (i *implClient) Post(ctx context.Context, endpoint string, requestBody interface{}, dest interface{}) *ResponseFail {
	response, err := i.do(ctx, http.MethodPost, endpoint, requestBody)
	if err != nil {
		return &ResponseFail{
			Err:        err,
			StatusCode: response.StatusCode,
			Header:     response.Header,
		}
	}
	return i.decode(response, dest)
}

func (i *implClient) Patch(ctx context.Context, endpoint string, requestBody interface{}, dest interface{}) *ResponseFail {
	response, err := i.do(ctx, http.MethodPatch, endpoint, requestBody)
	if err != nil {
		return &ResponseFail{
			Err:        err,
			StatusCode: response.StatusCode,
			Header:     response.Header,
		}
	}
	return i.decode(response, dest)
}

func (i *implClient) Put(ctx context.Context, endpoint string, requestBody interface{}, dest interface{}) *ResponseFail {
	response, err := i.do(ctx, http.MethodPut, endpoint, requestBody)
	if err != nil {
		return &ResponseFail{
			Err:        err,
			StatusCode: response.StatusCode,
			Header:     response.Header,
		}
	}
	return i.decode(response, dest)
}

func (i *implClient) Delete(ctx context.Context, endpoint string, requestBody interface{}, dest interface{}) *ResponseFail {
	response, err := i.do(ctx, http.MethodDelete, endpoint, requestBody)
	if err != nil {
		return &ResponseFail{
			Err:        err,
			StatusCode: response.StatusCode,
			Header:     response.Header,
		}
	}
	return i.decode(response, dest)
}

func (i implClient) SetAuthBasicToHeader(user, password string) {
	auth := []byte(user + ":" + password)
	i.header["Authorization"] = "Basic " + base64.StdEncoding.EncodeToString(auth)
}

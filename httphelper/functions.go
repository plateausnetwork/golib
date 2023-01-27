package httphelper

import (
	"encoding/json"
	"io"
	"net/http"
)

func DecodeResponse(response *http.Response, dest interface{}) *ResponseFail {
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
			Err:        ErrInvalidStatusCode,
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

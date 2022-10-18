package httpclient

import (
	"encoding/base64"
	"net/http"
)

func SetAuthBasicToHeader(header *http.Header, user, password string) {
	auth := []byte(user + ":" + password)
	header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString(auth))
}

func SetAppJSONToHeader(header *http.Header) {
	header.Set("Content-Type", "application/json")
}

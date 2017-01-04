package utils

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func GetRequestBody(r *http.Request) []byte {
	// Read the content
	var bodyBytes []byte
	if r.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(r.Body)
	}
	// Restore the io.ReadCloser to its original state
	r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	return bodyBytes
}

func GetRequestBodyString(r *http.Request) string {
	return string(GetRequestBody(r))
}

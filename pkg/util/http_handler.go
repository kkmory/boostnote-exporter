package util

import "net/http"

type HTTPClient struct {
	*http.Client
}

func NewHTTPHandler() *HTTPClient {
	client := new(http.Client)

	return &HTTPClient{client}
}

package request

import (
	"net/http"
)

type httpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type clientWrapper struct {
	client httpClient
}

func (c *clientWrapper) Do(req *http.Request) (*http.Response, error) {
	return c.client.Do(req)
}

package request

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseRequest(t *testing.T) {
	var newRequest InputRequest
	newRequest.Filepath = "../test/testRequest.json"
	parsedRequest := parseRequest(newRequest)

	assert.Equal(t, "GET", parsedRequest.Method)
	assert.Equal(t, "http://www.example.com/", parsedRequest.Url)
	assert.Equal(t, "application/json", parsedRequest.Headers["Content-Type"])
}

// TODO: Test executeHttpRequest, need to mock http response

func TestParseResponse(t *testing.T) {
	response := &http.Response{
		StatusCode: 200,
		Header:     http.Header{"Content-Type": []string{"application/json"}},
		Body:       io.NopCloser(bytes.NewBufferString(`{"foo":"bar"}`)),
	}
	parsedResponse := parseResponse(response)

	assert.Equal(t, 200, parsedResponse.StatusCode)
	assert.Equal(t, "application/json", parsedResponse.ContentType)
}

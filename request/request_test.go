package request

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseRequestJSON(t *testing.T) {
	var newRequest InputRequest
	newRequest.Filepath = "../test/testRequest.json"
	parsedRequest := parseRequest(newRequest)

	assert.Equal(t, "GET", parsedRequest.Method)
	assert.Equal(t, "http://www.example.com/", parsedRequest.Url)
	assert.Equal(t, "application/json", parsedRequest.Headers["Content-Type"])
}

func TestParseRequestYAML(t *testing.T) {
	var newRequest InputRequest
	newRequest.Filepath = "../test/testRequest.yaml"
	parsedRequest := parseRequest(newRequest)

	assert.Equal(t, "GET", parsedRequest.Method)
	assert.Equal(t, "http://www.example.com/", parsedRequest.Url)
	assert.Equal(t, "application/json", parsedRequest.Headers["Content-Type"])
}

func TestExecuteHttpRequest(t *testing.T) {
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, `{"foo": bar}`)
	}))
	defer svr.Close()

	newRequest := &HttpRequest{
		Method: "GET",
		Url:    svr.URL,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}

	response := executeHttpRequest(newRequest)

	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, "application/json", response.Header.Get("Content-Type"))
	responseBody, _ := io.ReadAll(response.Body)
	assert.Equal(t, `{"foo": bar}`, string(responseBody))
}

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

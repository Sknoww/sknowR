package request

import (
	"encoding/json"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func captureOutput(f func(response *HttpResponse), response *HttpResponse, isStdout bool) string {
	var orig *os.File
	if isStdout {
		orig = os.Stdout
	} else {
		orig = os.Stderr
	}

	defer func(orig *os.File) {
		if isStdout {
			os.Stdout = orig
		} else {
			os.Stderr = orig
		}
	}(orig)

	r, w, _ := os.Pipe()
	if isStdout {
		os.Stdout = w
	} else {
		os.Stderr = w
	}
	f(response)
	w.Close()
	out, _ := io.ReadAll(r)

	return string(out)
}

func TestOutputResponseBodyToStdout(t *testing.T) {
	response := &HttpResponse{
		StatusCode: 200,
		Body:       []byte(`{"foo":"bar"}`),
	}

	output := captureOutput(OutputResponseBodyToStdout, response, true)

	expected, _ := response.Body.MarshalJSON()
	assert.Equal(t, string(expected)+"\n", output)
}

func TestOutputResponseHeadersToSterr(t *testing.T) {
	response := &HttpResponse{
		StatusCode: 200,
		Headers:    map[string]string{"Content-Type": "application/json"},
	}

	output := captureOutput(OutputResponseHeadersToSterr, response, false)

	assert.Equal(t, "Content-Type: application/json\n", output)
}

func TestOutputResponseToFile(t *testing.T) {
	response := &HttpResponse{
		StatusCode: 200,
		Headers:    map[string]string{"Content-Type": "application/json"},
		Body:       []byte(`{"foo":"bar"}`),
	}

	outputFilePath := "testOutput.txt"
	OutputResponseToFile(outputFilePath, response)

	f, err := os.Open(outputFilePath)
	if err != nil {
		t.Error(err)
	}
	defer f.Close()

	b, err := io.ReadAll(f)
	if err != nil {
		t.Error(err)
	}

	expected, _ := json.MarshalIndent(response, "", "  ")
	assert.Equal(t, string(expected), string(b))

	t.Cleanup(func() {
		os.Remove(outputFilePath)
	})
}

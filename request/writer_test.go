package request

import (
	"encoding/json"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func captureOutput(f func(response *HttpResponse), response *HttpResponse) (string, string) {
	origStdout := os.Stdout
	origStderr := os.Stderr

	defer func() {
		os.Stdout = origStdout
		os.Stderr = origStderr
	}()

	rStdout, wStdout, _ := os.Pipe()
	rStderr, wStderr, _ := os.Pipe()

	os.Stdout = wStdout
	os.Stderr = wStderr

	f(response)

	wStdout.Close()
	wStderr.Close()

	outStdout, _ := io.ReadAll(rStdout)
	outStderr, _ := io.ReadAll(rStderr)

	return string(outStdout), string(outStderr)
}

func TestOutputResponseToStd(t *testing.T) {
	response := &HttpResponse{
		StatusCode: 200,
		Headers:    map[string]string{"Content-Type": "application/json"},
		Body:       []byte(`{"foo":"bar"}`),
	}

	stdOut, stdErr := captureOutput(OutputResponseToStd, response)

	expected, _ := response.Body.MarshalJSON()
	assert.Equal(t, string(expected)+"\n", stdOut)
	assert.Equal(t, "Content-Type: application/json\n", stdErr)
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

package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

type InputRequest struct {
	Filepath       string
	OutputFilePath string
}

type HttpRequest struct {
	Method  string            `json:"method"`
	Url     string            `json:"url"`
	Headers map[string]string `json:"headers"`
	Body    string            `json:"body"`
}

type HttpResponse struct {
	StatusCode  int               `yaml:"status.code" json:"statusCode"`
	ContentType string            `json:"contentType"`
	Headers     map[string]string `json:"headers"`
	Body        json.RawMessage   `json:"body"`
}

var NewRequest InputRequest

func HandleNewRequest(cmd *cobra.Command, args []string) {
	if NewRequest.Filepath != "" {
		// Parse request file
		parsedRequest := parseRequest(cmd, args)

		// Execute http request
		response := executeHttpRequest(parsedRequest)

		// Format response
		formattedResponse := parseResponse(response)

		// Write response to file if the user provided a filepath
		if NewRequest.OutputFilePath != "" {
			OutputResponseToFile(formattedResponse)
		} else {
			// Write response to stdout and stderr (default)
			OutputResponseBodyToStdout(formattedResponse)
			OutputResponseHeadersToSterr(formattedResponse)
		}

	}
}

// parseRequest parses the json request file provided by the user
func parseRequest(cmd *cobra.Command, args []string) *HttpRequest {
	fmt.Println("Parsing request file...")
	f, err := os.Open(NewRequest.Filepath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	byteValue, _ := io.ReadAll(f)

	var request HttpRequest
	json.Unmarshal(byteValue, &request)

	return &request
}

// executeHttpRequest executes the http request provided by the user
func executeHttpRequest(newRequest *HttpRequest) *http.Response {
	fmt.Println("Making http request...")

	request, err := http.NewRequest(newRequest.Method, newRequest.Url, bytes.NewBuffer([]byte(newRequest.Body)))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for key, value := range newRequest.Headers {
		request.Header.Set(key, value)
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return response
}

// parseResponse converts the reponse to a HttpResponse struct
func parseResponse(response *http.Response) *HttpResponse {
	var formattedResponse HttpResponse
	formattedResponse.StatusCode = response.StatusCode
	formattedResponse.ContentType = response.Header.Get("Content-Type")
	formattedResponse.Headers = make(map[string]string)
	for k, v := range response.Header {
		formattedResponse.Headers[k] = v[0]
	}

	var err error
	formattedResponse.Body, err = io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return &formattedResponse
}

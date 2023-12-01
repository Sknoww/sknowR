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

// InputRequest is the struct that is used to store the user input
type InputRequest struct {
	Filepath       string
	OutputFilePath string
}

// HttpRequest is the struct that is used to format the request
type HttpRequest struct {
	Method  string            `json:"method"`
	Url     string            `json:"url"`
	Headers map[string]string `json:"headers"`
	Body    string            `json:"body"`
}

// HttpResponse is the struct that is used to format the response
type HttpResponse struct {
	StatusCode  int               `yaml:"status.code" json:"statusCode"`
	ContentType string            `json:"contentType"`
	Headers     map[string]string `json:"headers"`
	Body        json.RawMessage   `json:"body"`
}

// HandleNewRequest handles input request
func HandleNewRequest(cmd *cobra.Command, args []string) {
	var newRequest InputRequest
	newRequest.Filepath, _ = cmd.Flags().GetString("filepath")
	newRequest.OutputFilePath, _ = cmd.Flags().GetString("output")

	// Check if user provided a filepath
	if newRequest.Filepath != "" {
		// Parse request file
		parsedRequest := parseRequest(newRequest)

		// Execute http request
		response := executeHttpRequest(parsedRequest)

		// Format response
		formattedResponse := parseResponse(response)

		// Write response to file if the user provided a filepath
		if newRequest.OutputFilePath != "" {
			OutputResponseToFile(newRequest.OutputFilePath, formattedResponse)
		} else {
			// Write response to stdout and stderr (default)
			OutputResponseToStd(formattedResponse)
		}

	} else {
		// Exit if no filepath provided
		fmt.Println("No request filepath provided")
		os.Exit(1)
	}
}

// parseRequest parses the json request file provided by the user
func parseRequest(newRequest InputRequest) *HttpRequest {

	// Open request file
	f, err := os.Open(newRequest.Filepath)
	if err != nil {
		fmt.Println("Error opening request file")
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	// Read request file
	byteValue, _ := io.ReadAll(f)

	// Parse request file
	var request HttpRequest
	json.Unmarshal(byteValue, &request)

	return &request
}

// executeHttpRequest executes the http request provided by the user
func executeHttpRequest(newRequest *HttpRequest) *http.Response {

	// Create http request
	request, err := http.NewRequest(newRequest.Method, newRequest.Url, bytes.NewBuffer([]byte(newRequest.Body)))
	if err != nil {
		fmt.Println("Error creating http request")
		fmt.Println(err)
		os.Exit(1)
	}

	// Add headers to request
	for key, value := range newRequest.Headers {
		request.Header.Set(key, value)
	}

	// Execute request
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Error executing http request")
		fmt.Println(err)
		os.Exit(1)
	}

	return response
}

// parseResponse converts the reponse to a HttpResponse struct
func parseResponse(response *http.Response) *HttpResponse {

	// Read response body
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body")
		fmt.Println(err)
		os.Exit(1)
	}

	// Create HttpResponse struct
	fResponse := &HttpResponse{
		StatusCode:  response.StatusCode,
		ContentType: response.Header.Get("Content-Type"),
		Headers:     make(map[string]string),
		Body:        responseBody,
	}

	// Convert response headers to map
	for k, v := range response.Header {
		fResponse.Headers[k] = v[0]
	}

	return fResponse
}

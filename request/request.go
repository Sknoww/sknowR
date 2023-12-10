package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

// InputRequest is the struct that is used to store the user input
type InputRequest struct {
	Filepath       string
	OutputFilePath string
	IsDownload     bool
}

// HttpRequest is the struct that is used to format the request
// Different versions needed for yaml and json due to body formatting
type HttpRequest struct {
	Method  string                 `json:"method" yaml:"method"`
	Url     string                 `json:"url" yaml:"url"`
	Params  map[string]string      `json:"params" yaml:"params"`
	Headers map[string]string      `json:"headers" yaml:"headers"`
	Cookies map[string]string      `json:"cookies" yaml:"cookies"`
	Data    map[string]interface{} `json:"data" yaml:"data"`
}

// HttpResponse is the struct that is used to format the response
type HttpResponse struct {
	StatusCode  int               `json:"statusCode"`
	ContentType string            `json:"contentType"`
	Headers     map[string]string `json:"headers"`
	Body        json.RawMessage   `json:"body"`
}

// HandleNewRequest handles input request
func HandleNewRequest(cmd *cobra.Command, args []string) {
	// Create new request struct
	var newRequest InputRequest
	newRequest.Filepath, _ = cmd.Flags().GetString("filepath")
	newRequest.OutputFilePath, _ = cmd.Flags().GetString("output")
	newRequest.IsDownload, _ = cmd.Flags().GetBool("download")

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

	// Check if file is yaml
	ext := path.Ext(newRequest.Filepath)
	if ext == ".yaml" || ext == ".yml" {
		// Convert yaml to json
		err = yaml.Unmarshal(byteValue, &request)
		if err != nil {
			fmt.Println("Error parsing yaml file")
			fmt.Println(err)
			os.Exit(1)
		}
	} else {
		err = json.Unmarshal(byteValue, &request)
		if err != nil {
			fmt.Println("Error parsing json file")
			fmt.Println(err)
			os.Exit(1)
		}
	}

	return &request
}

// executeHttpRequest executes the http request provided by the user
func executeHttpRequest(newRequest *HttpRequest) *http.Response {
	// Add params to url if provided
	if len(newRequest.Params) > 0 {
		newRequest.Url += "?"
		for key, value := range newRequest.Params {
			newRequest.Url += key + "=" + value + "&"
		}
		newRequest.Url = newRequest.Url[:len(newRequest.Url)-1]
	}

	// Marshal data to bytes for http request
	var data []byte
	var err error
	if len(newRequest.Data) > 0 {
		data, err = json.Marshal(newRequest.Data)
		if err != nil {
			fmt.Println("Error marshalling data")
			fmt.Println(err)
			os.Exit(1)
		}
	}

	// Create http request
	request, err := http.NewRequest(newRequest.Method, newRequest.Url, bytes.NewBuffer(data))
	if err != nil {
		fmt.Println("Error creating http request")
		fmt.Println(err)
		os.Exit(1)
	}

	// Add headers to request
	for key, value := range newRequest.Headers {
		request.Header.Set(key, value)
	}

	// Add cookies to request
	for key, value := range newRequest.Cookies {
		request.AddCookie(&http.Cookie{Name: key, Value: value})
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
	// Close response body
	defer response.Body.Close()

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

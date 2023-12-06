package request

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// OutputResponseToStd writes the response body to stdout and stderr (default)
func OutputResponseToStd(response *HttpResponse) {
	// Write response status code to stderr
	fmt.Fprintf(os.Stderr, "Status: %d\n", response.StatusCode)

	// Write response headers to stderr in json format
	h := marshalResponse(response.Headers)
	fmt.Fprintf(os.Stderr, "%s\n", string(h))

	// Convert response body to bytes
	b := []byte(response.Body)

	// Check if content type is json
	if strings.Contains(response.Headers["Content-Type"], "application/json") {
		// Marshal response struct to bytes
		b = marshalResponse(response.Body)
	}

	// Write response body to stdout
	os.Stdout.WriteString(string(b))

}

// OutputResponseToFile writes the response to a file if the user provided a filepath
func OutputResponseToFile(outputFilePath string, response *HttpResponse) {
	// Create output file
	f, err := os.Create(outputFilePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	// Convert response body to bytes
	b := []byte(response.Body)

	// Check if content type is json
	if strings.Contains(response.Headers["Content-Type"], "application/json") {
		// Marshal response struct to bytes
		b = marshalResponse(response)
	}

	// Write response to file
	_, err = f.WriteString(string(b))
	if err != nil {
		fmt.Println("Error writing response to file")
		fmt.Println(err)
		os.Exit(1)
	}
}

// marshalResponse marshals any part of the response to json
func marshalResponse(response interface{}) []byte {
	// Marshal response struct to bytes for writing to file
	// Adds indentation to json
	b, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling response to json")
		fmt.Println(err)
		os.Exit(1)
	}
	return b
}

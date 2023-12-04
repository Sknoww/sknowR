package request

import (
	"encoding/json"
	"fmt"
	"os"
)

// OutputResponseToStd writes the response body to stdout and stderr (default)
func OutputResponseToStd(response *HttpResponse) {

	// If content type is not json write response body to stdout
	if response.Headers["Content-Type"] == "application/pdf" {
		// Write pdf to file
		f, err := os.Create("output.pdf")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer f.Close()

		// Write response body to file
		_, err = f.Write(response.Body)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		return
	}

	// If content type is json, write response body to stdout in json format
	// Write response body to stdout in json format
	b := marshalResponse(response.Body)
	fmt.Printf("%s\n", string(b))

	// Write response headers to stderr in json format
	h := marshalResponse(response.Headers)
	fmt.Fprintf(os.Stderr, "%s\n", string(h))
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

	// Marshal response struct to bytes
	b := marshalResponse(response)

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

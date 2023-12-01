package request

import (
	"encoding/json"
	"fmt"
	"os"
)

// OutputResponseToStd writes the response body to stdout and stderr (default)
func OutputResponseToStd(response *HttpResponse) {
	// Write response body to stdout
	fmt.Printf("%s\n", response.Body)

	// Write response headers to stderr
	for k, v := range response.Headers {
		fmt.Fprintf(os.Stderr, "%s: %s\n", k, v)
	}
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

	// Marshal response struct to bytes for writing to file
	// Adds indentation to json
	b, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling response to json")
		fmt.Println(err)
		os.Exit(1)
	}

	// Write response to file
	_, err = f.WriteString(string(b))
	if err != nil {
		fmt.Println("Error writing response to file")
		fmt.Println(err)
		os.Exit(1)
	}
}

package request

import (
	"encoding/json"
	"fmt"
	"os"
)

// OutputResponseBodyToStdout writes the response body to stdout (default)
func OutputResponseBodyToStdout(response *HttpResponse) {
	fmt.Printf("%s\n", response.Body)
}

// OutputResponseHeadersToSterr writes the response headers to stderr (default)
func OutputResponseHeadersToSterr(response *HttpResponse) {
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
		fmt.Println("1")
		fmt.Println(err)
		os.Exit(1)
	}

	// Write response to file
	_, err = f.WriteString(string(b))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

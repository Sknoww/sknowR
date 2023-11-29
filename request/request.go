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

var NewRequest InputRequest

func ParseRequest(cmd *cobra.Command, args []string) {
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

	makeHttpRequest(&request)

}

func makeHttpRequest(newRequest *HttpRequest) {
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

	responseBody, err := io.ReadAll(response.Body)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	formattedResponseBody := formatJSON(responseBody)

	fmt.Println(response.Status)
	fmt.Println(formattedResponseBody)

}

// Helpers //

// formatJSON formats the json response body
func formatJSON(data []byte) string {
	var out bytes.Buffer
	err := json.Indent(&out, data, "", " ")

	if err != nil {
		fmt.Println(err)
	}

	d := out.Bytes()
	return string(d)
}

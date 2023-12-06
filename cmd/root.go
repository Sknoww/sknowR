/*
Copyright © 2023 sknoww sknow.codes@gmail.com
*/
package cmd

import (
	"os"

	"github.com/sknoww/sknowR/request"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sknowR",
	Short: "A CLI tool for making HTTP requests written in Go",
	Long: `A CLI tool for making HTTP requests written in Go. 
	Requests can be passed in using .json files. Response body 
	is written to stdout and response headers and status are 
	written to stderr by default. You can add the -o flag to 
	specify and output file.`,
	Run: request.HandleNewRequest,
}

func init() {
	// Add filepath flag
	rootCmd.PersistentFlags().StringP("filepath", "f", "", "The path to the request file")
	rootCmd.MarkPersistentFlagRequired("filepath")

	// Add output flag
	rootCmd.Flags().StringP("output", "o", "", "The path to the output file (optional if not downloading a file)")
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

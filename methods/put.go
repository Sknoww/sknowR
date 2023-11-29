/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package methods

import (
	"fmt"

	"github.com/sknoww/sknowR/cmd"
	"github.com/spf13/cobra"
)

// putCmd represents the put command
var PutCmd = &cobra.Command{
	Use:   "put",
	Short: "Use to send a PUT request to a server.",
	Long:  ``,
	Run:   parsePutRequest,
}

func init() {
	cmd.RootCmd.AddCommand(PutCmd)
}

func parsePutRequest(cmd *cobra.Command, args []string) {
	fmt.Println("put called")
}

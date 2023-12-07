package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = NewRootCmd()

func NewRootCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "tree",
		Short: "List directory and files in a tree-like view",
		Long:  "A command-utility which lists directories and files recursively in a tree-like view for the user.",
		Run: func(cmd *cobra.Command, args []string) {
			dir, _ := os.ReadDir("./")
			for _, x := range dir {
				fmt.Fprintf(cmd.OutOrStdout(), x.Name())
			}
		},
	}
}

func init() {

}

func Execute() {
	rootCmd.Execute()
}

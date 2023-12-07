package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tree",
	Short: "List directory and files in a tree-like view",
	Long:  "A command-utility which lists directories and files recursively in a tree-like view for the user.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("You have successfully set up the base!")
	},
}

func init() {

}

func Execute() {
	rootCmd.Execute()
}

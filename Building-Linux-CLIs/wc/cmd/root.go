package cmd

import (
	"fmt"
	"syscall"

	"github.com/spf13/cobra"
)

var l bool

var rootCmd = NewRootCmd()

func NewRootCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "wc",
		Short: "count lines, words and bytes from input",
		Long:  `wc is a command-line utility for counting lines words and bytes from the input.`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			filename := args[0]
			flags := syscall.O_RDONLY
			mode := uint32(0666) /* Does not matter  */
			_, err := syscall.Open(filename, flags, mode)
			if err != nil {
				fmt.Fprintf(cmd.OutOrStdout(), "%s: %s: open: %s", cmd.Name(), filename, err.Error())
				return
			}
		},
	}
}

func init() {
	rootCmd.Flags().BoolVarP(&l, "", "l", false, "The number of lines in each input file is written to the standard output")
}

func Execute() {
	rootCmd.Execute()
}

/* Code snippets that may be useful */

// RunE: func(cmd *cobra.Command, args []string) error {
// 	filename := args[0]
// 	_, err := os.Open(filename)
// 	if err != nil {
// 		fmt.Fprintf(cmd.OutOrStdout(), "%v", err)
// 		return err
// 	}

// 	return nil
// },

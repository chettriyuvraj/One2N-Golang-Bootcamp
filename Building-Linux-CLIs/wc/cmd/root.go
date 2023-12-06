package cmd

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

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
			f, err := os.Open(filename)
			if err != nil {
				fmt.Fprintf(cmd.OutOrStdout(), "%s: %s: open: %s", cmd.Name(), filename, GetBaseError(err.Error()))
				return
			}

			r := bufio.NewReader(f)
			lineCount, wordCount, charCount := 0, 0, 0
			for {
				s, err := r.ReadString('\n')
				if err != nil {
					if err != io.EOF {
						fmt.Fprintf(cmd.OutOrStdout(), "%s: %s: read: %s", cmd.Name(), filename, GetBaseError(err.Error()))
						return
					}
					break
				}

				charCount += len(s)
				lineCount += 1
				wordCount += len(strings.Fields(s))
			}

			if l {
				fmt.Fprintf(cmd.OutOrStdout(), "%8d", lineCount)
			}

			fmt.Fprintf(cmd.OutOrStdout(), " %s", filename)

		},
	}
}

func init() {
	rootCmd.Flags().BoolVarP(&l, "", "l", false, "The number of lines in each input file is written to the standard output")
}

func Execute() {
	rootCmd.Execute()
}

/***** Helpers *****/

func GetBaseError(s string) string {
	errStrs := strings.Split(s, ":")
	return strings.TrimSpace(errStrs[len(errStrs)-1])
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

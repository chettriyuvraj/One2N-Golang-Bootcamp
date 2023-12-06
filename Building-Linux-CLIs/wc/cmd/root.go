package cmd

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var l, w bool

var rootCmd = NewRootCmd()

func NewRootCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "wc",
		Short: "count lines, words and bytes from input",
		Long:  `wc is a command-line utility for counting lines words and bytes from the input.`,
		Run: func(cmd *cobra.Command, args []string) {
			f, fname := os.Stdin, ""
			if len(args) > 0 {
				fname = args[0]
				file, err := os.Open(fname)
				if err != nil {
					fmt.Fprintf(cmd.OutOrStdout(), "%s: %s: open: %s", cmd.Name(), fname, GetBaseError(err.Error()))
					return
				}

				f = file
			}

			r := bufio.NewReader(f)
			lineCount, wordCount, charCount := 0, 0, 0
			for {
				s, err := r.ReadString('\n')
				if err != nil {
					if err != io.EOF {
						fmt.Fprintf(cmd.OutOrStdout(), "%s: %s: read: %s", cmd.Name(), fname, GetBaseError(err.Error()))
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

			if w {
				fmt.Fprintf(cmd.OutOrStdout(), "%8d", wordCount)
			}

			if fname != "" {
				fmt.Fprintf(cmd.OutOrStdout(), " %s", fname)
			}

		},
	}
}

func init() {
	SetFlags(rootCmd)
}

func Execute() {
	rootCmd.Execute()
}

/***** Helpers *****/

func GetBaseError(s string) string {
	errStrs := strings.Split(s, ":")
	return strings.TrimSpace(errStrs[len(errStrs)-1])
}

func SetFlags(cmd *cobra.Command) {
	cmd.Flags().BoolVarP(&l, "line", "l", false, "The number of lines in each input file is written to the standard output")
	cmd.Flags().BoolVarP(&w, "word", "w", false, "The number of lines in each input file is written to the standard output")
}

/* Code snippets that may be useful */

// RunE: func(cmd *cobra.Command, args []string) error {
// 	fname := args[0]
// 	_, err := os.Open(fname)
// 	if err != nil {
// 		fmt.Fprintf(cmd.OutOrStdout(), "%v", err)
// 		return err
// 	}

// 	return nil
// },

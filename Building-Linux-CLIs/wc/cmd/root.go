package cmd

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var l, w, c bool

var rootCmd = NewRootCmd()

func NewRootCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "wc",
		Short: "count lines, words and bytes from input",
		Long:  `wc is a command-line utility for counting lines words and bytes from the input.`,
		Run: func(cmd *cobra.Command, args []string) {

			noFlagSet := !l && !w && !c
			if noFlagSet {
				SetAllFlags()
			}

			if len(args) < 1 {
				wc(cmd, "")
				return
			}

			totallc, totalwc, totalcc := 0, 0, 0
			for i, fname := range args {
				if i > 0 {
					fmt.Fprintln(cmd.OutOrStdout())
				}
				curlc, curwc, curcc := wc(cmd, fname)
				totallc += curlc
				totalwc += curwc
				totalcc += curcc
			}

			if len(args) > 1 {
				fmt.Fprintln(cmd.OutOrStdout())
				if l {
					fmt.Fprintf(cmd.OutOrStdout(), "%8d", totallc)
				}

				if w {
					fmt.Fprintf(cmd.OutOrStdout(), "%8d", totalwc)
				}

				if c {
					fmt.Fprintf(cmd.OutOrStdout(), "%8d", totalcc)
				}
				fmt.Fprintf(cmd.OutOrStdout(), " total\n")
			}
		},
	}
}

func init() {
	AddFlags(rootCmd)
}

func Execute() {
	rootCmd.Execute()
}

func wc(cmd *cobra.Command, fname string) (lineCount, wordCount, charCount int) {
	f := os.Stdin

	if fname != "" {
		file, err := os.Open(fname)
		if err != nil {
			fmt.Fprintf(cmd.OutOrStdout(), "%s: %s: open: %s", cmd.Name(), fname, GetBaseError(err.Error()))
			return
		}
		defer file.Close()

		f = file
	}

	r := bufio.NewReader(f)
	lineCount, wordCount, charCount = 0, 0, 0
	for {
		s, err := r.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				fmt.Fprintf(cmd.OutOrStdout(), "%s: %s: read: %s", cmd.Name(), fname, GetBaseError(err.Error()))
				return lineCount, wordCount, charCount
			}
			charCount += len(s)
			wordCount += len(strings.Fields(s))
			break
		}

		charCount += len(s)
		lineCount += 1
		wordCount += len(strings.Fields(s))
	}

	if l {
		fmt.Fprintf(cmd.OutOrStdout(), "%8d", lineCount)
		// if w {
		// 	fmt.Printf(" ")
		// }
	}

	if w {
		fmt.Fprintf(cmd.OutOrStdout(), "%8d", wordCount)
		// if c {
		// 	fmt.Printf(" ")
		// }
	}

	if c {
		fmt.Fprintf(cmd.OutOrStdout(), "%8d", charCount)
	}

	if fname != "" {
		fmt.Fprintf(cmd.OutOrStdout(), " %s", fname)
	}

	return lineCount, wordCount, charCount
}

/***** Helpers *****/

func GetBaseError(s string) string {
	errStrs := strings.Split(s, ":")
	return strings.TrimSpace(errStrs[len(errStrs)-1])
}

func AddFlags(cmd *cobra.Command) {
	cmd.Flags().BoolVarP(&l, "line", "l", false, "The number of lines in each input file is written to the standard output")
	cmd.Flags().BoolVarP(&w, "word", "w", false, "The number of lines in each input file is written to the standard output")
	cmd.Flags().BoolVarP(&c, "char", "c", false, "The number of bytes in each input file is written to the standard output")
}

func SetAllFlags() {
	l, w, c = true, true, true
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

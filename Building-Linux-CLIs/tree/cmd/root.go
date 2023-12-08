package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const (
	defaultbasepath = "."
	defaulterrmsg   = "error opening dir"
)

var rootCmd = NewRootCmd()

func NewRootCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "tree",
		Short: "List directory and files in a tree-like view",
		Long:  "A command-utility which lists directories and files recursively in a tree-like view for the user.",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				args = append(args, defaultbasepath)
			}

			fcount, dcount := 0, 1
			for _, path := range args {
				fmt.Fprintf(cmd.OutOrStdout(), "%s", path)
				fc, dc, err := tree(cmd, path, []bool{})
				if err != nil {
					fmt.Fprintf(cmd.OutOrStdout(), "  [error opening dir]")
					fcount, dcount = 0, 0
					break
				}
				fcount += fc
				dcount += dc
			}

			/* Edge case - empty directory */
			if fcount == 0 && dcount == 1 {
				dcount -= 1
			}

			dirStr := "directories"
			if dcount == 1 {
				dirStr = "directory"
			}
			fileStr := "files"
			if fcount == 1 {
				fileStr = "file"
			}
			fmt.Fprintf(cmd.OutOrStdout(), "\n\n%d %s, %d %s\n", dcount, dirStr, fcount, fileStr)
		},
	}
}

func init() {

}

func Execute() {
	rootCmd.Execute()
}

func tree(cmd *cobra.Command, path string, isAncestorLastDir []bool) (fcount int, dcount int, err error) {
	direntries, err := os.ReadDir(path)
	if err != nil {
		return -1, -1, err
	}

	for i, d := range direntries {
		isLastElem := i == len(direntries)-1
		printDirEntry(cmd, d, isLastElem, isAncestorLastDir)

		if !d.IsDir() {
			fcount += 1
			continue
		}

		fc, dc, err := tree(cmd, fmt.Sprintf("%s/%s", path, d.Name()), append(isAncestorLastDir, isLastElem))
		if err != nil {
			return -1, -1, err
		}
		fcount += fc
		dcount += dc + 1
	}

	return fcount, dcount, nil
}

/***** Helpers *****/

func printDirEntry(cmd *cobra.Command, d os.DirEntry, isLastElem bool, isAncestorLastDir []bool) {
	fmt.Fprintf(cmd.OutOrStdout(), "\n")

	for _, b := range isAncestorLastDir {
		if b {
			fmt.Fprintf(cmd.OutOrStdout(), " ")
		} else {
			fmt.Fprintf(cmd.OutOrStdout(), "%s", vline)
		}
		fmt.Fprintf(cmd.OutOrStdout(), "   ")
	}

	if isLastElem {
		fmt.Fprintf(cmd.OutOrStdout(), "%s%s%s %s", vshafthalf, hline, hline, d.Name())
		return
	}

	fmt.Fprintf(cmd.OutOrStdout(), "%s%s%s %s", vshaft, hline, hline, d.Name())
}

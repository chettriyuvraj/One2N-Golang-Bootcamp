package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

const (
	defaultbasepath = "."
	defaulterrmsg   = "error opening dir"
)

var rootCmd = NewRootCmd()

var ff, df bool

func NewRootCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "tree",
		Short: "List directory and files in a tree-like view",
		Long:  "A command-utility which lists directories and files recursively in a tree-like view for the user.",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				args = append(args, defaultbasepath)
			}

			tree(cmd, args)
		},
	}
}

func init() {
	setFlags(rootCmd)
}

func Execute() {
	rootCmd.Execute()
}

func setFlags(cmd *cobra.Command) {
	cmd.Flags().BoolVarP(&ff, "fullpath", "f", false, "Print the full path prefix for each file")
	cmd.Flags().BoolVarP(&df, "dir", "d", false, "Print only directories")
}

func tree(cmd *cobra.Command, args []string) {
	ancestors := []DirInfo{}
	if ff {
		ancestors = append(ancestors, DirInfo{isDummyEntry: true, DummyName: strings.Join(args, "")})
	}

	fcount, dcount := 0, 1
	for _, path := range args {

		if ff { /* Edge case - f flag removes trailing '/' */
			n := len(path)
			if path[n-1] == '/' {
				path = path[:n-1]
			}
		}

		fmt.Fprintf(cmd.OutOrStdout(), "%s", path)
		fc, dc, err := treedfs(cmd, path, ancestors)
		if err != nil {
			fmt.Fprintf(cmd.OutOrStdout(), "  [error opening dir]")
			fcount, dcount = 0, 0
			break
		}
		fcount += fc
		dcount += dc
	}

	printfiledircount(cmd, fcount, dcount)
}

func treedfs(cmd *cobra.Command, path string, dirAncestors []DirInfo) (fcount int, dcount int, err error) {
	direntries, err := os.ReadDir(path)
	if err != nil {
		return -1, -1, err
	}

	if df {
		direntries = filterDirs(direntries)
	}

	for i, d := range direntries {

		isLastElem := i == len(direntries)-1
		dinfo := DirInfo{dir: d, isLastElem: isLastElem}
		printDirEntry(cmd, d, isLastElem, dirAncestors)

		if !d.IsDir() {
			fcount += 1
			continue
		}

		fc, dc, err := treedfs(cmd, fmt.Sprintf("%s/%s", path, d.Name()), append(dirAncestors, dinfo))
		if err != nil {
			return -1, -1, err
		}
		fcount += fc
		dcount += dc + 1
	}

	return fcount, dcount, nil
}

/***** Helpers *****/

func printDirEntry(cmd *cobra.Command, d os.DirEntry, isLastElem bool, dirAncestors []DirInfo) {
	fmt.Fprintf(cmd.OutOrStdout(), "\n")

	var dName strings.Builder

	for _, di := range dirAncestors {
		if !di.isDummyEntry {
			if di.isLastElem {
				fmt.Fprintf(cmd.OutOrStdout(), " ")
			} else {
				fmt.Fprintf(cmd.OutOrStdout(), "%s", vline)
			}
			fmt.Fprintf(cmd.OutOrStdout(), "   ")

			if ff {
				dName.WriteString(di.dir.Name() + "/")
			}
			continue
		}

		if ff {
			dName.WriteString(di.DummyName)
			if di.DummyName[len(di.DummyName)-1] != '/' {
				dName.WriteString("/")
			}
		}

	}

	dName.WriteString(d.Name())

	if isLastElem {
		fmt.Fprintf(cmd.OutOrStdout(), "%s%s%s %s", vshafthalf, hline, hline, dName.String())
		return
	}

	fmt.Fprintf(cmd.OutOrStdout(), "%s%s%s %s", vshaft, hline, hline, dName.String())
}

func printfiledircount(cmd *cobra.Command, fcount, dcount int) {
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
	fmt.Fprintf(cmd.OutOrStdout(), "\n\n%d %s", dcount, dirStr)
	if !df {
		fmt.Fprintf(cmd.OutOrStdout(), ", %d %s", fcount, fileStr)
	}
	fmt.Fprintf(cmd.OutOrStdout(), "\n")
}

func filterDirs(de []os.DirEntry) []os.DirEntry {
	dirs := []os.DirEntry{}
	for _, elem := range de {
		if elem.IsDir() {
			dirs = append(dirs, elem)
		}
	}
	return dirs
}

func (d DirInfo) String() string {
	if !d.isDummyEntry {
		return fmt.Sprintf("%s ", d.dir.Name())
	}
	return fmt.Sprintf("%s ", d.DummyName)
}

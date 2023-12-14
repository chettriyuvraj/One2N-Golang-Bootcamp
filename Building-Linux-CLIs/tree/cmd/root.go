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

var ff, df, jf bool

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
	cmd.Flags().BoolVarP(&jf, "json", "j", false, "Print tree as JSON")
}

func tree(cmd *cobra.Command, args []string) {

	/*** DFS ***/
	fcount, dcount := 0, 0
	for i, path := range args {

		/** Pre-processing **/
		ancestors := []TreeElem{}

		if ff {
			ancestors = append(ancestors, TreeElem{isPrefixElem: true, Name: strings.Join(args, "")})
			n := len(path)
			/* Edge case - remove trailing '/' */
			if path[n-1] == '/' {
				path = path[:n-1]
			}
		}

		fmt.Fprintf(cmd.OutOrStdout(), "%s", path)
		te, err := treedfs(cmd, path, ancestors)
		if err != nil {
			fmt.Fprintf(cmd.OutOrStdout(), "  [error opening dir]")
			fcount, dcount = 0, 0
			break
		}

		fcount += te.fcount
		dcount += te.dcount + 1

		if i != len(args)-1 {
			fmt.Fprintf(cmd.OutOrStdout(), "\n")
		}
	}

	printfiledircount(cmd, fcount, dcount)
}

func treedfs(cmd *cobra.Command, path string, dirAncestors []TreeElem) (TreeElem, error) {
	parent := TreeElem{Name: path}

	direntries, err := os.ReadDir(path)
	if err != nil {
		return TreeElem{}, err
	}

	if df {
		direntries = filterDirs(direntries)
	}

	for i, d := range direntries {

		isLastElem := i == len(direntries)-1
		child := TreeElem{Name: d.Name(), isLastElem: isLastElem}
		printDirEntry(cmd, child, dirAncestors)

		if !d.IsDir() {
			parent.fcount += 1
			continue
		}

		grandchild, err := treedfs(cmd, fmt.Sprintf("%s/%s", path, d.Name()), append(dirAncestors, child))
		if err != nil {
			return grandchild, err
		}

		parent.fcount += grandchild.fcount
		parent.dcount += grandchild.dcount + 1
	}

	return parent, nil
}

/***** Helpers *****/

/**** Print Helpers ****/

func printDirEntry(cmd *cobra.Command, te TreeElem, dirAncestors []TreeElem) {
	var sb strings.Builder

	fmt.Fprintf(cmd.OutOrStdout(), "\n")

	for _, ancestor := range dirAncestors {

		if ancestor.isPrefixElem && ff {
			sb.WriteString(ancestor.Name)
			if ancestor.Name[len(ancestor.Name)-1] != '/' {
				sb.WriteString("/")
			}
			continue
		}

		if ancestor.isLastElem {
			fmt.Fprintf(cmd.OutOrStdout(), " ")
		} else {
			fmt.Fprintf(cmd.OutOrStdout(), "%s", vline)
		}

		fmt.Fprintf(cmd.OutOrStdout(), "   ")
		if ff {
			sb.WriteString(ancestor.Name + "/")
		}

	}

	sb.WriteString(te.Name)

	if te.isLastElem {
		fmt.Fprintf(cmd.OutOrStdout(), "%s%s%s %s", vshafthalf, hline, hline, sb.String())
		return
	}

	fmt.Fprintf(cmd.OutOrStdout(), "%s%s%s %s", vshaft, hline, hline, sb.String())
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

func printjson(cmd *cobra.Command, elemType string, name string) {

}

/**** Misc Helpers ****/

func filterDirs(de []os.DirEntry) []os.DirEntry {
	dirs := []os.DirEntry{}
	for _, elem := range de {
		if elem.IsDir() {
			dirs = append(dirs, elem)
		}
	}
	return dirs
}

func (d TreeElem) String() string {
	if !d.isPrefixElem {
		return fmt.Sprintf("%s ", d.Name)
	}
	return fmt.Sprintf("%s ", d.Name)
}

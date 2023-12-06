package cmd

import (
	"os"
	"strings"
	"testing"
)

func TestLFlag(t *testing.T) {
	tc := []struct {
		name string
		args []string
		want string
	}{
		{
			name: "Error when file does not exist",
			args: []string{"-l", "testfile"},
			want: "wc: testfile: open: no such file or directory",
		},
		{
			name: "Error when file is a directory",
			args: []string{"-l", "testdata"},
			want: "wc: testdata: read: is a directory",
		},
		{
			name: "Successful count of lines in file",
			args: []string{"-l", "testdata/test.txt"},
			want: "       3 testdata/test.txt",
		},
		{
			name: "Successful count of lines in stdin",
			args: []string{"-l"},
			want: "       3",
		},
	}

	cmd := rootCmd
	for _, test := range tc {
		var b strings.Builder

		/* Special case: mocking stdin as a file if no args provided */
		if len(test.args) < 2 {
			f, _ := os.Open("testdata/test.txt")
			os.Stdin = f
		}

		cmd.SetArgs(test.args)
		cmd.SetOut(&b)
		cmd.Execute()

		got := b.String()
		if got != test.want {
			t.Errorf("got %q, wanted %q", got, test.want)
		}

	}
}

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
		want map[string]string
		err  string
	}{
		{
			name: "Error when file does not exist",
			args: []string{"testfile"},
			want: map[string]string{},
			err:  "wc: testfile: open: no such file or directory",
		},
		{
			name: "Error when file is a directory",
			args: []string{"testdata"},
			want: map[string]string{},
			err:  "wc: testdata: read: is a directory",
		},
		{
			name: "Successful count of lines in file",
			args: []string{"testdata/test.txt"},
			want: map[string]string{"-l": "       3 testdata/test.txt"},
			err:  "",
		},
		{
			name: "Successful count of lines in stdin",
			args: []string{},
			want: map[string]string{"-l": "       3"},
			err:  "",
		},
	}

	flags := []string{"-l"} //, "-w"} //, "-m", "-lwm", "-mwl"}
	cmd := rootCmd

	for _, test := range tc {
		for _, flag := range flags {
			test.args = append([]string{flag}, test.args...)
			var b strings.Builder

			/* Special case: mocking stdin as a file if no args provided */
			if len(test.args) < 2 {
				stdin := os.Stdin
				f, _ := os.Open("testdata/test.txt")
				os.Stdin = f
				defer func() {
					os.Stdin = stdin
				}()
			}

			cmd.SetArgs(test.args)
			cmd.SetOut(&b)
			cmd.Execute()

			got := b.String()
			if test.err != "" {
				if got != test.err {
					t.Errorf("error testing: got %q, wanted %q", got, test.want)
				}
				return
			}
			if got != test.want[flag] {
				t.Errorf("got %q, wanted %q", got, test.want)
			}
		}

	}
}

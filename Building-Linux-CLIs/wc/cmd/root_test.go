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
			name: "Edge case with no newlines",
			args: []string{"testdata/testnonewline.txt"},
			want: map[string]string{
				"-l": "       0 testdata/testnonewline.txt",
				"-w": "       1 testdata/testnonewline.txt",
			},
			err: "",
		},
		{
			name: "Edge case with emoji",
			args: []string{"testdata/testemoji.txt"},
			want: map[string]string{
				"-l": "       0 testdata/testemoji.txt",
				"-w": "       2 testdata/testemoji.txt",
			},
			err: "",
		},
		{
			name: "Edge case using stdin instead of file",
			args: []string{},
			want: map[string]string{
				"-l": "       3",
				"-w": "       6",
			},
			err: "",
		},
		{
			name: "Happy case",
			args: []string{"testdata/test.txt"},
			want: map[string]string{
				"-l": "       3 testdata/test.txt",
				"-w": "       6 testdata/test.txt",
			},
			err: "",
		},
	}

	flags := []string{"-l", "-w"} //, "-m", "-lwm", "-mwl"}

	for _, test := range tc {
		initargs := test.args

		for _, flag := range flags {
			test.args = append([]string{flag}, initargs...)
			var b strings.Builder = strings.Builder{}

			/* Special case: mocking stdin as a file if no args provided */
			isStdinMocked := false
			stdininit := os.Stdin
			if len(test.args) < 2 {
				isStdinMocked = true
				f, _ := os.Open("testdata/test.txt")
				os.Stdin = f
			}

			cmd := NewRootCmd()
			SetFlags(cmd)
			cmd.SetArgs(test.args)
			cmd.SetOut(&b)
			cmd.Execute()

			got := b.String()
			if test.err != "" {
				if got != test.err {
					t.Errorf("\n\nName: %q:\nArgs: %q\nGot: %q\nWanted: %q\n\n", test.name, test.args, got, test.want)
				}
				continue
			}
			if got != test.want[flag] {
				t.Errorf("\n\nName: %q:\nArgs: %q\nGot: %q\nWanted: %q\n\n", test.name, test.args, got, test.want[flag])
			}

			if isStdinMocked {
				os.Stdin = stdininit
			}
		}

	}
}

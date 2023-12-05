package cmd

import (
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
		// {
		// 	name: "Error when file is a directory",
		// 	args: []string{"test"},
		// 	want: "error",
		// },
	}

	cmd := rootCmd
	for _, test := range tc {
		var b strings.Builder

		cmd.SetArgs(test.args)
		cmd.SetOut(&b)
		cmd.Execute()

		got := b.String()
		if got != test.want {
			t.Errorf("got %q, wanted %q", got, test.want)
		}

	}
}

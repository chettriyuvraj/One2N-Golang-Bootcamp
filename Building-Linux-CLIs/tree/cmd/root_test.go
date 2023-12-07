package cmd

import (
	"strings"
	"testing"
)

func TestTreeBasic(t *testing.T) {
	tc := []struct {
		name string
		args []string
		want string
	}{
		{
			name: "Happy case",
			args: []string{},
			want: `.
			├── cmd
			│   ├── root.go
			│   └── root_test.go
			├── go.mod
			├── go.sum
			├── main.go
			└── testdir
				├── testdirinner
				├── testfile1.txt
				└── testfile2.txt
			
			4 directories, 7 files`,
		},
	}

	for _, test := range tc {
		var b strings.Builder
		rootCmd = NewRootCmd()

		rootCmd.SetOut(&b)
		rootCmd.SetArgs(test.args)
		rootCmd.Execute()

		got := b.String()
		if got != test.want {
			t.Errorf("\n\nTest: %s\nArgs: %s\nWant: %s\nGot: %s", test.name, test.args, test.want, got)
		}
	}
}

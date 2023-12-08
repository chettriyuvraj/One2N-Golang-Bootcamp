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
			name: "Happy case current directory",
			args: []string{"."},
			want: `.
├── chars.go
├── root.go
├── root_test.go
└── testdata
    ├── emptydir
    └── testdir
        ├── testdirinner
        │   └── testdirinner.txt
        ├── testfile1.txt
        └── testfile2.txt

5 directories, 6 files
`,
		},
		{
			name: "Happy case parent directory",
			args: []string{"../"},
			want: `../
├── cmd
│   ├── chars.go
│   ├── root.go
│   ├── root_test.go
│   └── testdata
│       ├── emptydir
│       └── testdir
│           ├── testdirinner
│           │   └── testdirinner.txt
│           ├── testfile1.txt
│           └── testfile2.txt
├── go.mod
├── go.sum
├── main.go
└── notes.txt

6 directories, 10 files
`,
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
			t.Errorf("\nTest: %s\nArgs: %s\nWant: %s\nGot: %s\n\n", test.name, test.args, test.want, got)
		}
	}
}

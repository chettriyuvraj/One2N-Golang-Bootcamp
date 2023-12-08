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
			name: "Happy case: multiple directories + files",
			args: []string{"testdata"},
			want: `testdata
├── emptydir
├── testdir
│   ├── testdirinner
│   │   └── testdirinner.txt
│   ├── testfile1.txt
│   └── testfile2.txt
└── testdironlyfiles
    └── 1.txt

5 directories, 4 files
`,
		},
		{
			name: "Happy case: only files",
			args: []string{"./testdata/testdironlyfiles"},
			want: `./testdata/testdironlyfiles
└── 1.txt

1 directory, 1 file
`,
		},
		{
			name: "Edge case: empty directory",
			args: []string{"testdata/emptydir/"},
			want: `testdata/emptydir/

0 directories, 0 files
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

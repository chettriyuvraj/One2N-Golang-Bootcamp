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
			args: []string{"testdata/testdir"},
			want: `testdata/testdir
├── testdirinner
│   └── testdirinner.txt
├── testfile1.txt
└── testfile2.txt

2 directories, 3 files
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
		{
			name: "Error: file doesn't exist",
			args: []string{"testdata/x"},
			want: `testdata/x  [error opening dir]

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

func TestTreeRelativePathFlag(t *testing.T) {
	tc := []struct {
		name string
		args []string
		want string
	}{
		{
			name: "Happy case: relative path with multiple directories + files",
			args: []string{"-f", "testdata/testdir"},
			want: `testdata/testdir
├── testdata/testdir/testdirinner
│   └── testdata/testdir/testdirinner/testdirinner.txt
├── testdata/testdir/testfile1.txt
└── testdata/testdir/testfile2.txt

2 directories, 3 files
`,
		},
		{
			name: "Happy case: only files",
			args: []string{"-f", "./testdata/testdironlyfiles"},
			want: `./testdata/testdironlyfiles
└── ./testdata/testdironlyfiles/1.txt

1 directory, 1 file
`,
		},
		{
			name: "Edge case: empty directory",
			args: []string{"-f", "testdata/emptydir/"},
			want: `testdata/emptydir

0 directories, 0 files
`,
		},
		{
			name: "Error: file doesn't exist",
			args: []string{"-f", "testdata/x"},
			want: `testdata/x  [error opening dir]

0 directories, 0 files
`,
		},
	}

	for _, test := range tc {
		var b strings.Builder
		cmd := NewRootCmd()
		setFlags(cmd)
		cmd.SetOut(&b)
		cmd.SetArgs(test.args)

		cmd.Execute()
		got := b.String()
		if got != test.want {
			t.Errorf("\nTest: %s\nArgs: %s\nWant: %s\nGot: %s\n\n", test.name, test.args, test.want, got)
		}
	}
}

func TestTreeDirFlag(t *testing.T) {
	tc := []struct {
		name string
		args []string
		want string
	}{
		{
			name: "Happy case: relative path with multiple directories + files",
			args: []string{"-d", "testdata/testdir"},
			want: `testdata/testdir
└── testdirinner

2 directories
`,
		},
		{
			name: "Happy case: only files",
			args: []string{"-d", "./testdata/testdironlyfiles/"},
			want: `./testdata/testdironlyfiles/

0 directories
`,
		},
		{
			name: "Edge case: empty directory",
			args: []string{"-d", "testdata/emptydir/"},
			want: `testdata/emptydir/

0 directories
`},

		{
			name: "Error: file doesn't exist",
			args: []string{"-d", "testdata/x"},
			want: `testdata/x  [error opening dir]

0 directories
`,
		},
	}

	for _, test := range tc {
		var b strings.Builder
		cmd := NewRootCmd()
		setFlags(cmd)
		cmd.SetOut(&b)
		cmd.SetArgs(test.args)

		cmd.Execute()
		got := b.String()
		if got != test.want {
			t.Errorf("\nTest: %s\nArgs: %s\nWant: %s\nGot: %s\n\n", test.name, test.args, test.want, got)
		}
	}
}

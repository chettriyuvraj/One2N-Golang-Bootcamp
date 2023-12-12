package cmd

import "io/fs"

type DirInfo struct {
	dir          fs.DirEntry
	isLastElem   bool
	isDummyEntry bool
	DummyName    string
}

package main

import (
	"github.com/alexohneander/codesharego/cmd"
)

var (
	version    = "unknown"
	commitHash = "unknown"
)

func main() {
	cmd.Run(version, commitHash)
}

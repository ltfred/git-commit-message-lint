package main

import (
	"github.com/ltfred/git-commit-message-lint/cmd"
	"os"
)

func main() {
	err := cmd.LintCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

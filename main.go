// Package main is the entry point of the git-restore command line tool. This commands takes in 3 arguments (source repo path, destination directory and commit hash/branch/tag).
// It restore entire state of the repositroy in the destination directory at which the commit hash/branch/tag is pointing to.
package main

import (
	"fmt"
	"os"

	"github.com/Patel-Raj/git-restore/internal/logic"
	"github.com/Patel-Raj/git-restore/internal/validate"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("git-restore tool require these three arguments: source repo path, destination directory and commit hash/branch/tag")
		os.Exit(1)
	}

	sourceRepo, distiantionDir, gitObject := os.Args[1], os.Args[2], os.Args[3]

	allErrors := validate.ValidateInputs(sourceRepo, distiantionDir, gitObject)
	if allErrors != nil {
		fmt.Println(allErrors)
		os.Exit(1)
	}

	err := logic.CreateRepoCopy(sourceRepo, distiantionDir, gitObject)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Repository created successfully at", distiantionDir)
}

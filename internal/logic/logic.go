package logic

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/Patel-Raj/git-restore/internal/util"
)

func CreateRepoCopy(sourceRepo, distinationDir, gitObject string) error {
	cmd := exec.Command(util.GIT_EXECUTABLE, "cat-file", "-p", gitObject)
	cmd.Stdout = &strings.Builder{}

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("%s is not a commit hash/branch/tag", gitObject)
	}
	// TODO: Implement DFS logic
	return nil // TODO: remove this
}

package validate

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/Patel-Raj/git-restore/internal/util"
)

func ValidateInputs(repositoryPath, destinationPath, gitObject string) error {
	allErrors := make([]error, 0, 3)

	err := validateGitRepo(repositoryPath)
	if err != nil {
		allErrors = append(allErrors, err)
	}

	err = validateDir(destinationPath)
	if err != nil {
		allErrors = append(allErrors, err)
	}

	_, err = exec.LookPath(util.GIT_EXECUTABLE)
	if err != nil {
		allErrors = append(allErrors, fmt.Errorf("install git to use this tool"))
	}

	if len(allErrors) > 0 {
		return errors.Join(allErrors...)
	}

	return nil
}

func validateGitRepo(repoPath string) error {
	err := validateDir(repoPath)
	if err != nil {
		return err
	}

	gitDir := filepath.Clean(repoPath + "/.git")
	fs, err := os.Stat(gitDir)
	if err != nil || !fs.IsDir() {
		return fmt.Errorf("%s is not a git repository", repoPath)
	}

	return nil
}

func validateDir(path string) error {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("%s is not a file or directory", path)
	}

	if !fileInfo.IsDir() {
		return fmt.Errorf("%s is not a directory", path)
	}

	return nil
}

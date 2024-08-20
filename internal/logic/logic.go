package logic

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/Patel-Raj/git-restore/internal/util"
)

type Content struct {
	dir  bool
	hash string
	name string
}

func parseContent(output string) []Content {
	content := make([]Content, 0)
	lines := strings.Split(output, "\n")

	for _, line := range lines[:len(lines)-1] {
		tokens := strings.Fields(line)
		content = append(content, Content{dir: tokens[1] == "tree", hash: tokens[2], name: tokens[3]})
	}
	return content
}

func executeCommand(gitObject string) (string, error) {
	cmd := exec.Command(util.GIT_EXECUTABLE, "cat-file", "-p", gitObject)
	var out *strings.Builder = &strings.Builder{}
	cmd.Stdout = out

	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("%s is not a commit hash/branch/tag", gitObject)
	}

	return out.String(), nil
}

func handleDirCreation(root string, content Content) error {
	childDir := root + "/" + content.name

	err := os.Mkdir(childDir, 0755)
	if err != nil && !os.IsExist(err) {
		return err
	}

	return generate(childDir, content.hash)
}

func handleFileCreation(root string, content Content) error {
	output, err := executeCommand(content.hash)
	if err != nil {
		return err
	}

	err = os.WriteFile(root+"/"+content.name, []byte(output), 0644)
	if err != nil {
		return err
	}

	return nil
}

func generate(dir, gitObject string) error {
	output, err := executeCommand(gitObject)
	if err != nil {
		return err
	}

	content := parseContent(output)

	for _, obj := range content {
		var err error
		if obj.dir {
			err = handleDirCreation(dir, obj)
		} else {
			err = handleFileCreation(dir, obj)
		}

		if err != nil {
			return err
		}
	}
	return nil
}

func CreateRepoCopy(sourceRepo, distinationDir, gitObject string) error {
	rootOutput, err := executeCommand(gitObject)
	if err != nil {
		return err
	}

	absPath, _ := filepath.Abs(sourceRepo)
	repoRoot := distinationDir + "/" + filepath.Base(absPath)
	err = os.Mkdir(repoRoot, 0755)
	if err != nil && !os.IsExist(err) {
		return err
	}

	err = generate(repoRoot, strings.Split(strings.Split(rootOutput, "\n")[0], " ")[1])
	if err != nil {
		return err
	}

	return nil
}

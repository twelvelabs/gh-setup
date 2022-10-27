package git

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	"github.com/cli/safeexec"
)

// IsInstalled returns true if git is installed.
func IsInstalled() bool {
	_, err := Path()
	return err == nil
}

// IsInitialized returns true if the current working dir has been initialized.
func IsInitialized() bool {
	_, _, err := Exec("rev-parse", "--is-inside-work-tree")
	return err == nil
}

// IsDirty returns true if there are uncommitted files.
func IsDirty() bool {
	lines, _ := StatusLines()
	return len(lines) > 0
}

// StatusLines returns the result of `git status --porcelain`.
func StatusLines() ([]string, error) {
	stdout, _, err := Exec("status", "--porcelain")
	if err != nil {
		return []string{}, err
	}
	lines := []string{}
	for _, line := range strings.Split(stdout.String(), "\n") {
		if strings.TrimSpace(line) != "" {
			lines = append(lines, line)
		}
	}
	return lines, nil
}

// Path returns the path to git.
func Path() (string, error) {
	return safeexec.LookPath("git")
}

// Exec executes git with args.
// Note that any errors returned also include stderr text.
func Exec(args ...string) (bytes.Buffer, bytes.Buffer, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	path, err := Path()
	if err != nil {
		err = fmt.Errorf("could not find git executable in PATH: %w", err)
		return stdout, stderr, err
	}

	return run(path, args...)
}

func run(path string, args ...string) (bytes.Buffer, bytes.Buffer, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd := exec.Command(path, args...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		err = fmt.Errorf("failed to run git: %s. error: %w", stderr.String(), err)
		return stdout, stderr, err
	}

	return stdout, stderr, err
}

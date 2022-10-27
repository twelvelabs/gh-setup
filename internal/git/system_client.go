package git

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	"github.com/cli/safeexec"
)

type systemClient struct {
}

func (c *systemClient) IsInstalled() bool {
	_, err := path()
	return err == nil
}

func (c *systemClient) IsInitialized() bool {
	_, _, err := c.Exec("rev-parse", "--is-inside-work-tree")
	return err == nil
}

func (c *systemClient) IsDirty() bool {
	lines, _ := c.StatusLines()
	return len(lines) > 0
}

func (c *systemClient) StatusLines() ([]string, error) {
	stdout, _, err := c.Exec("status", "--porcelain")
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

func (c *systemClient) Exec(args ...string) (bytes.Buffer, bytes.Buffer, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	path, err := path()
	if err != nil {
		err = fmt.Errorf("could not find git executable in PATH: %w", err)
		return stdout, stderr, err
	}

	return run(path, args...)
}

func path() (string, error) {
	return safeexec.LookPath("git")
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

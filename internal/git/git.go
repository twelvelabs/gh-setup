package git

import (
	"bytes"
)

//go:generate moq -rm -out client_mock.go . Client

// A git CLI client.
type Client interface {
	IsInstalled() bool
	IsInitialized() bool
	IsDirty() bool
	StatusLines() ([]string, error)
	Exec(args ...string) (bytes.Buffer, bytes.Buffer, error)
}

var (
	// DefaultClient is the default Git client.
	DefaultClient Client = &systemClient{}
)

// IsInstalled returns true if git is installed.
func IsInstalled() bool {
	return DefaultClient.IsInstalled()
}

// IsInitialized returns true if the current working dir has been initialized.
func IsInitialized() bool {
	return DefaultClient.IsInitialized()
}

// IsDirty returns true if there are uncommitted files.
func IsDirty() bool {
	return DefaultClient.IsDirty()
}

// StatusLines returns the result of `git status --porcelain`.
func StatusLines() ([]string, error) {
	return DefaultClient.StatusLines()
}

// Exec executes git with args.
// Note that any errors returned also include stderr text.
func Exec(args ...string) (bytes.Buffer, bytes.Buffer, error) {
	return DefaultClient.Exec(args...)
}

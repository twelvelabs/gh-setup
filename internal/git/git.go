package git

import (
	"bytes"
)

//go:generate moq -rm -out client_mock.go . Client

// A git CLI client.
type Client interface {
	// Exec executes git with args.
	Exec(args ...string) (bytes.Buffer, bytes.Buffer, error)
	// HasRemote returns true if name has been configured as a remote.
	HasRemote(name string) bool
	// IsDirty returns true if there are uncommitted files.
	IsDirty() bool
	// IsInitialized returns true if the working dir has been initialized.
	IsInitialized() bool
	// IsInstalled returns true if git is installed.
	IsInstalled() bool
	// StatusLines returns the result of `git status --porcelain`.
	StatusLines() ([]string, error)
}

var (
	// DefaultClient is the default Git client.
	DefaultClient Client = &systemClient{}
)

// HasRemote returns true if name has been configured as a remote.
func HasRemote(name string) bool {
	return DefaultClient.HasRemote(name)
}

// Exec executes git with args.
// Note that any errors returned also include stderr text.
func Exec(args ...string) (bytes.Buffer, bytes.Buffer, error) {
	return DefaultClient.Exec(args...)
}

// IsDirty returns true if there are uncommitted files.
func IsDirty() bool {
	return DefaultClient.IsDirty()
}

// IsInitialized returns true if the working dir has been initialized.
func IsInitialized() bool {
	return DefaultClient.IsInitialized()
}

// IsInstalled returns true if git is installed.
func IsInstalled() bool {
	return DefaultClient.IsInstalled()
}

// StatusLines returns the result of `git status --porcelain`.
func StatusLines() ([]string, error) {
	return DefaultClient.StatusLines()
}

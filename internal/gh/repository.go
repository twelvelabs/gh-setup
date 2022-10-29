package gh

import "errors"

// Protocol is an enum representing the git URL protocol.
type Protocol string

const (
	ProtocolHTTPS   Protocol = "https"
	ProtocolSSH     Protocol = "ssh"
	ProtocolGit     Protocol = "git"
	ProtocolUnknown Protocol = ""
)

// Visibility is an enum representing a repo' visibility.
type Visibility string

func (v Visibility) Validate() error {
	switch v {
	case VisibilityPublic, VisibilityPrivate, VisibilityInternal:
		return nil
	default:
		return ErrInvalidVisibility
	}
}

const (
	VisibilityPublic   Visibility = "PUBLIC"
	VisibilityPrivate  Visibility = "PRIVATE"
	VisibilityInternal Visibility = "INTERNAL"
)

var (
	ErrInvalidVisibility = errors.New("invalid visibility")
)

// Repository is a GitHub repo.
type Repository struct {
	Name       string     `json:"name"`
	FullName   string     `json:"full_name"`
	Owner      *Account   `json:"owner"`
	Visibility Visibility `json:"visibility"`
	URL        string     `json:"html_url"`
	CloneURL   string     `json:"clone_url"`
	SSHURL     string     `json:"ssh_url"`
	GitURL     string     `json:"git_url"`
}

// RemoteURL returns the URL for the given protocol.
func (r *Repository) RemoteURL(protocol Protocol) string {
	// attempt to return the preferred type (if present)
	switch protocol {
	case ProtocolGit:
		return r.GitURL
	case ProtocolHTTPS:
		return r.CloneURL
	case ProtocolSSH:
		return r.SSHURL
	default:
		return r.CloneURL
	}
}

type RepositoryRequest struct {
	Name       string     `json:"name"`
	Private    bool       `json:"private"`
	Visibility Visibility `json:"visibility"`
}

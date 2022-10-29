package gh

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/cli/go-gh"
	"github.com/cli/go-gh/pkg/api"
)

//go:generate moq -rm -out client_mock.go . Client

type Client interface {
	CurrentUser() (*User, error)
	CurrentRemote() (*Repository, error)
	CreateRepo(owner string, name string, access Visibility) (*Repository, error)
	GetAccount(name string) (*Account, error)
	GetRepo(name string) (*Repository, error)
}

func NewClient(restClient RESTClient, exec ExecFunc) *SystemClient {
	if exec == nil {
		exec = gh.Exec
	}
	return &SystemClient{
		restClient: restClient,
		exec:       exec,
	}
}

type ExecFunc func(args ...string) (bytes.Buffer, bytes.Buffer, error)

type SystemClient struct {
	exec       ExecFunc
	restClient RESTClient
}

var _ Client = &SystemClient{}

func (c *SystemClient) CurrentUser() (*User, error) {
	user := &User{}
	if err := c.restClient.Get("user", user); err != nil {
		return nil, err
	}
	user.Orgs = []*Account{}
	if err := c.restClient.Get("user/orgs", &user.Orgs); err != nil {
		return nil, err
	}
	stdout, _, err := c.exec("config", "get", "git_protocol")
	if err != nil {
		return nil, err
	}
	user.GitProtocol = Protocol(strings.TrimSpace(stdout.String()))
	return user, nil
}

func (c *SystemClient) CurrentRemote() (*Repository, error) {
	r, err := gh.CurrentRepository()
	if err != nil {
		return nil, err
	}
	repo := &Repository{
		Name:     r.Name(),
		FullName: fmt.Sprintf("%s/%s", r.Owner(), r.Name()),
		Owner: &Account{
			Login: r.Owner(),
		},
		CloneURL: fmt.Sprintf("%s/%s/%s.git", r.Host(), r.Owner(), r.Name()),
	}
	return repo, nil
}

func (c *SystemClient) CreateRepo(owner string, name string, vis Visibility) (*Repository, error) {
	account, err := c.GetAccount(owner)
	if err != nil {
		return nil, err
	}

	err = vis.Validate()
	if err != nil {
		return nil, err
	}

	var path string
	switch account.Type {
	case AccountTypeOrg:
		path = fmt.Sprintf("orgs/%s/repos", owner)
	case AccountTypeUser:
		path = "user/repos"
	}
	request := &RepositoryRequest{
		Name:       name,
		Private:    (vis == VisibilityPrivate),
		Visibility: vis,
	}
	requestJSON, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	body := bytes.NewReader(requestJSON)

	repo := &Repository{}
	err = c.restClient.Post(path, body, repo)
	if err != nil {
		return nil, err
	}
	return repo, nil
}

func (c *SystemClient) GetAccount(name string) (*Account, error) {
	path := fmt.Sprintf("users/%s", name)
	account := &Account{}
	if err := c.restClient.Get(path, account); err != nil {
		httpErr := &api.HTTPError{}
		if errors.As(err, httpErr) {
			if httpErr.StatusCode == http.StatusNotFound {
				return nil, nil //nolint: nilnil
			}
		}
		return nil, err
	}
	return account, nil
}

func (c *SystemClient) GetRepo(name string) (*Repository, error) {
	path := fmt.Sprintf("repos/%s", name)
	repo := &Repository{}
	if err := c.restClient.Get(path, repo); err != nil {
		httpErr := &api.HTTPError{}
		if errors.As(err, httpErr) {
			if httpErr.StatusCode == http.StatusNotFound {
				return nil, nil //nolint: nilnil
			}
		}
		return nil, err
	}
	return repo, nil
}

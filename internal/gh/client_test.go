package gh

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"testing"

	"github.com/cli/go-gh/pkg/api"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestClient_CurrentUser(t *testing.T) {
	tests := []struct {
		desc       string
		execFunc   ExecFunc
		restClient *RESTClientMock
		expected   *User
		err        string
	}{
		{
			desc: "makes rest api and exec calls to return info about the current user",
			execFunc: func(args ...string) (bytes.Buffer, bytes.Buffer, error) {
				expected := []string{"config", "get", "git_protocol"}
				if !cmp.Equal(args, expected) {
					return bytes.Buffer{}, bytes.Buffer{}, errors.New("unexpected args")
				}
				stdout := bytes.NewBufferString("https\n")
				stderr := bytes.NewBufferString("")
				return *stdout, *stderr, nil
			},
			restClient: &RESTClientMock{
				GetFunc: func(path string, resp interface{}) error {
					if path == "user" {
						user := resp.(*User)
						user.Login = "test-user"
						return nil
					}
					if path == "user/orgs" {
						orgs := resp.(*[]*Account)
						*orgs = append(*orgs, &Account{Login: "org1"})
						*orgs = append(*orgs, &Account{Login: "org2"})
						return nil
					}
					return errors.New("unexpected path")
				},
			},
			expected: &User{
				Login: "test-user",
				Orgs: []*Account{
					{Login: "org1"},
					{Login: "org2"},
				},
				GitProtocol: ProtocolHTTPS,
			},
			err: "",
		},

		{
			desc: "returns rest api errors",
			restClient: &RESTClientMock{
				GetFunc: func(path string, resp interface{}) error {
					return errors.New("reticulating splines")
				},
			},
			expected: nil,
			err:      "reticulating splines",
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			client := NewClient(tt.restClient, tt.execFunc)
			actual, err := client.CurrentUser()

			if tt.err == "" {
				assert.NoError(t, err)
			} else {
				assert.ErrorContains(t, err, tt.err)
			}

			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestClient_CreateRepo(t *testing.T) {
	type args struct {
		owner string
		name  string
		vis   Visibility
	}
	tests := []struct {
		desc       string
		restClient *RESTClientMock
		args       args
		expected   *Repository
		err        string
	}{
		{
			desc: "can successfully create user owned repos",
			restClient: &RESTClientMock{
				GetFunc: func(path string, resp interface{}) error {
					if path == "users/test-user" {
						account := resp.(*Account)
						account.Login = "test-user"
						account.Type = "User"
						return nil
					}
					return errors.New("unexpected GET path: " + path)
				},
				PostFunc: func(path string, body io.Reader, resp interface{}) error {
					if path == "user/repos" {
						req := &RepositoryRequest{}
						_ = json.NewDecoder(body).Decode(req)

						repo := resp.(*Repository)
						repo.Name = req.Name
						repo.Visibility = req.Visibility
						return nil
					}
					return errors.New("unexpected POST path: " + path)
				},
			},
			args: args{
				owner: "test-user",
				name:  "test-repo",
				vis:   VisibilityPublic,
			},
			expected: &Repository{
				Name:       "test-repo",
				Visibility: VisibilityPublic,
			},
			err: "",
		},

		{
			desc: "can successfully create org owned repos",
			restClient: &RESTClientMock{
				GetFunc: func(path string, resp interface{}) error {
					if path == "users/test-org" {
						account := resp.(*Account)
						account.Login = "test-org"
						account.Type = "Organization"
						return nil
					}
					return errors.New("unexpected GET path: " + path)
				},
				PostFunc: func(path string, body io.Reader, resp interface{}) error {
					if path == "orgs/test-org/repos" {
						req := &RepositoryRequest{}
						_ = json.NewDecoder(body).Decode(req)

						repo := resp.(*Repository)
						repo.Name = req.Name
						repo.Visibility = req.Visibility
						return nil
					}
					return errors.New("unexpected POST path: " + path)
				},
			},
			args: args{
				owner: "test-org",
				name:  "test-repo",
				vis:   VisibilityPublic,
			},
			expected: &Repository{
				Name:       "test-repo",
				Visibility: VisibilityPublic,
			},
			err: "",
		},

		{
			desc: "returns rest api errors",
			restClient: &RESTClientMock{
				GetFunc: func(path string, resp interface{}) error {
					return errors.New("reticulating splines")
				},
			},
			args: args{
				owner: "test-user",
				name:  "test-repo",
				vis:   VisibilityPublic,
			},
			expected: nil,
			err:      "reticulating splines",
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			client := NewClient(tt.restClient, nil)
			actual, err := client.CreateRepo(tt.args.owner, tt.args.name, tt.args.vis)

			if tt.err == "" {
				assert.NoError(t, err)
			} else {
				assert.ErrorContains(t, err, tt.err)
			}

			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestClient_GetAccount(t *testing.T) {
	tests := []struct {
		desc       string
		execFunc   ExecFunc
		restClient *RESTClientMock
		expected   *Account
		err        string
	}{
		{
			desc: "calls the rest api and returns an account",
			restClient: &RESTClientMock{
				GetFunc: func(path string, resp interface{}) error {
					if path == "users/someone" {
						account := resp.(*Account)
						account.Login = "someone"
						return nil
					}
					return errors.New("unexpected path")
				},
			},
			expected: &Account{
				Login: "someone",
			},
			err: "",
		},

		{
			desc: "gracefully handles 404 responses from the api",
			restClient: &RESTClientMock{
				GetFunc: func(path string, resp interface{}) error {
					return api.HTTPError{
						Message:    "Not Found",
						StatusCode: 404,
					}
				},
			},
			expected: nil,
			err:      "",
		},

		{
			desc: "returns all other api errors",
			restClient: &RESTClientMock{
				GetFunc: func(path string, resp interface{}) error {
					return errors.New("reticulating splines")
				},
			},
			expected: nil,
			err:      "reticulating splines",
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			client := NewClient(tt.restClient, tt.execFunc)
			actual, err := client.GetAccount("someone")

			if tt.err == "" {
				assert.NoError(t, err)
			} else {
				assert.ErrorContains(t, err, tt.err)
			}

			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestClient_GetRepo(t *testing.T) {
	tests := []struct {
		desc       string
		execFunc   ExecFunc
		restClient *RESTClientMock
		expected   *Repository
		err        string
	}{
		{
			desc: "calls the rest api and returns a repo",
			restClient: &RESTClientMock{
				GetFunc: func(path string, resp interface{}) error {
					if path == "repos/test-owner/test-repo" {
						repo := resp.(*Repository)
						repo.Name = "test-repo"
						repo.FullName = "test-owner/test-repo"
						return nil
					}
					return errors.New("unexpected path")
				},
			},
			expected: &Repository{
				Name:     "test-repo",
				FullName: "test-owner/test-repo",
			},
			err: "",
		},

		{
			desc: "gracefully handles 404 responses from the api",
			restClient: &RESTClientMock{
				GetFunc: func(path string, resp interface{}) error {
					return api.HTTPError{
						Message:    "Not Found",
						StatusCode: 404,
					}
				},
			},
			expected: nil,
			err:      "",
		},

		{
			desc: "returns all other api errors",
			restClient: &RESTClientMock{
				GetFunc: func(path string, resp interface{}) error {
					return errors.New("reticulating splines")
				},
			},
			expected: nil,
			err:      "reticulating splines",
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			client := NewClient(tt.restClient, tt.execFunc)
			actual, err := client.GetRepo("test-owner/test-repo")

			if tt.err == "" {
				assert.NoError(t, err)
			} else {
				assert.ErrorContains(t, err, tt.err)
			}

			assert.Equal(t, tt.expected, actual)
		})
	}
}

package cmd

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/twelvelabs/termite/testutil"
	uimock "github.com/twelvelabs/termite/ui/mock"

	"github.com/twelvelabs/gh-setup/internal/core"
	"github.com/twelvelabs/gh-setup/internal/gh"
	"github.com/twelvelabs/gh-setup/internal/git"
)

func TestMain(m *testing.M) {
	os.Setenv("APP_ENV", "test")
	os.Exit(m.Run())
}

func TestRootAction_Run(t *testing.T) {
	tests := []struct {
		desc       string
		setup      func(t *testing.T, a *RootAction)
		assertions func(t *testing.T, a *RootAction)
		err        string
	}{
		{
			desc: "returns error if git not found",
			setup: func(t *testing.T, a *RootAction) {
				t.Helper()

				gc := a.GitClient.(*git.ClientMock)
				gc.IsInstalledFunc = func() bool {
					return false
				}
			},
			err: "could not find git",
		},

		{
			desc: "prompts to run git init if needed",
			setup: func(t *testing.T, a *RootAction) {
				t.Helper()

				a.GhClient = NewClientMock()
				a.GitClient = git.DefaultClient

				p := a.Prompter.(*uimock.PrompterMock)
				p.ConfirmFunc = func(msg string, value bool, help string) (bool, error) {
					switch msg {
					case "Initialize the repo?":
						return true, nil
					case "Create a new repo on GitHub?":
						return false, nil
					default:
						panic(fmt.Errorf("unexpected confirm call: %s", msg))
					}
				}

				assert.Equal(t, false, a.GitClient.IsInitialized())
			},
			assertions: func(t *testing.T, a *RootAction) {
				t.Helper()

				p := a.Prompter.(*uimock.PrompterMock)
				assert.Equal(t, 2, len(p.ConfirmCalls()))
				assert.Equal(t, true, a.GitClient.IsInitialized())
			},
			err: "aborted",
		},
		{
			desc: "aborts early if user does not want to run git init",
			setup: func(t *testing.T, a *RootAction) {
				t.Helper()

				a.GhClient = NewClientMock()
				a.GitClient = git.DefaultClient

				p := a.Prompter.(*uimock.PrompterMock)
				p.ConfirmFunc = func(msg string, value bool, help string) (bool, error) {
					switch msg {
					case "Initialize the repo?":
						return false, nil
					default:
						panic(fmt.Errorf("unexpected confirm call: %s", msg))
					}
				}

				assert.Equal(t, false, a.GitClient.IsInitialized())
			},
			assertions: func(t *testing.T, a *RootAction) {
				t.Helper()

				p := a.Prompter.(*uimock.PrompterMock)
				assert.Equal(t, 1, len(p.ConfirmCalls()))
				assert.Equal(t, false, a.GitClient.IsInitialized())
			},
			err: "aborted",
		},

		{
			desc: "creates a new repo",
			setup: func(t *testing.T, a *RootAction) {
				t.Helper()

				a.GitClient = git.DefaultClient

				a.GhClient = NewClientMock()
				ghc := a.GhClient.(*gh.ClientMock)
				ghc.CreateRepoFunc = func(owner, name string, vis gh.Visibility) (*gh.Repository, error) {
					url := fmt.Sprintf("http://github.com/%s/%s", owner, name)
					repo := &gh.Repository{
						Name:       name,
						Visibility: vis,
						URL:        url,
						CloneURL:   url + ".git",
					}
					return repo, nil
				}

				p := a.Prompter.(*uimock.PrompterMock)
				p.ConfirmFunc = func(msg string, value bool, help string) (bool, error) {
					switch msg {
					case "Create a new repo on GitHub?":
						return true, nil
					case "Add and commit?":
						return true, nil
					case "Push local commits to the remote?":
						return false, nil
					default:
						panic(fmt.Errorf("unexpected confirm call: %s", msg))
					}
				}

				p.InputFunc = func(msg, value, help string) (string, error) {
					switch msg {
					case "GitHub repo name":
						return value, nil
					case "Commit message":
						return value, nil
					default:
						panic(fmt.Errorf("unexpected input call: %s", msg))
					}
				}

				p.SelectFunc = func(msg string, options []string, value, help string) (string, error) {
					switch msg {
					case "GitHub repo owner":
						return value, nil
					case "GitHub repo visibility":
						return value, nil
					default:
						panic(fmt.Errorf("unexpected select call: %s", msg))
					}
				}

				assert.Equal(t, false, a.GitClient.IsInitialized())
				// dump two untracked files in there
				_ = os.WriteFile("foo.txt", []byte("aaa"), 0600)
				_ = os.WriteFile("bar.txt", []byte("bbb"), 0600)
				_, _, err := a.GitClient.Exec("init")
				assert.NoError(t, err)
				assert.Equal(t, true, a.GitClient.IsInitialized())
				assert.Equal(t, false, git.HasRemote("origin"))
				assert.Equal(t, true, a.GitClient.IsDirty())
			},
			assertions: func(t *testing.T, a *RootAction) {
				t.Helper()

				p := a.Prompter.(*uimock.PrompterMock)
				assert.Equal(t, 3, len(p.ConfirmCalls()))
				assert.Equal(t, 2, len(p.InputCalls()))
				assert.Equal(t, 2, len(p.SelectCalls()))

				assert.Equal(t, true, git.IsInitialized())
				assert.Equal(t, true, git.HasRemote("origin"))
				assert.Equal(t, false, git.IsDirty())
			},
			err: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			testutil.InTempDir(t, func(tmpDir string) {
				app := core.NewTestApp()
				action := NewRootAction(app)

				// pre-run setup or assertions
				if tt.setup != nil {
					tt.setup(t, action)
				}

				// run the action
				err := action.Run()
				// assert error
				if tt.err == "" {
					require.NoError(t, err)
				} else {
					require.ErrorContains(t, err, tt.err)
				}

				// post-run assertions
				if tt.assertions != nil {
					tt.assertions(t, action)
				}
			})
		})
	}
}

func NewClientMock() *gh.ClientMock {
	return &gh.ClientMock{
		CurrentRemoteFunc: func() (*gh.Repository, error) {
			return nil, nil
		},
		CurrentUserFunc: func() (*gh.User, error) {
			return &gh.User{
				Login: "test-user",
				Orgs: []*gh.Account{
					{Login: "org1"},
				},
				GitProtocol: gh.ProtocolHTTPS,
			}, nil
		},
		CreateRepoFunc: func(owner string, name string, access gh.Visibility) (*gh.Repository, error) {
			return nil, nil
		},
		GetAccountFunc: func(name string) (*gh.Account, error) {
			return nil, nil
		},
		GetRepoFunc: func(name string) (*gh.Repository, error) {
			return nil, nil
		},
	}
}

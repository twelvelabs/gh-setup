package cmd

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/twelvelabs/gh-setup/internal/core"
	"github.com/twelvelabs/gh-setup/internal/git"
	"github.com/twelvelabs/gh-setup/internal/prompt"
	"github.com/twelvelabs/gh-setup/internal/testutil"
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

				a.GitClient = git.DefaultClient

				p := a.Prompter.(*prompt.PrompterMock)
				p.ConfirmFunc = prompt.NewConfirmFuncSet(
					prompt.NewConfirmFunc(true, nil),
					prompt.NewConfirmFunc(false, nil),
				)

				assert.Equal(t, false, a.GitClient.IsInitialized())
			},
			assertions: func(t *testing.T, a *RootAction) {
				t.Helper()

				p := a.Prompter.(*prompt.PrompterMock)
				cc := p.ConfirmCalls()
				assert.Equal(t, 2, len(cc))
				assert.Equal(t, "Initialize the repo?", cc[0].Msg)
				assert.Equal(t, "Create a remote on GitHub?", cc[1].Msg)

				assert.Equal(t, true, a.GitClient.IsInitialized())
			},
		},
		{
			desc: "aborts early if user does not want to run git init",
			setup: func(t *testing.T, a *RootAction) {
				t.Helper()

				a.GitClient = git.DefaultClient

				p := a.Prompter.(*prompt.PrompterMock)
				p.ConfirmFunc = prompt.NewConfirmFuncSet(
					prompt.NewConfirmFunc(false, nil),
				)

				assert.Equal(t, false, a.GitClient.IsInitialized())
			},
			assertions: func(t *testing.T, a *RootAction) {
				t.Helper()

				p := a.Prompter.(*prompt.PrompterMock)
				cc := p.ConfirmCalls()
				assert.Equal(t, 1, len(cc))
				assert.Equal(t, "Initialize the repo?", cc[0].Msg)

				assert.Equal(t, false, a.GitClient.IsInitialized())
			},
		},

		{
			desc: "prompts to run git commit if the working directory is dirty",
			setup: func(t *testing.T, a *RootAction) {
				t.Helper()

				a.GitClient = git.DefaultClient

				p := a.Prompter.(*prompt.PrompterMock)
				p.ConfirmFunc = prompt.NewConfirmFuncSet(
					prompt.NewConfirmFunc(true, nil),
					prompt.NewConfirmFunc(false, nil),
				)
				p.InputFunc = prompt.NewInputFunc("custom commit msg", nil)

				assert.Equal(t, false, a.GitClient.IsInitialized())
				// dump two untracked files in there
				_ = os.WriteFile("foo.txt", []byte("aaa"), 0600)
				_ = os.WriteFile("bar.txt", []byte("bbb"), 0600)
				_, _, err := a.GitClient.Exec("init")
				assert.NoError(t, err)
				assert.Equal(t, true, a.GitClient.IsDirty())
			},
			assertions: func(t *testing.T, a *RootAction) {
				t.Helper()

				p := a.Prompter.(*prompt.PrompterMock)
				cc := p.ConfirmCalls()
				assert.Equal(t, 2, len(cc))
				assert.Equal(t, "Add and commit?", cc[0].Msg)
				assert.Equal(t, "Create a remote on GitHub?", cc[1].Msg)

				ic := p.InputCalls()
				assert.Equal(t, 1, len(ic))
				assert.Equal(t, "Commit message", ic[0].Msg)

				assert.Equal(t, true, git.IsInitialized())
				assert.Equal(t, false, git.IsDirty())
			},
		},
		{
			desc: "aborts early if user does not want to run git commit",
			setup: func(t *testing.T, a *RootAction) {
				t.Helper()

				a.GitClient = git.DefaultClient

				p := a.Prompter.(*prompt.PrompterMock)
				p.ConfirmFunc = prompt.NewConfirmFuncSet(
					prompt.NewConfirmFunc(false, nil),
				)

				assert.Equal(t, false, a.GitClient.IsInitialized())
				// dump two untracked files in there
				_ = os.WriteFile("foo.txt", []byte("aaa"), 0600)
				_ = os.WriteFile("bar.txt", []byte("bbb"), 0600)
				_, _, err := a.GitClient.Exec("init")
				assert.NoError(t, err)
				assert.Equal(t, true, a.GitClient.IsDirty())
			},
			assertions: func(t *testing.T, a *RootAction) {
				t.Helper()

				p := a.Prompter.(*prompt.PrompterMock)
				cc := p.ConfirmCalls()
				assert.Equal(t, 1, len(cc))
				assert.Equal(t, "Add and commit?", cc[0].Msg)

				assert.Equal(t, true, git.IsInitialized())
				assert.Equal(t, true, git.IsDirty())
			},
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

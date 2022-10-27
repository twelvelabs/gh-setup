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

func TestRootAction(t *testing.T) {
	tests := []struct {
		desc     string
		before   func(t *testing.T, a *RootAction)
		prompter *prompt.PrompterMock
		after    func(t *testing.T, a *RootAction, p *prompt.PrompterMock)
		err      string
	}{
		{
			desc: "running in an empty directory only prompts to init",
			before: func(t *testing.T, a *RootAction) {
				t.Helper()
				assert.Equal(t, false, git.IsInitialized())
			},
			prompter: &prompt.PrompterMock{
				ConfirmFunc: prompt.NewConfirmFuncSet(
					prompt.NewConfirmFunc(true, nil),
				),
			},
			after: func(t *testing.T, a *RootAction, p *prompt.PrompterMock) {
				t.Helper()
				prompter := a.Prompter.(*prompt.PrompterMock)
				calls := prompter.ConfirmCalls()

				assert.Equal(t, 1, len(calls))
				assert.Equal(t, "Initialize the repo?", calls[0].Msg)

				assert.Equal(t, true, git.IsInitialized())
			},
		},

		{
			desc: "running in a directory with files also prompts to commit",
			before: func(t *testing.T, a *RootAction) {
				t.Helper()
				assert.Equal(t, false, git.IsInitialized())
				_ = os.WriteFile("foo.txt", []byte("aaa"), 0600)
				_ = os.WriteFile("bar.txt", []byte("bbb"), 0600)
			},
			prompter: &prompt.PrompterMock{
				ConfirmFunc: prompt.NewConfirmFuncSet(
					prompt.NewConfirmFunc(true, nil),
					prompt.NewConfirmFunc(true, nil),
				),
				InputFunc: prompt.NewInputFunc("custom commit msg", nil),
			},
			after: func(t *testing.T, a *RootAction, p *prompt.PrompterMock) {
				t.Helper()

				cc := p.ConfirmCalls()
				assert.Equal(t, 2, len(cc))
				assert.Equal(t, "Initialize the repo?", cc[0].Msg)
				assert.Equal(t, "Add and commit?", cc[1].Msg)

				ic := p.InputCalls()
				assert.Equal(t, 1, len(ic))
				assert.Equal(t, "Commit message", ic[0].Msg)

				assert.Equal(t, true, git.IsInitialized())
				assert.Equal(t, false, git.IsDirty())
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			testutil.InTempDir(t, func(tmpDir string) {
				app := core.NewTestApp()
				action := NewRootAction(app)
				action.Prompter = tt.prompter

				// pre-run setup or assertions
				tt.before(t, action)
				// run the action
				err := action.Run()
				// assert error
				if tt.err == "" {
					require.NoError(t, err)
				} else {
					require.ErrorContains(t, err, tt.err)
				}
				// post-run assertions
				tt.after(t, action, tt.prompter)
			})
		})
	}
}

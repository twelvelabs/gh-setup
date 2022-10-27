package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/twelvelabs/gh-setup/internal/core"
	"github.com/twelvelabs/gh-setup/internal/git"
	"github.com/twelvelabs/gh-setup/internal/iostreams"
	"github.com/twelvelabs/gh-setup/internal/prompt"
)

func NewRootCmd(app *core.App) *cobra.Command {
	action := NewRootAction(app)

	cmd := &cobra.Command{
		Use:   "gh-setup",
		Short: "Setup new GitHub repositories",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := action.Setup(cmd, args); err != nil {
				return err
			}
			if err := action.Validate(); err != nil {
				return err
			}
			if err := action.Run(); err != nil {
				return err
			}
			return nil
		},
		Version:      "X.X.X",
		SilenceUsage: true,
	}

	cmd.Flags().BoolVar(&action.NoPrompt, "no-prompt", false, "Do not prompt for input")
	cmd.Flags().Lookup("no-prompt").NoOptDefVal = "true"

	return cmd
}

func NewRootAction(app *core.App) *RootAction {
	return &RootAction{
		IO:        app.IO,
		Logger:    app.Logger,
		Prompter:  app.Prompter,
		GitClient: app.GitClient,
	}
}

type RootAction struct {
	IO        *iostreams.IOStreams
	Logger    *iostreams.IconLogger
	Prompter  prompt.Prompter
	GitClient git.Client

	NoPrompt bool
}

func (a *RootAction) Setup(cmd *cobra.Command, args []string) error {
	if a.NoPrompt {
		a.IO.SetInteractive(false)
	}
	return nil
}
func (a *RootAction) Validate() error {
	return nil
}
func (a *RootAction) Run() error {
	if !a.GitClient.IsInstalled() {
		return fmt.Errorf("could not find git executable in PATH")
	}

	if !a.GitClient.IsInitialized() {
		resp, err := a.Prompter.Confirm("Initialize the repo?", true, "")
		if err != nil {
			return err
		}
		if !resp {
			a.Logger.Failure("Unable to continue until the working directory is initialized.\n")
			return nil
		}
		_, _, err = a.GitClient.Exec("init")
		if err != nil {
			return err
		}
	}

	lines, err := a.GitClient.StatusLines()
	if err != nil {
		return err
	}
	if len(lines) > 0 {
		ok, err := a.promptToCommit(lines)
		if err != nil {
			return err
		}
		if !ok {
			a.Logger.Failure("Unable to continue until the working directory is clean.\n")
			return nil
		}
		err = a.commit()
		if err != nil {
			return err
		}
	}

	a.Logger.Success("Setup complete.\n")

	// client, err := gh.RESTClient(nil)
	// if err != nil {
	// 	fmt.Fprintln(a.IO.Err, err)
	// 	os.Exit(1)
	// }
	// response := struct{ Login string }{}
	// err = client.Get("user", &response)
	// if err != nil {
	// 	fmt.Fprintln(a.IO.Err, err)
	// 	os.Exit(1)
	// }

	// fmt.Fprintf(a.IO.Err, "running as %s\n", response.Login)
	return nil
}

func (a *RootAction) promptToCommit(lines []string) (bool, error) {
	a.Logger.Info("There are uncommitted files in the working directory:\n")
	fmt.Fprintf(a.IO.Err, "\n")
	for _, line := range lines {
		fmt.Fprintf(a.IO.Err, "%s\n", line)
	}
	fmt.Fprintf(a.IO.Err, "\n")
	return a.Prompter.Confirm("Add and commit?", true, "")
}

func (a *RootAction) commit() error {
	_, _, err := a.GitClient.Exec("add", ".")
	if err != nil {
		return err
	}
	msg, err := a.Prompter.Input("Commit message", "Initial commit", "")
	if err != nil {
		return err
	}
	args := []string{"commit", "-m", msg}
	if os.Getenv("APP_ENV") == "test" {
		// Don't sign when testing because some folks (:raise_hand:)
		// use a YubiKey and it will block waiting for approval.
		args = append(args, "--no-gpg-sign", "--no-verify")
	}
	// Starting a progress indicator for the YubiKey folks as a reminder
	// that they need to touch to approve.
	a.IO.StartProgressIndicatorWithLabel("Committing")
	_, _, err = a.GitClient.Exec(args...)
	a.IO.StopProgressIndicator()
	return err
}

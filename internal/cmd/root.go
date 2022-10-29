package cmd

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"

	"github.com/twelvelabs/gh-setup/internal/core"
	"github.com/twelvelabs/gh-setup/internal/gh"
	"github.com/twelvelabs/gh-setup/internal/git"
	"github.com/twelvelabs/gh-setup/internal/iostreams"
	"github.com/twelvelabs/gh-setup/internal/prompt"
)

var (
	ErrAborted = errors.New("aborted")
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
		GhClient:  app.GhClient,
		GitClient: app.GitClient,
	}
}

type RootAction struct {
	IO        *iostreams.IOStreams
	Logger    *iostreams.IconLogger
	Prompter  prompt.Prompter
	GhClient  gh.Client
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
	if err := a.ensureGitInstalled(); err != nil {
		return err
	}

	if err := a.ensureWorkingDirInit(); err != nil {
		return err
	}

	if err := a.ensureWorkingDirClean(); err != nil {
		return err
	}

	if err := a.ensureRemote(); err != nil {
		return err
	}

	if err := a.ensurePush(); err != nil {
		return err
	}

	a.Logger.Success("Setup complete.\n")
	return nil
}

func (a *RootAction) ensureGitInstalled() error {
	if a.GitClient.IsInstalled() {
		return nil // git installed
	}
	return fmt.Errorf("could not find git executable in PATH")
}

func (a *RootAction) ensureWorkingDirInit() error {
	if a.GitClient.IsInitialized() {
		return nil // working dir already initialized
	}
	ok, err := a.Prompter.Confirm("Initialize the repo?", true, "")
	if err != nil {
		return err
	}
	if !ok {
		a.Logger.Failure("Unable to continue until the working directory is initialized.\n")
		return ErrAborted
	}
	_, _, err = a.GitClient.Exec("init")
	if err != nil {
		return err
	}
	return nil
}

func (a *RootAction) ensureWorkingDirClean() error {
	if !a.GitClient.IsDirty() {
		return nil // working dir clean
	}

	lines, err := a.GitClient.StatusLines()
	if err != nil {
		return err
	}
	ok, err := a.promptToCommit(lines)
	if err != nil {
		return err
	}
	if !ok {
		a.Logger.Failure("Unable to continue until the working directory is clean.\n")
		return ErrAborted
	}
	err = a.commit()
	if err != nil {
		return err
	}
	return nil
}

func (a *RootAction) ensureRemote() error {
	repo, _ := a.GhClient.CurrentRemote()
	if repo != nil {
		return nil // remote already configured - nothing to do
	}

	// 1. Resolve the assumed repoName of the working directory.
	user, err := a.GhClient.CurrentUser()
	if err != nil {
		return err
	}
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	dir = filepath.Base(dir)
	repoName := fmt.Sprintf("%s/%s", user.Login, dir)

	// 2. Check to see if a repo already exists with that name.
	repo, err = a.GhClient.GetRepo(repoName)
	if err != nil {
		return err
	}
	if repo != nil {
		// 2a. Prompt to select existing repo
		a.Logger.Info("A repo named '%s' already exists on GitHub.\n", repoName)
		ok, err := a.Prompter.Confirm("Add it as a remote?", true, "")
		if err != nil {
			return err
		}
		if ok {
			// 2b. Set remote...
			if err := a.setRemote(repo, user); err != nil {
				return err
			}
			return nil
		}
	}

	// 3. Prompt to create new repo.
	ok, err := a.Prompter.Confirm("Create a new repo on GitHub?", true, "")
	if err != nil {
		return err
	}
	if !ok {
		a.Logger.Failure("Unable to continue until a remote has been configured.\n")
		return ErrAborted
	}

	owners := []string{user.Login}
	for _, org := range user.Orgs {
		owners = append(owners, org.Login)
	}
	owner, err := a.Prompter.Select("GitHub repo owner", owners, user.Login, "")
	if err != nil {
		return err
	}
	name, err := a.Prompter.Input("GitHub repo name", dir, "")
	if err != nil {
		return err
	}
	vis, err := a.Prompter.Select(
		"GitHub repo visibility",
		[]string{"Public", "Private", "Internal"},
		"Public",
		"",
	)
	if err != nil {
		return err
	}
	visibility := gh.Visibility(strings.ToUpper(vis))

	a.IO.StartProgressIndicatorWithLabel("Creating repo")
	repo, err = a.GhClient.CreateRepo(owner, name, visibility)
	a.IO.StopProgressIndicator()
	if err != nil {
		return err
	}
	a.Logger.Success("Repo created: %s\n", repo.URL)

	if err := a.setRemote(repo, user); err != nil {
		return err
	}
	return nil
}

func (a *RootAction) ensurePush() error {
	if !a.GitClient.HasCommits() {
		return nil // no commits - nothing to push
	}
	ok, err := a.Prompter.Confirm("Push local commits to the remote?", true, "")
	if err != nil {
		return err
	}
	if ok {
		_, _, err = a.GitClient.Exec("push", "-u", "origin", "HEAD")
		return err
	}
	return nil
}

func (a *RootAction) setRemote(repo *gh.Repository, user *gh.User) error {
	remote := "origin"
	url := repo.RemoteURL(user.GitProtocol)
	_, _, err := a.GitClient.Exec("remote", "add", remote, url)
	if err != nil {
		return err
	}
	a.Logger.Success("Remote added: %s %s\n", remote, url)
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

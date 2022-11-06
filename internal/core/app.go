package core

import (
	"fmt"

	"github.com/twelvelabs/termite/ioutil"
	"github.com/twelvelabs/termite/ui"

	"github.com/twelvelabs/gh-setup/internal/gh"
	"github.com/twelvelabs/gh-setup/internal/git"
)

type App struct {
	IO           *ioutil.IOStreams
	Messenger    *ui.Messenger
	Prompter     ui.Prompter
	GhClient     gh.Client
	GhRestClient gh.RESTClient
	GitClient    git.Client
}

func NewApp() (*App, error) {
	ios := ioutil.System()
	messenger := ui.NewMessenger(ios)
	prompter := ui.NewSurveyPrompter(ios.In, ios.Out, ios.Err, ios)
	gitClient := git.DefaultClient

	ghRestClient, err := gh.NewRESTClient()
	if err != nil {
		return nil, fmt.Errorf("gh: %w", err)
	}
	ghClient := gh.NewClient(ghRestClient, nil)

	app := &App{
		IO:           ios,
		Messenger:    messenger,
		Prompter:     prompter,
		GhClient:     ghClient,
		GhRestClient: ghRestClient,
		GitClient:    gitClient,
	}
	return app, nil
}

func NewTestApp() *App {
	ios := ioutil.Test()
	messenger := ui.NewMessenger(ios)
	prompter := ui.NewPrompterMock()
	ghClient := &gh.ClientMock{}
	ghRestClient := &gh.RESTClientMock{}
	gitClient := &git.ClientMock{}

	return &App{
		IO:           ios,
		Messenger:    messenger,
		Prompter:     prompter,
		GhClient:     ghClient,
		GhRestClient: ghRestClient,
		GitClient:    gitClient,
	}
}

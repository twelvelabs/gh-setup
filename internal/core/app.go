package core

import (
	"fmt"

	"github.com/twelvelabs/gh-setup/internal/gh"
	"github.com/twelvelabs/gh-setup/internal/git"
	"github.com/twelvelabs/gh-setup/internal/iostreams"
	"github.com/twelvelabs/gh-setup/internal/prompt"
)

type App struct {
	IO           *iostreams.IOStreams
	Logger       *iostreams.IconLogger
	Prompter     prompt.Prompter
	GhClient     gh.Client
	GhRestClient gh.RESTClient
	GitClient    git.Client
}

func NewApp() (*App, error) {
	ios := iostreams.System()
	logger := iostreams.NewIconLogger(ios)
	prompter := prompt.NewSurveyPrompter(ios.In, ios.Out, ios.Err, ios)
	gitClient := git.DefaultClient

	ghRestClient, err := gh.NewRESTClient()
	if err != nil {
		return nil, fmt.Errorf("gh: %w", err)
	}
	ghClient := gh.NewClient(ghRestClient, nil)

	app := &App{
		IO:           ios,
		Logger:       logger,
		Prompter:     prompter,
		GhClient:     ghClient,
		GhRestClient: ghRestClient,
		GitClient:    gitClient,
	}
	return app, nil
}

func NewTestApp() *App {
	ios := iostreams.Test()
	logger := iostreams.NewIconLogger(ios)
	prompter := &prompt.PrompterMock{}
	ghClient := &gh.ClientMock{}
	ghRestClient := &gh.RESTClientMock{}
	gitClient := &git.ClientMock{}

	return &App{
		IO:           ios,
		Logger:       logger,
		Prompter:     prompter,
		GhClient:     ghClient,
		GhRestClient: ghRestClient,
		GitClient:    gitClient,
	}
}

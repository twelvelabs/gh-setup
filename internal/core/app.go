package core

import (
	"github.com/twelvelabs/gh-setup/internal/git"
	"github.com/twelvelabs/gh-setup/internal/iostreams"
	"github.com/twelvelabs/gh-setup/internal/prompt"
)

type App struct {
	IO        *iostreams.IOStreams
	Logger    *iostreams.IconLogger
	Prompter  prompt.Prompter
	GitClient git.Client
}

func NewApp() *App {
	ios := iostreams.System()
	logger := iostreams.NewIconLogger(ios)
	prompter := prompt.NewSurveyPrompter(ios.In, ios.Out, ios.Err, ios)
	gitClient := git.DefaultClient

	return &App{
		IO:        ios,
		Logger:    logger,
		Prompter:  prompter,
		GitClient: gitClient,
	}
}

func NewTestApp() *App {
	ios := iostreams.Test()
	logger := iostreams.NewIconLogger(ios)
	prompter := &prompt.PrompterMock{}
	gitClient := &git.ClientMock{}

	return &App{
		IO:        ios,
		Logger:    logger,
		Prompter:  prompter,
		GitClient: gitClient,
	}
}

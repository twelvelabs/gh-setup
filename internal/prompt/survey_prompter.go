package prompt

import (
	"fmt"
	"io"
	"strings"

	"github.com/AlecAivazis/survey/v2" //cspell: disable-line
)

type fileReader interface {
	io.Reader
	Fd() uintptr
}

type fileWriter interface {
	io.Writer
	Fd() uintptr
}

type ioSession interface {
	IsInteractive() bool
}

func NewSurveyPrompter(in fileReader, out fileWriter, err fileWriter, ios ioSession) *SurveyPrompter {
	return &SurveyPrompter{
		stdin:  in,
		stdout: out,
		stderr: err,
		ios:    ios,
	}
}

type SurveyPrompter struct {
	stdin  fileReader
	stdout fileWriter
	stderr fileWriter
	ios    ioSession
}

func (p *SurveyPrompter) Confirm(prompt string, defaultValue bool, help string) (bool, error) {
	result := defaultValue
	err := p.ask(&survey.Confirm{
		Message: prompt,
		Help:    help,
		Default: defaultValue,
	}, &result)
	return result, err
}

func (p *SurveyPrompter) Input(prompt string, defaultValue string, help string) (string, error) {
	result := defaultValue
	err := p.ask(&survey.Input{
		Message: prompt,
		Help:    help,
		Default: defaultValue,
	}, &result)
	return result, err
}

func (p *SurveyPrompter) MultiSelect(
	prompt string, options []string, defaultValues []string, help string,
) ([]string, error) {
	result := defaultValues
	err := p.ask(&survey.MultiSelect{
		Message: prompt,
		Help:    help,
		Options: options,
		Default: defaultValues,
	}, &result)
	return result, err
}

func (p *SurveyPrompter) Select(
	prompt string, options []string, defaultValue string, help string,
) (string, error) {
	result := defaultValue
	err := p.ask(&survey.Select{
		Message: prompt,
		Help:    help,
		Options: options,
		Default: defaultValue,
	}, &result)
	return result, err
}

func (p *SurveyPrompter) ask(q survey.Prompt, response interface{}) error {
	if !p.ios.IsInteractive() {
		return nil
	}
	// survey.AskOne() doesn't allow passing in a transform func,
	// so we need to call survey.Ask().
	qs := []*survey.Question{
		{
			Prompt:    q,
			Transform: TrimSpace,
		},
	}
	err := survey.Ask(qs, response, survey.WithStdio(p.stdin, p.stdout, p.stderr))
	if err == nil {
		return nil
	}
	return fmt.Errorf("could not prompt: %w", err)
}

var (
	_ survey.Transformer = TrimSpace
)

// Custom survey.Transformer that removes leading and trailing whitespace
// from string values. Non-string values are a no-op.
func TrimSpace(val any) any {
	if str, ok := val.(string); ok {
		return strings.TrimSpace(str)
	}
	return val
}

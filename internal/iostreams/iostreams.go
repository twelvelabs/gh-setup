package iostreams

/*
This file started out as a copy of https://github.com/cli/cli/blob/trunk/pkg/iostreams/iostreams.go
Original license:

MIT License

Copyright (c) 2019 GitHub Inc.

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.
*/

import (
	"bytes"
	"io"
	"os"
	"sync"
	"time"

	"github.com/briandowns/spinner"
	"github.com/mattn/go-isatty" // cspell: disable-line
)

// IOStream represents an input or output stream.
type IOStream interface {
	io.Reader
	io.Writer
	Fd() uintptr
	String() string
}

// Container for the three main CLI I/O streams.
type IOStreams struct {
	// os.Stdin (or mock when unit testing)
	In IOStream
	// os.Stdout (or mock when unit testing)
	Out IOStream
	// os.Stderr (or mock when unit testing)
	Err IOStream

	colorEnabled bool

	isInteractiveOverride bool
	isInteractive         bool

	progressIndicatorEnabled bool
	progressIndicator        *spinner.Spinner
	progressIndicatorMu      sync.Mutex

	stdinTTYOverride  bool
	stdinIsTTY        bool
	stdoutTTYOverride bool
	stdoutIsTTY       bool
	stderrTTYOverride bool
	stderrIsTTY       bool
}

// IsColorEnabled returns true if color output is enabled.
func (s *IOStreams) IsColorEnabled() bool {
	return s.colorEnabled
}

// SetColorEnabled sets whether color is enabled.
func (s *IOStreams) SetColorEnabled(v bool) {
	s.colorEnabled = v
}

// Formatter returns a ANSI string formatter.
func (s *IOStreams) Formatter() *Formatter {
	return NewFormatter(s.IsColorEnabled())
}

// SetStdinTTY explicitly flags [IOStreams.In] as a TTY.
func (s *IOStreams) SetStdinTTY(isTTY bool) {
	s.stdinTTYOverride = true
	s.stdinIsTTY = isTTY
}

// IsStdinTTY returns true if [IOStreams.In] is a TTY.
func (s *IOStreams) IsStdinTTY() bool {
	if s.stdinTTYOverride {
		return s.stdinIsTTY
	}
	return IsTerminal(s.In)
}

// SetStdoutTTY explicitly flags [IOStreams.Out] as a TTY.
func (s *IOStreams) SetStdoutTTY(isTTY bool) {
	s.stdoutTTYOverride = true
	s.stdoutIsTTY = isTTY
}

// IsStdoutTTY returns true if [IOStreams.Out] is a TTY.
func (s *IOStreams) IsStdoutTTY() bool {
	if s.stdoutTTYOverride {
		return s.stdoutIsTTY
	}
	return IsTerminal(s.Out)
}

// SetStderrTTY explicitly flags [IOStreams.Err] as a TTY.
func (s *IOStreams) SetStderrTTY(isTTY bool) {
	s.stderrTTYOverride = true
	s.stderrIsTTY = isTTY
}

// IsStderrTTY returns true if [IOStreams.Err] is a TTY.
func (s *IOStreams) IsStderrTTY() bool {
	if s.stderrTTYOverride {
		return s.stderrIsTTY
	}
	return IsTerminal(s.Err)
}

// IsInteractive returns true if running interactively.
// Will be false if either (a) std in/out is not a TTY,
// or (b) the user has explicitly requested not to be prompted.
func (s *IOStreams) IsInteractive() bool {
	if s.isInteractiveOverride {
		return s.isInteractive
	}
	return s.IsStdinTTY() && s.IsStdoutTTY()
}

// SetInteractive explicitly sets whether this is an interactive session.
func (s *IOStreams) SetInteractive(v bool) {
	s.isInteractiveOverride = true
	s.isInteractive = v
}

// ProgressIndicatorEnabled returns true if the progress indicator is enabled.
func (s *IOStreams) ProgressIndicatorEnabled() bool {
	return s.progressIndicatorEnabled
}

// SetProgressIndicatorEnabled sets whether the progress indicator is enabled.
func (s *IOStreams) SetProgressIndicatorEnabled(v bool) {
	s.progressIndicatorEnabled = v
}

func (s *IOStreams) StartProgressIndicator() {
	s.StartProgressIndicatorWithLabel("")
}

func (s *IOStreams) StartProgressIndicatorWithLabel(label string) {
	if !s.progressIndicatorEnabled {
		return
	}

	s.progressIndicatorMu.Lock()
	defer s.progressIndicatorMu.Unlock()

	if s.progressIndicator != nil {
		if label == "" {
			s.progressIndicator.Prefix = ""
		} else {
			s.progressIndicator.Prefix = label + " "
		}
		return
	}

	// https://github.com/briandowns/spinner#available-character-sets
	dotStyle := spinner.CharSets[11]
	sp := spinner.New(dotStyle, 120*time.Millisecond, spinner.WithWriter(s.Err), spinner.WithColor("fgCyan"))
	if label != "" {
		sp.Prefix = label + " "
	}

	sp.Start()
	s.progressIndicator = sp
}

func (s *IOStreams) StopProgressIndicator() {
	s.progressIndicatorMu.Lock()
	defer s.progressIndicatorMu.Unlock()
	if s.progressIndicator == nil {
		return
	}
	s.progressIndicator.Stop()
	s.progressIndicator = nil
}

// IsTerminal returns true if the stream is a terminal.
func IsTerminal(stream IOStream) bool {
	return isatty.IsTerminal(stream.Fd()) || isatty.IsCygwinTerminal(stream.Fd())
}

// Returns an IOStreams containing os.Stdin, os.Stdout, and os.Stderr.
func System() *IOStreams {
	ios := &IOStreams{
		In:  &systemIOStream{File: os.Stdin},
		Out: &systemIOStream{File: os.Stdout},
		Err: &systemIOStream{File: os.Stderr},
	}
	stdoutIsTTY := ios.IsStdoutTTY()
	stderrIsTTY := ios.IsStderrTTY()
	ios.SetColorEnabled(EnvColorForced() || (stdoutIsTTY && !EnvColorDisabled()))
	ios.SetProgressIndicatorEnabled(stdoutIsTTY && stderrIsTTY)
	ios.SetStdoutTTY(stdoutIsTTY)
	ios.SetStderrTTY(stdoutIsTTY)
	return ios
}

// Returns an IOStreams with mock in/out/err values.
func Test() *IOStreams {
	ios := &IOStreams{
		In:  &mockIOStream{Buffer: &bytes.Buffer{}, fd: 0},
		Out: &mockIOStream{Buffer: &bytes.Buffer{}, fd: 1},
		Err: &mockIOStream{Buffer: &bytes.Buffer{}, fd: 2},
	}
	ios.SetColorEnabled(false)
	ios.SetProgressIndicatorEnabled(false)
	ios.SetStdinTTY(false)
	ios.SetStdoutTTY(false)
	ios.SetStderrTTY(false)
	return ios
}

var (
	_ IOStream = &systemIOStream{}
	_ IOStream = &mockIOStream{}
)

// Wrapper so we can make os.Stdin and friends fulfill IOStream.
type systemIOStream struct {
	*os.File
}

func (f *systemIOStream) String() string {
	buf, _ := io.ReadAll(f)
	return string(buf)
}

// Wrapper so we can make bytes.Buffer fulfill IOStream.
type mockIOStream struct {
	*bytes.Buffer
	fd uintptr
}

func (m *mockIOStream) Fd() uintptr {
	return m.fd
}

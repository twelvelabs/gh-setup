package iostreams

/*
This file started out as a copy of https://github.com/cli/cli/blob/trunk/pkg/iostreams/color.go
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
	"fmt"
	"strings"

	"github.com/mgutz/ansi" //cspell: disable-line
)

var (
	magenta = ansi.ColorFunc("magenta")
	cyan    = ansi.ColorFunc("cyan")
	red     = ansi.ColorFunc("red")
	yellow  = ansi.ColorFunc("yellow")
	blue    = ansi.ColorFunc("blue")
	green   = ansi.ColorFunc("green")
	gray    = ansi.ColorFunc("black+h")
	bold    = ansi.ColorFunc("default+b")
)

func NewFormatter(enabled bool) *Formatter {
	return &Formatter{
		enabled: enabled,
	}
}

// Formatter wraps strings in ANSI escape sequences.
type Formatter struct {
	enabled bool
}

func (f *Formatter) Bold(t string) string {
	if !f.enabled {
		return t
	}
	return bold(t)
}

func (f *Formatter) Boldf(t string, args ...interface{}) string {
	return f.Bold(fmt.Sprintf(t, args...))
}

func (f *Formatter) Red(t string) string {
	if !f.enabled {
		return t
	}
	return red(t)
}

func (f *Formatter) Redf(t string, args ...interface{}) string {
	return f.Red(fmt.Sprintf(t, args...))
}

func (f *Formatter) Yellow(t string) string {
	if !f.enabled {
		return t
	}
	return yellow(t)
}

func (f *Formatter) Yellowf(t string, args ...interface{}) string {
	return f.Yellow(fmt.Sprintf(t, args...))
}

func (f *Formatter) Green(t string) string {
	if !f.enabled {
		return t
	}
	return green(t)
}

func (f *Formatter) Greenf(t string, args ...interface{}) string {
	return f.Green(fmt.Sprintf(t, args...))
}

func (f *Formatter) Gray(t string) string {
	if !f.enabled {
		return t
	}
	return gray(t)
}

func (f *Formatter) Grayf(t string, args ...interface{}) string {
	return f.Gray(fmt.Sprintf(t, args...))
}

func (f *Formatter) Magenta(t string) string {
	if !f.enabled {
		return t
	}
	return magenta(t)
}

func (f *Formatter) Magentaf(t string, args ...interface{}) string {
	return f.Magenta(fmt.Sprintf(t, args...))
}

func (f *Formatter) Cyan(t string) string {
	if !f.enabled {
		return t
	}
	return cyan(t)
}

func (f *Formatter) Cyanf(t string, args ...interface{}) string {
	return f.Cyan(fmt.Sprintf(t, args...))
}

func (f *Formatter) Blue(t string) string {
	if !f.enabled {
		return t
	}
	return blue(t)
}

func (f *Formatter) Bluef(t string, args ...interface{}) string {
	return f.Blue(fmt.Sprintf(t, args...))
}

func (f *Formatter) Info(t string) string {
	return f.Gray(t)
}

func (f *Formatter) InfoIcon() string {
	return f.Info("•")
}

func (f *Formatter) Success(t string) string {
	return f.Green(t)
}

func (f *Formatter) SuccessIcon() string {
	return f.Success("✓")
}

func (f *Formatter) Warning(t string) string {
	return f.Yellow(t)
}

func (f *Formatter) WarningIcon() string {
	return f.Warning("!")
}

func (f *Formatter) Failure(t string) string {
	return f.Red(t)
}

func (f *Formatter) FailureIcon() string {
	return f.Failure("✖")
}

type ColorFunc func(string) string

func (f *Formatter) ColorFromString(s string) ColorFunc {
	var fn ColorFunc

	switch strings.ToLower(s) {
	case "bold":
		fn = f.Bold
	case "red":
		fn = f.Red
	case "yellow":
		fn = f.Yellow
	case "green":
		fn = f.Green
	case "gray":
		fn = f.Gray
	case "magenta":
		fn = f.Magenta
	case "cyan":
		fn = f.Cyan
	case "blue":
		fn = f.Blue
	default:
		fn = func(s string) string {
			return s
		}
	}

	return fn
}

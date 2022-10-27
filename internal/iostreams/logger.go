package iostreams

import (
	"fmt"
	"strings"
)

const (
	ActionWidth = 10
)

// NewActionLogger returns a new ActionLogger.
func NewActionLogger(ios *IOStreams, dryRun bool) *ActionLogger {
	return &ActionLogger{
		ios:    ios,
		format: ios.Formatter(),
		dryRun: dryRun,
	}
}

// ActionLogger logs formatted Task actions.
type ActionLogger struct {
	ios    *IOStreams
	format *Formatter
	dryRun bool
}

// Info logs a line to os.StdErr prefixed with an info icon and action label.
//
// Example:
//
//	// Prints "• [    action]: hello, world\n"
//	Info("action", "hello, %s", "world")
func (l *ActionLogger) Info(action string, line string, args ...any) {
	icon := l.format.InfoIcon()
	action = l.format.Info(l.rightJustify(action))
	l.log(icon, action, line, args...)
}

// Success logs a line to os.StdErr prefixed with a success icon and action label.
//
// Example:
//
//	// Prints "✓ [    action]: hello, world\n"
//	Success("action", "hello, %s", "world")
func (l *ActionLogger) Success(action string, line string, args ...any) {
	icon := l.format.SuccessIcon()
	action = l.format.Success(l.rightJustify(action))
	l.log(icon, action, line, args...)
}

// Warning logs a line to os.StdErr prefixed with a warning icon and action label.
//
// Example:
//
//	// Prints "! [    action]: hello, world\n"
//	Warning("action", "hello, %s", "world")
func (l *ActionLogger) Warning(action string, line string, args ...any) {
	icon := l.format.WarningIcon()
	action = l.format.Warning(l.rightJustify(action))
	l.log(icon, action, line, args...)
}

// Failure logs a line to os.StdErr prefixed with a failure icon and action label.
//
// Example:
//
//	// Prints "✖ [    action]: hello, world\n"
//	Failure("action", "hello, %s", "world")
func (l *ActionLogger) Failure(action string, line string, args ...any) {
	icon := l.format.FailureIcon()
	action = l.format.Failure(l.rightJustify(action))
	l.log(icon, action, line, args...)
}

// logs the formatted icon, action, and line to StdErr.
func (l *ActionLogger) log(icon, action, line string, args ...any) {
	prefix := icon + " "
	if l.dryRun {
		prefix += "[DRY RUN]"
	}
	line = l.ensureNewline(prefix + "[" + action + "]: " + line)
	fmt.Fprintf(l.ios.Err, line, args...)
}

func (l *ActionLogger) ensureNewline(text string) string {
	if !strings.HasSuffix(text, "\n") {
		text += "\n"
	}
	return text
}

func (l *ActionLogger) rightJustify(text string) string {
	return fmt.Sprintf("%*s", ActionWidth, text)
}

// NewIconLogger returns a new IconLogger.
func NewIconLogger(ios *IOStreams) *IconLogger {
	return &IconLogger{
		ios:    ios,
		format: ios.Formatter(),
	}
}

// IconLogger is a generic logger that prefixes lines with status icons.
type IconLogger struct {
	ios    *IOStreams
	format *Formatter
}

func (l *IconLogger) Info(line string, args ...any) {
	icon := l.format.InfoIcon()
	fmt.Fprintf(l.ios.Err, (icon + " " + line), args...)
}

func (l *IconLogger) Success(line string, args ...any) {
	icon := l.format.SuccessIcon()
	fmt.Fprintf(l.ios.Err, (icon + " " + line), args...)
}

func (l *IconLogger) Warning(line string, args ...any) {
	icon := l.format.WarningIcon()
	fmt.Fprintf(l.ios.Err, (icon + " " + line), args...)
}

func (l *IconLogger) Failure(line string, args ...any) {
	icon := l.format.FailureIcon()
	fmt.Fprintf(l.ios.Err, (icon + " " + line), args...)
}

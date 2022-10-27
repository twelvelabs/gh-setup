package iostreams

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestActionLogger_Info(t *testing.T) {
	tests := []struct {
		name     string
		dryRun   bool
		action   string
		line     string
		args     []any
		expected string
	}{
		{
			name:     "logs to stderr with an info icon",
			action:   "test",
			line:     "hello",
			args:     []any{},
			expected: "• [      test]: hello\n",
		},
		{
			name:     "handles printf strings and args",
			action:   "test",
			line:     "hello %s",
			args:     []any{"world"},
			expected: "• [      test]: hello world\n",
		},
		{
			name:     "only adds trailing newline if needed",
			action:   "test",
			line:     "hello\n",
			args:     []any{},
			expected: "• [      test]: hello\n",
		},
		{
			name:     "adds dry run prefix",
			dryRun:   true,
			action:   "test",
			line:     "hello\n",
			args:     []any{},
			expected: "• [DRY RUN][      test]: hello\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ios := Test()

			logger := NewActionLogger(ios, tt.dryRun)
			logger.Info(tt.action, tt.line, tt.args...)

			assert.Equal(t, tt.expected, ios.Err.String())
		})
	}
}

func TestActionLogger_Success(t *testing.T) {
	tests := []struct {
		name     string
		dryRun   bool
		action   string
		line     string
		args     []any
		expected string
	}{
		{
			name:     "logs to stderr with an info icon",
			action:   "test",
			line:     "hello",
			args:     []any{},
			expected: "✓ [      test]: hello\n",
		},
		{
			name:     "handles printf strings and args",
			action:   "test",
			line:     "hello %s",
			args:     []any{"world"},
			expected: "✓ [      test]: hello world\n",
		},
		{
			name:     "only adds trailing newline if needed",
			action:   "test",
			line:     "hello\n",
			args:     []any{},
			expected: "✓ [      test]: hello\n",
		},
		{
			name:     "adds dry run prefix",
			dryRun:   true,
			action:   "test",
			line:     "hello\n",
			args:     []any{},
			expected: "✓ [DRY RUN][      test]: hello\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ios := Test()

			logger := NewActionLogger(ios, tt.dryRun)
			logger.Success(tt.action, tt.line, tt.args...)

			assert.Equal(t, tt.expected, ios.Err.String())
		})
	}
}

func TestActionLogger_Warning(t *testing.T) {
	tests := []struct {
		name     string
		dryRun   bool
		action   string
		line     string
		args     []any
		expected string
	}{
		{
			name:     "logs to stderr with an info icon",
			action:   "test",
			line:     "hello",
			args:     []any{},
			expected: "! [      test]: hello\n",
		},
		{
			name:     "handles printf strings and args",
			action:   "test",
			line:     "hello %s",
			args:     []any{"world"},
			expected: "! [      test]: hello world\n",
		},
		{
			name:     "only adds trailing newline if needed",
			action:   "test",
			line:     "hello\n",
			args:     []any{},
			expected: "! [      test]: hello\n",
		},
		{
			name:     "adds dry run prefix",
			dryRun:   true,
			action:   "test",
			line:     "hello\n",
			args:     []any{},
			expected: "! [DRY RUN][      test]: hello\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ios := Test()

			logger := NewActionLogger(ios, tt.dryRun)
			logger.Warning(tt.action, tt.line, tt.args...)

			assert.Equal(t, tt.expected, ios.Err.String())
		})
	}
}

func TestActionLogger_Failure(t *testing.T) {
	tests := []struct {
		name     string
		dryRun   bool
		action   string
		line     string
		args     []any
		expected string
	}{
		{
			name:     "logs to stderr with an info icon",
			action:   "test",
			line:     "hello",
			args:     []any{},
			expected: "✖ [      test]: hello\n",
		},
		{
			name:     "handles printf strings and args",
			action:   "test",
			line:     "hello %s",
			args:     []any{"world"},
			expected: "✖ [      test]: hello world\n",
		},
		{
			name:     "only adds trailing newline if needed",
			action:   "test",
			line:     "hello\n",
			args:     []any{},
			expected: "✖ [      test]: hello\n",
		},
		{
			name:     "adds dry run prefix",
			dryRun:   true,
			action:   "test",
			line:     "hello\n",
			args:     []any{},
			expected: "✖ [DRY RUN][      test]: hello\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ios := Test()

			logger := NewActionLogger(ios, tt.dryRun)
			logger.Failure(tt.action, tt.line, tt.args...)

			assert.Equal(t, tt.expected, ios.Err.String())
		})
	}
}

func TestIconLogger_Info(t *testing.T) {
	tests := []struct {
		name     string
		line     string
		args     []any
		expected string
	}{
		{
			name:     "logs to stderr with an info icon",
			line:     "hello",
			args:     []any{},
			expected: "• hello",
		},
		{
			name:     "handles printf strings and args",
			line:     "hello %s",
			args:     []any{"world"},
			expected: "• hello world",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ios := Test()

			logger := NewIconLogger(ios)
			logger.Info(tt.line, tt.args...)

			assert.Equal(t, tt.expected, ios.Err.String())
		})
	}
}

func TestIconLogger_Success(t *testing.T) {
	tests := []struct {
		name     string
		line     string
		args     []any
		expected string
	}{
		{
			name:     "logs to stderr with a success icon",
			line:     "hello",
			args:     []any{},
			expected: "✓ hello",
		},
		{
			name:     "handles printf strings and args",
			line:     "hello %s",
			args:     []any{"world"},
			expected: "✓ hello world",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ios := Test()

			logger := NewIconLogger(ios)
			logger.Success(tt.line, tt.args...)

			assert.Equal(t, tt.expected, ios.Err.String())
		})
	}
}

func TestIconLogger_Warning(t *testing.T) {
	tests := []struct {
		name     string
		line     string
		args     []any
		expected string
	}{
		{
			name:     "logs to stderr with a warning icon",
			line:     "hello",
			args:     []any{},
			expected: "! hello",
		},
		{
			name:     "handles printf strings and args",
			line:     "hello %s",
			args:     []any{"world"},
			expected: "! hello world",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ios := Test()

			logger := NewIconLogger(ios)
			logger.Warning(tt.line, tt.args...)

			assert.Equal(t, tt.expected, ios.Err.String())
		})
	}
}

func TestIconLogger_Failure(t *testing.T) {
	tests := []struct {
		name     string
		line     string
		args     []any
		expected string
	}{
		{
			name:     "logs to stderr with a failure icon",
			line:     "hello",
			args:     []any{},
			expected: "✖ hello",
		},
		{
			name:     "handles printf strings and args",
			line:     "hello %s",
			args:     []any{"world"},
			expected: "✖ hello world",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ios := Test()

			logger := NewIconLogger(ios)
			logger.Failure(tt.line, tt.args...)

			assert.Equal(t, tt.expected, ios.Err.String())
		})
	}
}

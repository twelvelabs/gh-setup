package prompt

//go:generate moq -rm -out prompter_mock.go . Prompter
type Prompter interface {
	// Prompt for a boolean yes/no value.
	Confirm(msg string, value bool, help string) (bool, error)
	// Prompt for single string value.
	Input(msg string, value string, help string) (string, error)
	// Prompt for a slice of string values w/ a fixed set of options.
	MultiSelect(msg string, options []string, values []string, help string) ([]string, error)
	// Prompt for single string value w/ a fixed set of options.
	Select(msg string, options []string, value string, help string) (string, error)
}

type ConfirmFunc func(msg string, value bool, help string) (bool, error)
type InputFunc func(msg string, value string, help string) (string, error)
type MultiSelectFunc func(msg string, options []string, values []string, help string) ([]string, error)
type SelectFunc func(msg string, options []string, value string, help string) (string, error)

// Creates a new ConfirmFunc that returns the given result and err.
func NewConfirmFunc(result bool, err error) ConfirmFunc {
	return func(msg string, value bool, help string) (bool, error) {
		return result, err
	}
}

// Creates a new ConfirmFunc that delegates to funcs in series.
func NewConfirmFuncSet(funcs ...ConfirmFunc) ConfirmFunc {
	var f ConfirmFunc
	return func(msg string, value bool, help string) (bool, error) {
		f, funcs = funcs[0], funcs[1:]
		return f(msg, value, help)
	}
}

// Creates a new ConfirmFunc that returns the default value.
func NewNoopConfirmFunc() ConfirmFunc {
	return func(msg string, value bool, help string) (bool, error) {
		return value, nil
	}
}

// Creates a new InputFunc that returns the given result and err.
func NewInputFunc(result string, err error) InputFunc {
	return func(msg string, value string, help string) (string, error) {
		return result, err
	}
}

// Creates a new InputFunc that returns the default value.
func NewNoopInputFunc() InputFunc {
	return func(msg string, value string, help string) (string, error) {
		return value, nil
	}
}

// Creates a new SelectFunc that returns the given result and err.
func NewMultiSelectFunc(result []string, err error) MultiSelectFunc {
	return func(msg string, options []string, values []string, help string) ([]string, error) {
		return result, err
	}
}

// Creates a new InputFunc that returns the default value.
func NewNoopMultiSelectFunc() MultiSelectFunc {
	return func(msg string, options []string, values []string, help string) ([]string, error) {
		return values, nil
	}
}

// Creates a new SelectFunc that returns the given result and err.
func NewSelectFunc(result string, err error) SelectFunc {
	return func(msg string, options []string, value string, help string) (string, error) {
		return result, err
	}
}

// Creates a new InputFunc that returns the default value.
func NewNoopSelectFunc() SelectFunc {
	return func(msg string, options []string, value string, help string) (string, error) {
		return value, nil
	}
}

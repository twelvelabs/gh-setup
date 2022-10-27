package testutil

import (
	"os"
	"testing"
)

// InTempDir runs handler inside a temp dir, then returns back into the cwd.
func InTempDir(tb testing.TB, handler func(tmpDir string)) {
	tb.Helper()

	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	tmp := tb.TempDir()
	if err := os.Chdir(tmp); err != nil {
		panic(err)
	}
	handler(tmp)
	if err := os.Chdir(cwd); err != nil {
		panic(err)
	}
}

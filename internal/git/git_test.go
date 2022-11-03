package git

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/twelvelabs/termite/testutil"
)

func TestHasCommits(t *testing.T) {
	testutil.InTempDir(t, func(tmpDir string) {
		_, _, err := Exec("init")
		assert.NoError(t, err)
		assert.Equal(t, false, HasCommits())

		err = os.WriteFile("foo.txt", []byte("aaa"), 0600)
		assert.NoError(t, err)
		_, _, err = Exec("add", "foo.txt")
		assert.NoError(t, err)
		_, _, err = Exec("commit", "--no-gpg-sign", "--no-verify", "-m", "add foo")
		assert.NoError(t, err)

		assert.Equal(t, true, HasCommits())
	})
}

func TestHasRemote(t *testing.T) {
	testutil.InTempDir(t, func(tmpDir string) {
		_, _, err := Exec("init")
		assert.NoError(t, err)

		assert.Equal(t, false, HasRemote("origin"))

		_, _, err = Exec("remote", "add", "origin", "git@github.com:cli/cli.git")
		assert.NoError(t, err)

		assert.Equal(t, true, HasRemote("origin"))
		assert.Equal(t, false, HasRemote("unknown"))
	})
}

func TestExec(t *testing.T) {
	stdout, stderr, err := Exec("--version")
	assert.NoError(t, err)
	assert.Equal(t, "", stderr.String())
	assert.Contains(t, stdout.String(), "git version")
}

func TestIsInitialized(t *testing.T) {
	testutil.InTempDir(t, func(tmpDir string) {
		assert.False(t, IsInitialized())

		_, _, err := Exec("init")
		assert.NoError(t, err)

		assert.True(t, IsInitialized())
	})
}

func TestIsInstalled(t *testing.T) {
	// not going to bother w/ testing the inverse :shrug:
	assert.True(t, IsInstalled())
}

func TestStatusLinesAndIsDirty(t *testing.T) {
	testutil.InTempDir(t, func(tmpDir string) {
		_, _, err := Exec("init")
		assert.NoError(t, err)

		assert.Equal(t, false, IsDirty())
		lines, err := StatusLines()
		assert.NoError(t, err)
		assert.Equal(t, 0, len(lines))

		err = os.WriteFile("foo.txt", []byte("aaa"), 0600)
		assert.NoError(t, err)

		assert.Equal(t, true, IsDirty())
		lines, err = StatusLines()
		assert.NoError(t, err)
		assert.Equal(t, 1, len(lines))
		assert.Equal(t, "?? foo.txt", lines[0])

		_, _, err = Exec("add", "foo.txt")
		assert.NoError(t, err)

		assert.Equal(t, true, IsDirty())
		lines, err = StatusLines()
		assert.NoError(t, err)
		assert.Equal(t, 1, len(lines))
		assert.Equal(t, "A  foo.txt", lines[0])

		_, _, err = Exec("commit", "--no-gpg-sign", "--no-verify", "-m", "add foo")
		assert.NoError(t, err)

		assert.Equal(t, false, IsDirty())
		lines, err = StatusLines()
		assert.NoError(t, err)
		assert.Equal(t, 0, len(lines))

		err = os.WriteFile("foo.txt", []byte("bbb"), 0600)
		assert.NoError(t, err)
		err = os.WriteFile("bar.txt", []byte("bbb"), 0600)
		assert.NoError(t, err)

		assert.Equal(t, true, IsDirty())
		lines, err = StatusLines()
		assert.NoError(t, err)
		assert.Equal(t, 2, len(lines))
		assert.Equal(t, " M foo.txt", lines[0])
		assert.Equal(t, "?? bar.txt", lines[1])

		_, _, err = Exec("add", "foo.txt")
		assert.NoError(t, err)

		assert.Equal(t, true, IsDirty())
		lines, err = StatusLines()
		assert.NoError(t, err)
		assert.Equal(t, 2, len(lines))
		assert.Equal(t, "M  foo.txt", lines[0])
		assert.Equal(t, "?? bar.txt", lines[1])

		_, _, err = Exec("add", "bar.txt")
		assert.NoError(t, err)
		_, _, err = Exec("commit", "--no-gpg-sign", "--no-verify", "-m", "update foo, add bar")
		assert.NoError(t, err)

		assert.Equal(t, false, IsDirty())
		lines, err = StatusLines()
		assert.NoError(t, err)
		assert.Equal(t, 0, len(lines))
	})
}

package filesystem_test

import (
	"afero-objstor/filesystem"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	fs, cleanup := filesystem.NewTestFsOrFail(t)
	defer cleanup()
	assert.NotNil(t, fs)
}

package filesystem_test

import (
	"afero-objstor/filesystem"
	"github.com/Pallinder/go-randomdata"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestObjStorFs_CreateInRoot(t *testing.T) {
	fs, cleanup := filesystem.NewTestFsOrFail(t)
	defer cleanup()

	filename := randomdata.RandStringRunes(10)
	filePath := "/" + filename
	f, err := fs.Create(filePath)
	assert.Nil(t, err)
	assert.Equal(t, filename, f.Name())
}

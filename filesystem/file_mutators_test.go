package filesystem_test

import (
	"github.com/spf13/afero"
	"testing"

	"github.com/Pallinder/go-randomdata"
	"github.com/stretchr/testify/assert"

	"afero-objstor/filesystem"
)

func TestProjectedObject_Write(t *testing.T) {
	fs, cleanup := filesystem.NewTestFsOrFail(t)
	defer cleanup()

	fileName := "/" + randomdata.RandStringRunes(10)
	f, err := fs.Create(fileName)
	assert.Nil(t, err)

	fileContents := []byte(randomdata.RandStringRunes(10))
	amountWritten, writeErr := f.Write(fileContents)
	assert.Nil(t, writeErr, "error while writing to file")
	assert.Equal(t, 10, amountWritten,
		"did not report full amount of data written")

	syncErr := f.Sync()
	assert.Nil(t, syncErr, "could not sync file")

	f.Seek(0, 0)

	fileBytes, readAllErr := afero.ReadAll(f)
	assert.Nil(t, readAllErr, "could not read back file contents")

	assert.Equal(t, fileContents, fileBytes,
		"file did not have expected contnets")
}

package filesystem_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/google/go-cloud/blob/fileblob"
	"github.com/stretchr/testify/assert"

	"afero-objstor/filesystem"
)

func TestNew(t *testing.T) {
	tmpDir, err := ioutil.TempDir("", "test")
	defer os.RemoveAll(tmpDir)
	assert.Nil(t, err)

	fsBucket, err := fileblob.NewBucket(tmpDir)
	assert.Nil(t, err)

	fs, err := filesystem.NewFs(filesystem.WithBucket(fsBucket))
	assert.Nil(t, err)
	assert.NotNil(t, fs)
}

package filesystem

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/google/go-cloud/blob/fileblob"
	"github.com/stretchr/testify/assert"
)

// NewTestFsOrFail will attempt to create a filesystem backed by a bucket stored
// in a temporary directory on the local filesystem. The function returned will
// delete all files backing the cloud bucket.
func NewTestFsOrFail(t *testing.T) (*ObjStorFs, func()) {
	tmpDir, err := ioutil.TempDir("", "test")
	cleanupFunc := func() { os.RemoveAll(tmpDir) }
	assert.Nil(t, err)

	fsBucket, err := fileblob.NewBucket(tmpDir)
	assert.Nil(t, err)

	fs, err := NewFs(WithBucket(fsBucket))
	assert.Nil(t, err)

	return fs, cleanupFunc
}

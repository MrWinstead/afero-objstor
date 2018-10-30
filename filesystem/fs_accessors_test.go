package filesystem_test

import (
	"afero-objstor/filesystem"
	"context"
	"github.com/Pallinder/go-randomdata"
	"github.com/google/go-cloud/blob"
	"github.com/google/go-cloud/blob/fileblob"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func TestFsName(t *testing.T) {
	fs, cleanup := filesystem.NewTestFsOrFail(t)
	defer cleanup()

	n := fs.Name()
	assert.NotEqual(t, "", n)
}

func TestObjStorFs_OpenNotExist(t *testing.T) {
	fs, cleanup := filesystem.NewTestFsOrFail(t)
	defer cleanup()

	filePath := "/" + randomdata.RandStringRunes(10)
	_, err := fs.Open(filePath)
	assert.Equal(t, os.ErrNotExist, err)
}

func TestObjStorFs_Open(t *testing.T) {
	fs, cleanup := filesystem.NewTestFsOrFail(t)
	defer cleanup()

	filePath := "/" + randomdata.RandStringRunes(10)
	_, err := fs.Open(filePath)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "does not exist")
}

func TestObjStorFs_OpenNotInLocal(t *testing.T) {
	tmpDir, err := ioutil.TempDir("", "test")
	assert.Nil(t, err)
	defer os.RemoveAll(tmpDir)

	bucket, err := fileblob.NewBucket(tmpDir)
	assert.Nil(t, err)

	fileContents := []byte(randomdata.RandStringRunes(10))
	filename := "/" + randomdata.RandStringRunes(10)
	writerOpts := blob.WriterOptions{
		ContentType: "application/octet-stream",
	}
	writer, err := bucket.NewWriter(context.Background(), filename[1:],
		&writerOpts)
	assert.Nil(t, err)
	writer.Write(fileContents)
	writer.Close()

	fs, err := filesystem.NewFs(filesystem.WithBucket(bucket))
	assert.Nil(t, err)

	f, err := fs.Open(filename)
	assert.Nil(t, err)
	assert.NotNil(t, f)

	err = f.Sync()
	assert.Nil(t, err)
}

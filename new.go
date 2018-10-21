package afero_objstor

import (
	"afero-objstor/filesystem"
	"github.com/spf13/afero"
	"github.com/google/go-cloud/blob"
)

func New(bucket *blob.Bucket, options... filesystem.FsConstructorOption) (afero.Fs, error) {

	bucketOption := filesystem.WithBucket(bucket)
	inMemoryBackingFs := filesystem.WithBackingFilesystem(afero.NewMemMapFs())

	allOptions := append(
		[]filesystem.FsConstructorOption{bucketOption, inMemoryBackingFs},
		options...)

	created, createErr := filesystem.NewFs(allOptions...)
	return created, createErr
}

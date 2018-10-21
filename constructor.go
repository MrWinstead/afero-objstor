package objstorfs

import (
	"afero-objstor/filesystem"
	"github.com/google/go-cloud/blob"
	"github.com/spf13/afero"
)

// New constructs a filesystem while exposing the afero.Fs interface
func New(bucket *blob.Bucket, options ...filesystem.ConstructorOption) (afero.Fs, error) {
	return NewDeadlined(bucket, options...)
}

// NewDeadlined construct a filesytem while exposing the enhanced DeadlineFs
// interface
func NewDeadlined(bucket *blob.Bucket, options ...filesystem.ConstructorOption,
) (filesystem.DeadlineFs, error) {
	allOptions := append(options, filesystem.WithBucket(bucket))
	created, createErr := filesystem.NewFs(allOptions...)
	return created, createErr
}

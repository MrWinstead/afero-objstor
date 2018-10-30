package filesystem

import (
	"context"
	"github.com/google/go-cloud/blob"
	"github.com/spf13/afero"
)

const (
	fileTypeFile = iota
	fileTypeDirectory
	stateClosed
	stateOpenedForRead
	stateOpenedForWrite
)

type contextGenerator func(operationName, operationClass string) (context.Context, context.CancelFunc)

type options struct {
	WriteThrough bool
}

// ProjectedObject is a set of state describing an object in object storage
// projected into the local filesystem cache
type ProjectedObject struct {
	localInstance afero.File

	fileType int

	bucket *blob.Bucket
	fs     *ObjStorFs

	ctxGen contextGenerator
	opts   options
}

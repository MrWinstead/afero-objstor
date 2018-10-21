package filesystem

import (
	"context"
	"os"

	"github.com/google/go-cloud/blob"
)

const (
	fileTypeFile = iota
	fileTypeDirectory
)

type contextGenerator func(operationName, operationClass string) (context.Context, context.CancelFunc)

type options struct {
	WriteThrough bool
}

type fileinfo struct {
	obj *ProjectedObject
}

// ProjectedObject is a set of state describing an object in object storage
// projected into the local filesystem cache
type ProjectedObject struct {
	size            int64
	mode            os.FileMode
	currentPosition int64

	name     string
	fullKey  string
	fileType int

	bucket *blob.Bucket
	fs     *ObjStorFs

	ctxGen contextGenerator
	opts   options
}

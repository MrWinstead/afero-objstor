package filesystem

import (
	"context"
	"os"

	"github.com/google/go-cloud/blob"
)

type ContextGenerator func(operationName, operationClass string) (context.Context, context.CancelFunc)

type options struct {
	WriteThrough bool
}

type fileinfo struct {
	obj *ProjectedObject
}

type ProjectedObject struct {
	size int64
	mode os.FileMode
	currentPosition int64

	name string
	fullKey string

	bucket *blob.Bucket
	fs     *ObjStorFs

	ctxGen ContextGenerator
	opts        options
}

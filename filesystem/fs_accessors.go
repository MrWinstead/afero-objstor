package filesystem

import (
	"context"
	"github.com/google/go-cloud/blob"
	"github.com/spf13/afero"
	"os"
)

func (fs *ObjStorFs) GetKeyPrefix() string {
	return fs.keyPrefix
}

func (fs *ObjStorFs) GetBucket() *blob.Bucket {
	return fs.bucket
}

func (fs *ObjStorFs) Name() string {
	return "object-storage"
}

func (fs *ObjStorFs) Open(name string) (afero.File, error) {
	ctx, _ := fs.getOperationContext("Open", DeadlineKeyRead)
	return fs.OpenEx(ctx, name)
}

func (fs *ObjStorFs) OpenEx(ctx context.Context, name string) (*ProjectedObject, error) {
	return nil, nil
}

func (fs *ObjStorFs) OpenFile(name string, flag int, perm os.FileMode) (afero.File, error) {
	ctx, _ := fs.getOperationContext("OpenFile", DeadlineKeyRead)
	return fs.OpenFileEx(ctx, name, flag, perm)
}

func (fs *ObjStorFs) OpenFileEx(ctx context.Context, name string, flag int,
	perm os.FileMode) (*ProjectedObject, error) {
	return nil, nil
}

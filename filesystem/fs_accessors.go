package filesystem

import (
	"context"
	"github.com/spf13/afero"
	"os"
)

// Name fetches the name of the filesystem driver
func (fs *ObjStorFs) Name() string {
	return "object-storage"
}

// Open will create a local instance of the named file. Nothing is read or
// written to object storage until other functions are called. If the file does
// not exist os.ErrNotExist is returned
func (fs *ObjStorFs) Open(name string) (afero.File, error) {
	ctx, _ := fs.getOperationContext("Open", deadlineKeyRead)
	return fs.OpenEx(ctx, name)
}

// OpenEx behaves like Open, but also allows passing a context with possible
// deadline to lower layers
func (fs *ObjStorFs) OpenEx(ctx context.Context, name string) (DeadlineFile, error) {
	return nil, nil
}

// OpenFile will open a file in the local file cache. If O_CREATE is specified
// in flags, the file will be touched to object storage. If the parent directory
// does not exist, os.ErrNotExist will be returned
func (fs *ObjStorFs) OpenFile(name string, flag int, perm os.FileMode) (afero.File, error) {
	ctx, _ := fs.getOperationContext("OpenFile", deadlineKeyRead)
	return fs.OpenFileEx(ctx, name, flag, perm)
}

// OpenFileEx behaves like OpenFile, but also allows passing a context with possible
// deadline to lower layers
func (fs *ObjStorFs) OpenFileEx(ctx context.Context, name string, flag int,
	perm os.FileMode) (DeadlineFile, error) {
	return nil, nil
}

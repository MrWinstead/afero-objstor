package filesystem

import (
	"context"
	"os"
	"time"
)

// Chmod will set the file mode of the specified file. If the file does not
// exist, it will return os.ErrNotExist. If the permission is invalid, it will
// return os.ErrPermission
func (fs *ObjStorFs) Chmod(name string, mode os.FileMode) error {
	ctx, _ := fs.getOperationContext("Chmod", deadlineKeyMetadataWrite)
	return fs.ChmodEx(ctx, name, mode)
}

// ChmodEx behaves like Chmod, but also allows passing a context with possible
// deadline to lower layers
func (fs *ObjStorFs) ChmodEx(ctx context.Context, name string, mode os.FileMode) error {
	return nil
}

// Chtimes will set the access (atime) and modified (mtime) timestamps of the
// specified file. If the file does not exist, it will return os.ErrNotExist
func (fs *ObjStorFs) Chtimes(name string, atime, mtime time.Time) error {
	ctx, _ := fs.getOperationContext("Chtimes", deadlineKeyMetadataWrite)
	return fs.ChtimesEx(ctx, name, atime, mtime)
}

// ChtimesEx behaves like Chtimes, but also allows passing a context with possible
// deadline to lower layers
func (fs *ObjStorFs) ChtimesEx(ctx context.Context, name string, atime, mtime time.Time) error {
	return nil
}

// Stat will fetch file metadata. If the file is not found, it wil return
// os.ErrNotExist
func (fs *ObjStorFs) Stat(name string) (os.FileInfo, error) {
	ctx, _ := fs.getOperationContext("Stat", deadlineKeyMetadataRead)
	return fs.StatEx(ctx, name)
}

// StatEx behaves like Stat, but also allows passing a context with possible
// deadline to lower layers
func (fs *ObjStorFs) StatEx(ctx context.Context, name string) (os.FileInfo, error) {
	return nil, nil
}

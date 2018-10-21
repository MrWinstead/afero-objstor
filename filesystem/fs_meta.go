package filesystem

import (
	"context"
	"os"
	"time"
)

func (fs *ObjStorFs) Chmod(name string, mode os.FileMode) error {
	ctx, _ := fs.getOperationContext("Chmod", DeadlineKeyMetadataWrite)
	return fs.ChmodEx(ctx, name, mode)
}

func (fs *ObjStorFs) ChmodEx(ctx context.Context, name string, mode os.FileMode) error {
	return nil
}

func (fs *ObjStorFs) Chtimes(name string, atime, mtime time.Time) error {
	ctx, _ := fs.getOperationContext("Chtimes", DeadlineKeyMetadataWrite)
	return fs.ChtimesEx(ctx, name, atime, mtime)
}

func (fs *ObjStorFs) ChtimesEx(ctx context.Context, name string, atime, mtime time.Time) error {
	return nil
}

func (fs *ObjStorFs) Stat(name string) (os.FileInfo, error) {
	ctx, _ := fs.getOperationContext("Stat", DeadlineKeyMetadataRead)
	return fs.StatEx(ctx, name)
}

func (fs *ObjStorFs) StatEx(ctx context.Context, name string) (os.FileInfo, error) {
	return nil, nil
}

package filesystem

import (
	"context"
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/afero"
)

func (fs *ObjStorFs) Create(name string) (afero.File, error) {
	ctx, _ := fs.getOperationContext("Create", DeadlineKeyMetadataWrite)
	return fs.CreateEx(ctx, name)
}

func (fs *ObjStorFs) CreateEx(ctx context.Context, name string) (
	*ProjectedObject, error) {
	created, createErr := NewFile(fs, fs.getOperationContext)
	if nil != createErr {
		err := errors.Wrap(createErr, "could not create object")
		return nil, err
	}
	return created, nil
}

func (fs *ObjStorFs) Mkdir(name string, perm os.FileMode) error {
	ctx, _ := fs.getOperationContext("Mkdir", DeadlineKeyWrite)
	return fs.MkdirEx(ctx, name, perm)
}

func (fs *ObjStorFs) MkdirEx(ctx context.Context, name string, perm os.FileMode) error {
	return nil
}

func (fs *ObjStorFs) MkdirAll(path string, perm os.FileMode) error {
	ctx, _ := fs.getOperationContext("MkdirAll", DeadlineKeyWrite)
	return fs.MkdirAllEx(ctx, path, perm)
}

func (fs *ObjStorFs) MkdirAllEx(ctx context.Context, path string, perm os.FileMode) error {
	return nil
}

func (fs *ObjStorFs) Remove(name string) error {
	ctx, _ := fs.getOperationContext("Remove", DeadlineKeyWrite)
	return fs.RemoveEx(ctx, name)
}

func (fs *ObjStorFs) RemoveEx(ctx context.Context, name string) error {
	return nil
}

func (fs *ObjStorFs) RemoveAll(name string) error {
	ctx, _ := fs.getOperationContext("RemoveAll", DeadlineKeyWrite)
	return fs.RemoveAllEx(ctx, name)
}

func (fs *ObjStorFs) RemoveAllEx(ctx context.Context, name string) error {
	return nil
}

func (fs *ObjStorFs) Rename(oldName, newName string) error {
	ctx, _ := fs.getOperationContext("Rename", DeadlineKeyWrite)
	return fs.RenameEx(ctx, oldName, newName)
}

func (fs *ObjStorFs) RenameEx(ctx context.Context, oldName, newName string) error {
	return nil
}

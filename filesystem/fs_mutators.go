package filesystem

import (
	errors2 "afero-objstor/errors"
	"context"
	"fmt"
	"os"
	"path"

	"github.com/pkg/errors"
	"github.com/spf13/afero"
)

// Create will:
// * If the file does not exist in object storage, touch a blank one
// * If the file does exist, copy it into the local storage cache
func (fs *ObjStorFs) Create(name string) (afero.File, error) {
	ctx, _ := fs.getOperationContext("Create", deadlineKeyMetadataWrite)
	return fs.CreateEx(ctx, name)
}

// CreateEx behaves like Create, but also allows passing a context with possible
// deadline to lower layers
func (fs *ObjStorFs) CreateEx(ctx context.Context, name string) (
	DeadlineFile, error) {
	parentDirs, filename := path.Split(name)
	if "/" != parentDirs {
		err := &errors2.UnsupportedOperation{
			OperationFriendlyName: "create",
			Reason:                fmt.Sprintf("file '%v' has parent directories", name),
		}
		return nil, err
	}

	local, err := fs.localFs.Create(filename)
	if nil != err {
		return nil, err
	}

	created, createErr := newFile(fs, local, fileTypeFile)
	if nil != createErr {
		err := errors.Wrap(createErr, "could not create object")
		return nil, err
	}
	fs.filesCache.Add(filename, created)
	return created, nil
}

// Mkdir will not create a directory in object storage. It will create a pseudo-
// directory in the local file cache
func (fs *ObjStorFs) Mkdir(name string, perm os.FileMode) error {
	ctx, _ := fs.getOperationContext("Mkdir", deadlineKeyWrite)
	return fs.MkdirEx(ctx, name, perm)
}

// MkdirEx behaves like Mkdir, but also allows passing a context with possible
// deadline to lower layers
func (fs *ObjStorFs) MkdirEx(ctx context.Context, name string, perm os.FileMode) error {
	return nil
}

// MkdirAll behaves like Mkdir but will also create the parents of a directory
func (fs *ObjStorFs) MkdirAll(path string, perm os.FileMode) error {
	ctx, _ := fs.getOperationContext("MkdirAll", deadlineKeyWrite)
	return fs.MkdirAllEx(ctx, path, perm)
}

// MkdirAllEx behaves like MkdirAll, but also allows passing a context with
// possible deadline to lower layers
func (fs *ObjStorFs) MkdirAllEx(ctx context.Context, path string, perm os.FileMode) error {
	return nil
}

// Remove will attempt to remove the specified file. If the NameField specified is a
// directory, it will only succeed if there are no files under the directory.
func (fs *ObjStorFs) Remove(name string) error {
	ctx, _ := fs.getOperationContext("Remove", deadlineKeyWrite)
	return fs.RemoveEx(ctx, name)
}

// RemoveEx behaves like Remove, but also allows passing a context with
// possible deadline to lower layers
func (fs *ObjStorFs) RemoveEx(ctx context.Context, name string) error {
	return nil
}

// RemoveAll will attempt to remove the specified file. If the NameField specified is a
// directory, it will delete all children recursively
func (fs *ObjStorFs) RemoveAll(name string) error {
	ctx, _ := fs.getOperationContext("RemoveAll", deadlineKeyWrite)
	return fs.RemoveAllEx(ctx, name)
}

// RemoveAllEx behaves like RemoveAll, but also allows passing a context with
// possible deadline to lower layers
func (fs *ObjStorFs) RemoveAllEx(ctx context.Context, name string) error {
	return nil
}

// Rename will attempt to rename the provided filename to the provided new one.
// If the destination NameField already exists, it will fail with os.ErrExist. If
// the file being moved is a directory, all children will also be moved
func (fs *ObjStorFs) Rename(oldName, newName string) error {
	ctx, _ := fs.getOperationContext("Rename", deadlineKeyWrite)
	return fs.RenameEx(ctx, oldName, newName)
}

// RenameEx behaves like Rename, but also allows passing a context with
// possible deadline to lower layers
func (fs *ObjStorFs) RenameEx(ctx context.Context, oldName, newName string) error {
	return nil
}

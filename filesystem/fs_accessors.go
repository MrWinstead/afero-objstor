package filesystem

import (
	"context"
	"github.com/google/go-cloud/blob"
	"github.com/spf13/afero"
	"os"
	"path"
)

// Name fetches the NameField of the filesystem driver
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
func (fs *ObjStorFs) OpenEx(ctx context.Context, name string) (
	file DeadlineFile, err error) {
	fileIface, exists := fs.filesCache.Get(name)
	if exists {
		return fileIface.(*ProjectedObject), nil
	}

	defer func() {
		if nil != err {
			fs.Remove(name)
		}
	}()

	prefixedName := path.Join(fs.keyPrefix, name)
	if '/' == prefixedName[0] {
		prefixedName = prefixedName[1:]
	}
	attrs, err := fs.bucket.Attributes(ctx, prefixedName)
	if nil != err {
		if blob.IsNotExist(err) {
			return nil, os.ErrNotExist
		}
		return nil, err
	}

	createdFile, localCreateErr := fs.localFs.Create(name)
	if nil != localCreateErr {
		return nil, localCreateErr
	}

	serializedFileInfo, exists := attrs.Metadata[metadataKeyAttributes]
	if exists {
		fileInfo, deserializationErr := fileinfoFromJSONStr(serializedFileInfo)
		if nil == deserializationErr {
			applyErr := applyFileInfo(fs, fileInfo)
			if nil != applyErr {
				return nil, applyErr
			}
		}
	}

	projectedFile, projectedCreateErr := newFile(fs, createdFile,
		fileTypeFile)
	if nil != projectedCreateErr {
		return nil, projectedCreateErr
	}
	fs.filesCache.Add(name, projectedFile)

	return projectedFile, nil
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

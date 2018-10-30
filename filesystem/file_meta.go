package filesystem

import (
	"context"
	"github.com/google/go-cloud/blob"
	"github.com/pkg/errors"
	"os"
)

const (
	contentTypeBinary     = "application/octet-stream"
	metadataKeyAttributes = "x-afero-objstor-attrs"
)

// Close flushes changes to the file to object storage and releases all local
// resources related to the file
func (o *ProjectedObject) Close() error {
	ctx, _ := o.ctxGen("Close", deadlineKeyWrite)
	return o.CloseEx(ctx)
}

// CloseEx behaves like Close, but also allows passing a context with possible
// deadline to lower layers
func (o *ProjectedObject) CloseEx(ctx context.Context) error {
	err := o.SyncEx(ctx)
	if nil != err {
		return err
	}

	err = o.localInstance.Close()
	return err
}

// Stat fetches the file's attributes
func (o *ProjectedObject) Stat() (os.FileInfo, error) {
	ctx, _ := o.ctxGen("Stat", deadlineKeyMetadataRead)
	return o.StatEx(ctx)
}

// StatEx fetches file attributes while providing a context with possible
// deadline to lower layers
func (o *ProjectedObject) StatEx(ctx context.Context) (os.FileInfo, error) {
	fi, err := o.localInstance.Stat()
	if nil != err {
		return nil, err
	}
	wrapped := fileinfoFromFileInfo(fi)
	return wrapped, nil
}

// Sync will flush all changes to the file to object storage
func (o *ProjectedObject) Sync() error {
	if fileTypeDirectory == o.fileType { // don't sync dirs
		return nil
	}

	ctx, _ := o.ctxGen("Sync", deadlineKeyWrite)
	return o.SyncEx(ctx)
}

// SyncEx will flush all changes to the file to object storage while providing a
// context with possible deadline to lower layers
func (o *ProjectedObject) SyncEx(ctx context.Context) error {
	if fileTypeDirectory == o.fileType { // don't sync dirs
		return nil
	}

	fileInfoIface, err := o.Stat()
	if nil != err {
		return err
	}
	fi := fileInfoIface.(*fileinfo)

	serializedAttrs := fi.toJSONStr()

	err = o.localInstance.Sync()
	if nil != err {
		return err
	}

	fileContents := make([]byte, fileInfoIface.Size())
	_, err = o.localInstance.ReadAt(fileContents, 0)
	if nil != err {
		return err
	}

	objName := nameToObjKey(o, o.fs.keyPrefix)
	writerOpts := blob.WriterOptions{
		ContentType: contentTypeBinary,
		Metadata: map[string]string{
			metadataKeyAttributes: serializedAttrs,
		},
	}
	bucketWriter, err := o.fs.bucket.NewWriter(ctx, objName, &writerOpts)
	if nil != err {
		return err
	}
	defer bucketWriter.Close()

	var bytesWritten int
	for bytesWritten < len(fileContents) {
		iterationWritten, err := bucketWriter.Write(fileContents)
		if nil != err {
			return err
		}
		if bytesWritten+iterationWritten < bytesWritten {
			err := errors.Errorf(
				"integer overflow while writing data to object storage")
			return err
		}
		bytesWritten += iterationWritten
	}
	return err
}

// SetMode will overwrite the file's current ModeField
func (o *ProjectedObject) SetMode(m os.FileMode) error {
	ctx, _ := o.ctxGen("SetMode", deadlineKeyMetadataWrite)
	return o.SetModeEx(ctx, m)
}

// SetModeEx behaves like SetMode, but also allows passing a context with possible
// deadline to lower layers
func (o *ProjectedObject) SetModeEx(ctx context.Context, m os.FileMode) error {
	err := o.fs.localFs.Chmod(o.Name(), m)
	if nil != err {
		return err
	}
	err = o.Sync()
	return err
}

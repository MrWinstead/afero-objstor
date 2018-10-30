package filesystem

import (
	"context"
	"os"

	"afero-objstor/errors"
)

// Name returns the filename of the object
func (o *ProjectedObject) Name() string {
	return o.localInstance.Name()
}

// Read from the object into the provided buffer
func (o *ProjectedObject) Read(p []byte) (int, error) {
	ctx, _ := o.ctxGen("Read", deadlineKeyRead)
	return o.ReadEx(ctx, p)
}

// ReadEx behaves like Read, but also allows passing a context with possible
// deadline to lower layers
func (o *ProjectedObject) ReadEx(ctx context.Context, p []byte) (int, error) {
	if o.fs.opts.writeOptimize || !o.fs.opts.readOptimize {
		o.Sync()
	}
	return o.localInstance.Read(p)
}

// ReadAt pulls bytes from the object into the provided buffer at the file's
// current offset
func (o *ProjectedObject) ReadAt(p []byte, off int64) (int, error) {
	ctx, _ := o.ctxGen("ReadAt", deadlineKeyRead)
	return o.ReadAtEx(ctx, p, off)
}

// ReadAtEx behaves like ReadAt, but also allows passing a context with possible
// deadline to lower layers
func (o *ProjectedObject) ReadAtEx(ctx context.Context, p []byte, off int64) (
	int, error) {
	if o.fs.opts.writeOptimize || !o.fs.opts.readOptimize {
		o.Sync()
	}
	return o.localInstance.ReadAt(p, off)
}

// Readdir enumerates file information about all immediate children of this file
// if it is a directory. It will return os.ErrInvalid when called on a file
func (o *ProjectedObject) Readdir(count int) ([]os.FileInfo, error) {
	ctx, _ := o.ctxGen("Readdir", deadlineKeyMetadataRead)
	return o.ReaddirEx(ctx, count)
}

// ReaddirEx behaves like Readdir, but also allows passing a context with
// possible deadline to lower layers
func (o *ProjectedObject) ReaddirEx(ctx context.Context, count int) (
	[]os.FileInfo, error) {
	if fileTypeFile == o.fileType {
		return nil, os.ErrInvalid
	}
	err := &errors.UnsupportedOperation{
		OperationFriendlyName: "Readdir",
		Reason:                "only files currenly supported",
	}
	return nil, err
}

// Readdirnames behaves like Readdir, but only returns the filenames
func (o *ProjectedObject) Readdirnames(count int) ([]string, error) {
	ctx, _ := o.ctxGen("Readdirnames", deadlineKeyMetadataRead)
	return o.ReaddirnamesEx(ctx, count)
}

// ReaddirnamesEx behaves like Readdirnames, but also allows passing a context
// with possible deadline to lower layers
func (o *ProjectedObject) ReaddirnamesEx(ctx context.Context, count int) (
	[]string, error) {
	fullInfo, infoGatherErr := o.Readdir(count)
	if nil != infoGatherErr {
		return nil, infoGatherErr
	}
	names := make([]string, len(fullInfo))
	for i := range fullInfo {
		names[i] = fullInfo[i].Name()
	}
	return names, nil
}

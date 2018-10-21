package filesystem

import (
	"afero-objstor/errors"
	"context"
	"os"
)

// Truncate will cut the file to the specified size
func (o *ProjectedObject) Truncate(size int64) error {
	ctx, _ := o.ctxGen("Truncate", deadlineKeyWrite)
	return o.TruncateEx(ctx, size)
}

// TruncateEx behaves like Truncate, but also allows passing a context with possible
// deadline to lower layers
func (o *ProjectedObject) TruncateEx(ctx context.Context, size int64) error {
	err := &errors.UnsupportedOperation{
		OperationFriendlyName: "Truncate",
		Reason:                "not implemented",
	}
	return err
}

// Seek will set the read/write offset within the file. If an invalid whence
// is provided, os.ErrInvalid is returned. This functionality does not check for
// integer over/underflows
func (o *ProjectedObject) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case 0:
		o.currentPosition = offset
	case 1:
		o.currentPosition += offset
	case 2:
		o.currentPosition = o.size - offset
	default:
		return 0, os.ErrInvalid
	}
	return o.currentPosition, nil
}

// WriteAt writes the provided bytes at the provided offset
func (o *ProjectedObject) WriteAt(p []byte, offset int64) (int, error) {
	ctx, _ := o.ctxGen("WriteAt", deadlineKeyWrite)
	return o.WriteAtEx(ctx, p, offset)
}

// WriteAtEx behaves like WriteAt, but also allows passing a context with possible
// deadline to lower layers
func (o *ProjectedObject) WriteAtEx(ctx context.Context, p []byte, offset int64) (int, error) {
	err := &errors.UnsupportedOperation{
		OperationFriendlyName: "WriteAt",
		Reason:                "not implemented",
	}
	return 0, err
}

// Write writes the provided bytes to the current read/write file offset
func (o *ProjectedObject) Write(p []byte) (int, error) {
	ctx, _ := o.ctxGen("Write", deadlineKeyWrite)
	return o.WriteEx(ctx, p)
}

// WriteEx behaves like Write, but also allows passing a context with possible
// deadline to lower layers
func (o *ProjectedObject) WriteEx(ctx context.Context, p []byte) (int, error) {
	return o.WriteAtEx(ctx, p, o.currentPosition)
}

// WriteString calls Write with the string converted to UTF-8 bytes
func (o *ProjectedObject) WriteString(s string) (int, error) {
	ctx, _ := o.ctxGen("WriteString", deadlineKeyWrite)
	return o.WriteStringEx(ctx, s)
}

// WriteStringEx behaves like WriteString, but also allows passing a context with possible
// deadline to lower layers
func (o *ProjectedObject) WriteStringEx(ctx context.Context, s string) (int, error) {
	stringBytes := []byte(s)
	return o.WriteEx(ctx, stringBytes)
}

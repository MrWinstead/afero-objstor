package filesystem

import (
	"context"

	"afero-objstor/errors"
)

func (o *ProjectedObject) Truncate(size int64) error {
	ctx, _ := o.ctxGen("Truncate", DeadlineKeyWrite)
	return o.TruncateEx(ctx, size)
}

func (o *ProjectedObject) TruncateEx(ctx context.Context, size int64) error {
	err := &errors.UnsupportedOperation{
		OperationFriendlyName: "Truncate",
		Reason: "not implemented",
	}
	return err
}

func (o *ProjectedObject) Seek(offset int64, whence int) (int64, error) {
	ctx, _ := o.ctxGen("Seek", DeadlineKeyMetadataWrite)
	return o.SeekEx(ctx, offset, whence)
}

func (o *ProjectedObject) SeekEx(ctx context.Context, offset int64, whence int) (int64, error) {
	err := &errors.UnsupportedOperation{
		OperationFriendlyName: "Seek",
		Reason: "not currently supported",
	}
	return 0, err
}

func (o *ProjectedObject) WriteAt(p []byte, offset int64) (int, error) {
	ctx, _ := o.ctxGen("WriteAt", DeadlineKeyWrite)
	return o.WriteAtEx(ctx, p, offset)
}

func (o *ProjectedObject) WriteAtEx(ctx context.Context, p []byte, offset int64) (int, error) {
	err := &errors.UnsupportedOperation{
		OperationFriendlyName: "WriteAt",
		Reason: "not implemented",
	}
	return 0, err
}

func (o *ProjectedObject) Write(p []byte) (int, error) {
	ctx, _ := o.ctxGen("Write", DeadlineKeyWrite)
	return o.WriteEx(ctx, p)
}

func (o *ProjectedObject) WriteEx(ctx context.Context, p []byte) (int, error) {
	return o.WriteAtEx(ctx, p, o.currentPosition)
}

func (o *ProjectedObject) WriteString(s string) (int, error) {
	ctx, _ := o.ctxGen("WriteString", DeadlineKeyWrite)
	return o.WriteStringEx(ctx, s)
}

func (o *ProjectedObject) WriteStringEx(ctx context.Context, s string) (int, error) {
	stringBytes := []byte(s)
	return o.WriteEx(ctx, stringBytes)
}

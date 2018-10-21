package filesystem

import (
	"context"
	"os"

	"afero-objstor/errors"
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
	err := &errors.UnsupportedOperation{
		OperationFriendlyName: "Close",
		Reason:                "not implemented",
	}
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
	info := &fileinfo{obj: o}
	return info, nil
}

// Sync will flush all changes to the file to object storage
func (o *ProjectedObject) Sync() error {
	ctx, _ := o.ctxGen("Sync", deadlineKeyWrite)
	return o.SyncEx(ctx)
}

// SyncEx will flush all changes to the file to object storage while providng a
// context with possible deadline to lower layers
func (o *ProjectedObject) SyncEx(ctx context.Context) error {
	err := &errors.UnsupportedOperation{
		OperationFriendlyName: "Sync",
		Reason:                "not implemented",
	}
	return err
}

func (o *ProjectedObject) syncMetadata(ctx context.Context) error {
	return nil
}

// SetMode will overwrite the file's current mode
func (o *ProjectedObject) SetMode(m os.FileMode) error {
	ctx, _ := o.ctxGen("SetMode", deadlineKeyMetadataWrite)
	return o.SetModeEx(ctx, m)
}

// SetModeEx behaves like SetMode, but also allows passing a context with possible
// deadline to lower layers
func (o *ProjectedObject) SetModeEx(ctx context.Context, m os.FileMode) error {
	o.mode = m
	if o.opts.WriteThrough {
		o.Sync()
	}
	return nil
}

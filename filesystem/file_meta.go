package filesystem

import (
	"context"
	"os"

	"afero-objstor/errors"
)

func (o *ProjectedObject) Close() error {
	ctx, _ := o.ctxGen("Close", DeadlineKeyWrite)
	return o.CloseEx(ctx)
}

func (o *ProjectedObject) CloseEx(ctx context.Context) error {
	err := &errors.UnsupportedOperation{
		OperationFriendlyName: "Close",
		Reason: "not implemented",
	}
	return err
}


func (o *ProjectedObject) Stat() (os.FileInfo, error) {
	ctx, _ := o.ctxGen("Stat", DeadlineKeyMetadataRead)
	return o.StatEx(ctx)
}

func (o *ProjectedObject) StatEx(ctx context.Context) (os.FileInfo, error) {
	info := &fileinfo{obj: o}
	return info, nil
}

func (o *ProjectedObject) Sync() error {
	ctx, _ := o.ctxGen("Sync", DeadlineKeyWrite)
	return o.SyncEx(ctx)
}

func (o *ProjectedObject) SyncEx(ctx context.Context) error {
	err := &errors.UnsupportedOperation{
		OperationFriendlyName: "Sync",
		Reason: "not implemented",
	}
	return err
}

func (o *ProjectedObject) SetMode(m os.FileMode) error {
	ctx, _ := o.ctxGen("SetMode", DeadlineKeyMetadataWrite)
	return o.SetModeEx(ctx, m)
}

func (o *ProjectedObject) SetModeEx(ctx context.Context, m os.FileMode) error {
	o.mode = m
	if o.opts.WriteThrough {
		o.Sync()
	}
	return nil
}

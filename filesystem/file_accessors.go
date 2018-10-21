package filesystem

import (
	"context"
	"os"

	"afero-objstor/errors"
)

func (o *ProjectedObject) Name() string {
	return o.Name()
}

func (o *ProjectedObject) Read(p []byte) (int, error) {
	ctx, _ := o.ctxGen("Read", DeadlineKeyRead)
	return o.ReadEx(ctx, p)
}

func (o *ProjectedObject) ReadEx(ctx context.Context, p []byte) (int, error) {
	err := &errors.UnsupportedOperation{
		OperationFriendlyName: "Read",
		Reason: "not implemented",
	}
	return 0, err
}

func (o *ProjectedObject) ReadAt(p []byte, off int64) (int, error) {
	ctx, _ := o.ctxGen("ReadAt", DeadlineKeyRead)
	return o.ReadAtEx(ctx, p, off)
}

func (o *ProjectedObject) ReadAtEx(ctx context.Context, p []byte, off int64) (int, error) {
	err := &errors.UnsupportedOperation{
		OperationFriendlyName: "ReadAt",
		Reason: "not implemented",
	}
	return 0, err
}

func (o *ProjectedObject) Readdir(count int) ([]os.FileInfo, error) {
	ctx, _ := o.ctxGen("Readdir", DeadlineKeyMetadataRead)
	return o.ReaddirEx(ctx, count)
}

func (o *ProjectedObject) ReaddirEx(ctx context.Context, count int) ([]os.FileInfo, error) {
	err := &errors.UnsupportedOperation{
		OperationFriendlyName:"Readdir",
		Reason: "only files currenly supported",
	}
	return nil, err
}

func (o *ProjectedObject) Readdirnames(count int) ([]string, error) {
	ctx, _ := o.ctxGen("Readdirnames", DeadlineKeyMetadataRead)
	return o.ReaddirnamesEx(ctx, count)
}

func (o *ProjectedObject) ReaddirnamesEx(ctx context.Context, count int) ([]string, error) {
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
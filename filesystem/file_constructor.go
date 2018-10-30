package filesystem

import (
	"afero-objstor/errors"
	"fmt"
	"github.com/spf13/afero"
)

func newFile(fs *ObjStorFs, localInstance afero.File, fileType int) (
	*ProjectedObject, error) {
	switch fileType {
	case fileTypeFile:
	case fileTypeDirectory:
		break
	default:
		err := &errors.UnsupportedOperation{
			OperationFriendlyName: "newFile",
			Reason:                fmt.Sprintf("unsupported fileType %v", fileType),
		}
		return nil, err
	}

	created := &ProjectedObject{
		localInstance: localInstance,
		fileType:      fileType,
		fs:            fs,
		ctxGen:        fs.getOperationContext,
	}

	return created, nil
}

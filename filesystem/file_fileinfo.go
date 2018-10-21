package filesystem

import (
	"os"
	"time"
)

func (fi *fileinfo) Name() string {
	return fi.obj.name
}

func (fi *fileinfo) Size() int64 {
	return 0
}

func (fi *fileinfo) Mode() os.FileMode {
	return 0
}

func (fi *fileinfo) ModTime() time.Time {
	return time.Unix(0, 0)
}

func (fi *fileinfo) IsDir() bool {
	return false
}

func (fi *fileinfo) Sys() interface{} {
	return nil
}

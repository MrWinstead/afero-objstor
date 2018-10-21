package filesystem

import (
	"os"
	"time"
)

// Get the filename
func (fi *fileinfo) Name() string {
	return fi.obj.name
}

// Size which the file occupies in object storage
func (fi *fileinfo) Size() int64 {
	return 0
}

// Fetch the emulated UNIX-style file mode of the object
func (fi *fileinfo) Mode() os.FileMode {
	return 0
}

// Fetch the timestamp of the file
func (fi *fileinfo) ModTime() time.Time {
	return time.Unix(0, 0)
}

// Check if the file is a directory. Since object storage does not have
// directories, this is true for pseudo-directories maintained using key
// prefixes
func (fi *fileinfo) IsDir() bool {
	return false
}

// Returns nil and is unsupported
func (fi *fileinfo) Sys() interface{} {
	return nil
}

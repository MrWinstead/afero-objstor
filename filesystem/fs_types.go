package filesystem

import (
	"time"

	"github.com/google/go-cloud/blob"
	"github.com/hashicorp/golang-lru"
	"github.com/spf13/afero"
)

// ConstructorOption to provide to the constructor of an Object storage filesystem
type ConstructorOption func(fs *ObjStorFs) error

type fsOptions struct {
	deadlines map[string]time.Duration
}

// ObjStorFs is an Object Storage-backed filesystem
type ObjStorFs struct {
	bucket    *blob.Bucket
	keyPrefix string
	localFs   afero.Fs

	directoryCacheSizeMax int
	directoryCache        *lru.Cache

	totalCachedFilesSizeMax uint64
	filesCacheSizeMax       int
	filesCache              *lru.Cache

	opts fsOptions
}

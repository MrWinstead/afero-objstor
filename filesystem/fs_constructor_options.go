package filesystem

import (
	"time"

	"github.com/google/go-cloud/blob"
	"github.com/spf13/afero"
)

const (
	// DefaultMaxCachedDirectories is the number of directories kept in local
	// cache
	DefaultMaxCachedDirectories = 10000

	// DefaultMaxCachedFiles is the number of directories kept in local cache
	DefaultMaxCachedFiles = 10000

	// DefaultMaxCachedFilesSize is the total sizes of files kept in the local
	// cache, default is 1 GB
	DefaultMaxCachedFilesSize = uint64(1 * 1024 * 1024 * 1024)

	defaultReadDeadline          = 1 * time.Minute
	defaultWriteDeadline         = 1 * time.Minute
	defaultMetadataReadDeadline  = 1 * time.Minute
	defaultMetadataWriteDeadline = 1 * time.Minute

	deadlineKeyRead          = "operation-read"
	deadlineKeyWrite         = "operation-write"
	deadlineKeyMetadataRead  = "operation-metadata-read"
	deadlineKeyMetadataWrite = "operation-metadata-write"
)

var (
	defaultOptions = []ConstructorOption{
		WithKeyPrefix("/"),
		WithBackingFilesystem(afero.NewMemMapFs()),
		WithDefaultReadDeadline(defaultReadDeadline),
		WithDefaultWriteDeadline(defaultWriteDeadline),
		WithDefaultMetadataReadDeadline(defaultMetadataReadDeadline),
		WithDefaultMetadataWriteDeadline(defaultMetadataWriteDeadline),
		WithMaxCachedFilesCount(DefaultMaxCachedFiles),
		WithMaxCachedDirectoryCount(DefaultMaxCachedDirectories),
		WithMaxCachedFilesSize(DefaultMaxCachedFilesSize),
	}
)

// WithBucket sets the object storage bucket to which objects will be read and
// written
func WithBucket(bucket *blob.Bucket) ConstructorOption {
	return func(fs *ObjStorFs) error {
		fs.bucket = bucket
		return nil
	}
}

// WithKeyPrefix will prepend the specified prefix to all objects created in the
// underlying bucket
func WithKeyPrefix(prefix string) ConstructorOption {
	return func(fs *ObjStorFs) error {
		fs.keyPrefix = prefix
		return nil
	}
}

// WithBackingFilesystem will set the afero filesystem to be used as a local
// file cache for reads and writes
func WithBackingFilesystem(fs afero.Fs) ConstructorOption {
	return func(objStoreFs *ObjStorFs) error {
		objStoreFs.localFs = fs
		return nil
	}
}

// WithMaxCachedFilesCount will set the maximum number of files which will be
// kept in the local filesystem cache
func WithMaxCachedFilesCount(count int) ConstructorOption {
	return func(fs *ObjStorFs) error {
		fs.filesCacheSizeMax = count
		return nil
	}
}

// WithMaxCachedDirectoryCount sets the max number of directories which are
// kept in the local filesystem cache
func WithMaxCachedDirectoryCount(count int) ConstructorOption {
	return func(fs *ObjStorFs) error {
		fs.directoryCacheSizeMax = count
		return nil
	}
}

// WithMaxCachedFilesSize sets the total bytes of files which will be stored
// in the local filesystem cache
func WithMaxCachedFilesSize(maxSize uint64) ConstructorOption {
	return func(fs *ObjStorFs) error {
		fs.totalCachedFilesSizeMax = maxSize
		return nil
	}
}

// WithDefaultReadDeadline the default deadline for all data read operations
// (e.g. Read, ReadAt)
func WithDefaultReadDeadline(d time.Duration) ConstructorOption {
	return func(fs *ObjStorFs) error {
		fs.opts.deadlines[deadlineKeyRead] = d
		return nil
	}
}

// WithDefaultWriteDeadline sets the default deadline for all data write
// operations (e.g. Write, WriteAt)
func WithDefaultWriteDeadline(d time.Duration) ConstructorOption {
	return func(fs *ObjStorFs) error {
		fs.opts.deadlines[deadlineKeyWrite] = d
		return nil
	}
}

// WithDefaultMetadataReadDeadline the default deadline for all metadata read
// operations (e.g. Stat)
func WithDefaultMetadataReadDeadline(d time.Duration) ConstructorOption {
	return func(fs *ObjStorFs) error {
		fs.opts.deadlines[deadlineKeyMetadataRead] = d
		return nil
	}
}

// WithDefaultMetadataWriteDeadline sets the default deadline for all metadata
// write operations (e.g. Chtimes, Chmod)
func WithDefaultMetadataWriteDeadline(d time.Duration) ConstructorOption {
	return func(fs *ObjStorFs) error {
		fs.opts.deadlines[deadlineKeyMetadataWrite] = d
		return nil
	}
}

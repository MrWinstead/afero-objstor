package filesystem

import (
	"time"

	"github.com/google/go-cloud/blob"
	"github.com/spf13/afero"
)

const (
	DefaultMaxCachedDirectories = 10000
	DefaultMaxCachedFiles = 10000

	// 1 GB
	DefaultMaxCachedFilesSize = uint64(1 * 1024 * 1024 * 1024)

	DefaultReadDeadline          = 1 * time.Minute
	DefaultWriteDeadline         = 1 * time.Minute
	DefaultMetadataReadDeadline  = 1 * time.Minute
	DefaultMetadataWriteDeadline = 1 * time.Minute

	DeadlineKeyRead          = "operation-read"
	DeadlineKeyWrite         = "operation-write"
	DeadlineKeyMetadataRead  = "operation-metadata-read"
	DeadlineKeyMetadataWrite = "operation-metadata-write"
)

var (
	defaultOptions = []FsConstructorOption{
		WithKeyPrefix("/"),
		WithDefaultReadDeadline(DefaultReadDeadline),
		WithDefaultWriteDeadline(DefaultWriteDeadline),
		WithDefaultMetadataReadDeadline(DefaultMetadataReadDeadline),
		WithDefaultMetadataWriteDeadline(DefaultMetadataWriteDeadline),
	}
)

func WithBucket(bucket *blob.Bucket) FsConstructorOption {
	return func(fs *ObjStorFs) error {
		fs.bucket = bucket
		return nil
	}
}

func WithKeyPrefix(prefix string) FsConstructorOption {
	return func(fs *ObjStorFs) error {
		fs.keyPrefix = prefix
		return nil
	}
}

func WithBackingFilesystem(fs afero.Fs) FsConstructorOption {
	return func(objStoreFs *ObjStorFs) error {
		objStoreFs.localFs = fs
		return nil
	}
}

// Use this to set the maximum runtime for a specific operation (e.g. Open)
func WithOperationDeadline(operationName string, d time.Duration) FsConstructorOption {
	return func(fs *ObjStorFs) error {
		fs.opts.deadlines[operationName] = d
		return nil
	}
}

// Set the default deadline for all data read operations (e.g. Read, ReadAt)
func WithDefaultReadDeadline(d time.Duration) FsConstructorOption {
	return func(fs *ObjStorFs) error {
		fs.opts.deadlines[DeadlineKeyRead] = d
		return nil
	}
}

// Set the default deadline for all data write operations (e.g. Write, WriteAt)
func WithDefaultWriteDeadline(d time.Duration) FsConstructorOption {
	return func(fs *ObjStorFs) error {
		fs.opts.deadlines[DeadlineKeyWrite] = d
		return nil
	}
}

// Set the default deadline for all metadata read operations (e.g. Stat)
func WithDefaultMetadataReadDeadline(d time.Duration) FsConstructorOption {
	return func(fs *ObjStorFs) error {
		fs.opts.deadlines[DeadlineKeyMetadataRead] = d
		return nil
	}
}

// Set the default deadline for all metadata write operations (e.g. Chtimes, Chmod)
func WithDefaultMetadataWriteDeadline(d time.Duration) FsConstructorOption {
	return func(fs *ObjStorFs) error {
		fs.opts.deadlines[DeadlineKeyMetadataWrite] = d
		return nil
	}
}

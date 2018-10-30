# Afero Object Storage Driver

[![CircleCI](https://circleci.com/gh/MrWinstead/afero-objstor/tree/master.svg?style=svg)](https://circleci.com/gh/MrWinstead/afero-objstor/tree/master)
[![codecov](https://codecov.io/gh/MrWinstead/afero-objstor/branch/master/graph/badge.svg)](https://codecov.io/gh/MrWinstead/afero-objstor)
[![GoDoc](https://godoc.org/github.com/MrWinstead/afero-objstor?status.svg)](https://godoc.org/github.com/MrWinstead/afero-objstor)

## Overview

A filesystem backend for the
[afero filesystem abstraction](https://github.com/spf13/afero) which is designed
to work with [go-cloud](https://github.com/google/go-cloud) supported object
storage systems. Additionally, it provides an extended interface beyond that of
afero which offers deadline sensitivity provided by golang contexts.

## Notes on Distributed System Use

This library is likely to be used in a cloud-native context with multiple
workers all using the same backend object storage bucket. This library does not
make promises regarding file locking, consistency, or multiple-writer
situations. This would best be solved with locks currently managed outside of
this library.

## Features
### Currently Supported

```go
type DeadlineFs interface {
	Name() string
	
	Create(name string) (File, error)
	CreateEx(ctx context.Context, name string) (DeadlineFile, error)
	
	Open(name string) (File, error)
	OpenEx(ctx context.Context, name string) (DeadlineFile, error)
}

type DeadlineFile interface {
	io.Reader
	ReadEx(ctx context.Context, p []byte) (count int, err error)
	io.ReaderAt
	ReadAtEx(ctx context.Context, p []byte, off int64) (count int, err error)
	
	io.Writer
	WriteEx(ctx context.Context, p []byte) (count int, err error)
	io.WriterAt
	WriteAtEx(ctx context.Context, p []byte, off int64) (count int, err error)
	WriteString(s string) (ret int, err error)
	WriteStringEx(ctx context.Context, s string) (int, error)
	
	Sync() error
	SyncEx(ctx context.Context) error
}
```

### Currently Unsupported
```go
type DeadlineFs interface {
	Mkdir(name string, perm os.FileMode) error
	MkdirEx(ctx context.Context, name string, perm os.FileMode) error
	MkdirAll(path string, perm os.FileMode) error
	MkdirAllEx(ctx context.Context, path string, perm os.FileMode) error
	
	OpenFile(name string, flag int, perm os.FileMode) (File, error)
	OpenFileEx(ctx context.Context, name string, flag int, perm os.FileMode) (DeadlineFile, error)
	
	Remove(name string) error
	RemoveEx(ctx context.Context, name string) error
	RemoveAll(path string) error
	RemoveAllEx(ctx context.Context, path string) error
	
	Rename(oldname, newname string) error
	RenameEx(ctx context.Context, oldname, newname string) error
	
	Stat(name string) (os.FileInfo, error)
	StatEx(ctx context.Context, name string) (os.FileInfo, error)
	
	Chmod(name string, mode os.FileMode) error
	ChmodEx(ctx context.Context, name string, mode os.FileMode) error
	Chtimes(name string, atime time.Time, mtime time.Time) error
	ChtimesEx(ctx context.Context, name string, atime time.Time, mtime time.Time) error
}

type DeadlineFile interface {
	io.Seeker
	Name() string
	
	io.Closer
	CloseEx(ctx context.Context) error

	Readdir(count int) ([]os.FileInfo, error)
	ReaddirEx(ctx context.Context, count int) ([]os.FileInfo, error)
	Readdirnames(n int) ([]string, error)
	ReaddirnamesEx(ctx context.Context, count int) ([]string, error)
	
	Stat() (os.FileInfo, error)
	StatEx(ctx context.Context) (os.FileInfo, error)
	
	Truncate(size int64) error
	TruncateEx(ctx context.Context, size int64) error
}
```

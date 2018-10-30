package filesystem

import (
	"context"
	"os"
	"time"

	"github.com/spf13/afero"
)

// DeadlineFs describes a super-set of afero.Fs which respects contexts
type DeadlineFs interface {
	afero.Fs

	// Create creates a file in the filesystem, returning the file and an
	// error, if any happens.
	CreateEx(ctx context.Context, name string) (DeadlineFile, error)

	// Mkdir creates a directory in the filesystem, return an error if any
	// happens.
	MkdirEx(ctx context.Context, name string, perm os.FileMode) error

	// MkdirAll creates a directory path and all parents that does not exist
	// yet.
	MkdirAllEx(ctx context.Context, path string, perm os.FileMode) error

	// Open opens a file, returning it or an error, if any happens.
	OpenEx(ctx context.Context, name string) (DeadlineFile, error)

	// OpenFile opens a file using the given flags and the given ModeField.
	OpenFileEx(ctx context.Context, name string, flag int, perm os.FileMode) (DeadlineFile, error)

	// Remove removes a file identified by NameField, returning an error, if any
	// happens.
	RemoveEx(ctx context.Context, name string) error

	// RemoveAll removes a directory path and any children it contains. It
	// does not fail if the path does not exist (return nil).
	RemoveAllEx(ctx context.Context, path string) error

	// Rename renames a file.
	RenameEx(ctx context.Context, oldname, newname string) error

	// Stat returns a FileInfo describing the named file, or an error, if any
	// happens.
	StatEx(ctx context.Context, name string) (os.FileInfo, error)

	//Chmod changes the ModeField of the named file to ModeField.
	ChmodEx(ctx context.Context, name string, mode os.FileMode) error

	//Chtimes changes the access and modification times of the named file
	ChtimesEx(ctx context.Context, name string, atime time.Time, mtime time.Time) error
}

// DeadlineFile describes a super-set of afero.File which respects contexts
type DeadlineFile interface {
	afero.File

	// Close the file. If the operation cannot complete within the specified time,
	// the file is left open
	CloseEx(ctx context.Context) error

	// Attempt to read until the context expires or another error is encountered
	ReadEx(ctx context.Context, p []byte) (count int, err error)

	// Attempt to read at position until the context expires or another error is
	// encountered
	ReadAtEx(ctx context.Context, p []byte, off int64) (count int, err error)

	// Attempt to write the provided bytes until the context expires
	WriteEx(ctx context.Context, p []byte) (count int, err error)

	// Attempt to write at position the provided bytes until the context expires
	WriteAtEx(ctx context.Context, p []byte, off int64) (count int, err error)

	ReaddirEx(ctx context.Context, count int) ([]os.FileInfo, error)
	ReaddirnamesEx(ctx context.Context, count int) ([]string, error)
	StatEx(ctx context.Context) (os.FileInfo, error)
	SyncEx(ctx context.Context) error
	TruncateEx(ctx context.Context, size int64) error
	WriteStringEx(ctx context.Context, s string) (int, error)
}

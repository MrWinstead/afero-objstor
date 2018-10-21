package filesystem

import (
	"time"

	"github.com/hashicorp/golang-lru"
	"github.com/pkg/errors"
)

// NewFs creates a new object storage filesytem with the specified options
func NewFs(options ...ConstructorOption) (*ObjStorFs, error) {
	created := &ObjStorFs{
		opts: fsOptions{
			deadlines: make(map[string]time.Duration),
		},
	}

	options = append(defaultOptions, options...)

	for _, opt := range options {
		optionErr := opt(created)
		if nil != optionErr {
			return nil, optionErr
		}
	}

	if 0 == created.directoryCacheSizeMax {
		created.directoryCacheSizeMax = DefaultMaxCachedDirectories
	}
	directoryCache, err := lru.NewWithEvict(created.directoryCacheSizeMax,
		created.evictDirectoryCallback)
	if nil != err {
		err = errors.Wrap(err, "could not create directory cache")
		return nil, err
	}
	created.directoryCache = directoryCache

	if 0 == created.totalCachedFilesSizeMax {
		created.totalCachedFilesSizeMax = DefaultMaxCachedFilesSize
	}
	if 0 == created.filesCacheSizeMax {
		created.filesCacheSizeMax = DefaultMaxCachedFiles
	}
	filesCache, err := lru.NewWithEvict(created.filesCacheSizeMax,
		created.evictFileCallback)
	if nil != err {
		err = errors.Wrap(err, "could not create files cache")
		return nil, err
	}
	created.filesCache = filesCache

	return created, nil
}

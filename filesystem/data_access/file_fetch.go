package data_access

import (
	"context"
	"path"

	"afero-objstor/filesystem"
	"github.com/hashicorp/golang-lru"
)

func PopulateFileIntoCache(ctx context.Context, cache *lru.Cache,
	fs *filesystem.ObjStorFs, fileName string) error {
	fullPath := path.Clean(path.Join(fs.GetKeyPrefix(), fileName))
	backingBucket := fs.GetBucket()

	_, attrFetchErr := backingBucket.Attributes(ctx, fullPath)
	if nil != attrFetchErr {
		return attrFetchErr
	}
	return nil
}

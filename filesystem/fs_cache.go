package filesystem

func (fs *ObjStorFs) evictFileCallback(k, v interface{}) {
	fileName := k.(string)
	fs.localFs.Remove(fileName)
}

func (fs *ObjStorFs) evictDirectoryCallback(k, v interface{}) {
	dirName := k.(string)
	fs.localFs.RemoveAll(dirName)
}

package filesystem

func (fs *ObjStorFs) EvictFileCallback(k, v interface{}) {
	fileName := k.(string)
	fs.localFs.Remove(fileName)
}

func (fs *ObjStorFs) EvictDirectoryCallback(k, v interface{}) {
	dirName := k.(string)
	fs.localFs.RemoveAll(dirName)
}


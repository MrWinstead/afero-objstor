package filesystem

func newFile(fs *ObjStorFs) (*ProjectedObject, error) {
	created := &ProjectedObject{
		fs:     fs,
		ctxGen: fs.getOperationContext,
	}

	return created, nil
}

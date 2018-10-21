package filesystem

func NewFile(fs *ObjStorFs, ctxGen ContextGenerator) (*ProjectedObject, error) {
	created := &ProjectedObject{
		fs:     fs,
		ctxGen: ctxGen,
	}

	return created, nil
}

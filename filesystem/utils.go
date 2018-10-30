package filesystem

import "path"

func nameToObjKey(f *ProjectedObject, keyPrefix string) string {
	fullPath := path.Join(keyPrefix, f.Name())
	if '/' == fullPath[0] {
		fullPath = fullPath[1:]
	}
	cleaned := path.Clean(fullPath)
	return cleaned
}

package filesystem

import (
	"encoding/json"
	"os"
	"time"
)

type fileinfo struct {
	NameField  string      `json:"Name"`
	ModeField  os.FileMode `json:"Mode"`
	SizeField  int64       `json:"Size"`
	MtimeField time.Time   `json:"Mtime"`
	IsDirField bool        `json:"IsDir"`
}

func (f fileinfo) Name() string {
	return f.NameField
}

func (f fileinfo) Size() int64 {
	return f.SizeField
}

func (f fileinfo) Mode() os.FileMode {
	return f.ModeField
}

func (f fileinfo) ModTime() time.Time {
	return f.MtimeField
}

func (f fileinfo) IsDir() bool {
	return f.IsDirField
}

func (f fileinfo) Sys() interface{} {
	return nil
}

func (f fileinfo) toJsonStr() string {
	serialized, serializeErr := json.Marshal(f)
	if nil != serializeErr {
		panic(serializeErr)
	}
	return string(serialized)
}

func fileinfoFromJsonStr(serialized string) (*fileinfo, error) {
	fi := &fileinfo{}
	err := json.Unmarshal([]byte(serialized), fi)
	if nil != err {
		return nil, err
	}

	return fi, nil
}

func fileinfoFromFileInfo(fi os.FileInfo) *fileinfo {
	created := &fileinfo{
		NameField:  fi.Name(),
		SizeField:  fi.Size(),
		ModeField:  fi.Mode(),
		MtimeField: fi.ModTime(),
		IsDirField: fi.IsDir(),
	}
	return created
}

func applyFileInfo(fs DeadlineFs, info *fileinfo) error {
	chtimeErr := fs.Chtimes(info.Name(), time.Unix(0, 0), info.MtimeField)
	if nil != chtimeErr {
		return chtimeErr
	}

	chmodErr := fs.Chmod(info.NameField, info.Mode())
	return chmodErr
}

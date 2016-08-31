package session

import (
	"encoding/json"
	"os"
	"russ/system/util"
)

type FileStore struct {
}

func NewFileStore() *FileStore {
	return &FileStore{}
}

func (f *FileStore) Init(id string) SessionInterface {
	path := os.Getenv("GOPATH") + "/src/russ/cache/session"
	if ok := util.IsExist(path); !ok {
		os.MkdirAll(path, 0755)
	}

	filename := path + "/sess_" + id
	fd, _ := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0666)
	defer fd.Close()

	fs := &fileStore{
		filename: filename,
		Values:   make(map[string]interface{}),
	}
	data, err := util.ReadAll(fd)

	if err == nil {
		json.Unmarshal(data, &fs.Values)
	}

	/*decoder := gob.NewDecoder(fd)
	decoder.Decode(fs.Values)*/

	return fs
}

func (f *FileStore) GC() {

}

type fileStore struct {
	filename string
	Values   map[string]interface{}
	fd       *os.File
}

func (f *fileStore) Get(key string) interface{} {
	value, ok := f.Values[key]
	if !ok {
		return nil
	}
	return value
}

func (f *fileStore) Set(key string, value interface{}) {
	f.Values[key] = value
}

func (f *fileStore) Remove(key string) {
	delete(f.Values, key)
}

func (f *fileStore) Flush() {

}

func (f *fileStore) Save() error {
	fd, err := os.OpenFile(f.filename, os.O_RDWR, 0666)
	defer fd.Close()

	if err != nil {
		return err
	}

	buf, _ := json.Marshal(f.Values)
	fd.Write(buf)
	/*encoder := gob.NewEncoder(fd)
	err = encoder.Encode(f.Values)*/
	return err
}

package file

import (
	"io/ioutil"
)

type FsOperations interface {
	ReadFile() ([]byte, error)
}

type File struct {
	filename string
}

func New(filename string) File {
	return File{
		filename: filename,
	}
}

func (f File) ReadFile() ([]byte, error) {
	b, err := ioutil.ReadFile(f.filename)
	if err != nil {
		return nil, err
	}
	return b, nil
}

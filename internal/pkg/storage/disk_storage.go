package storage

import (
	"os"
	"path/filepath"
)

type DiskStorage struct {
	dir string
}

func New(dir string) *DiskStorage {
	return &DiskStorage{
		dir: dir,
	}
}

func (d *DiskStorage) Load(name string) ([]byte, error) {
	return os.ReadFile(filepath.Join(d.dir, name))
}

func (d *DiskStorage) Save(name string, data []byte) error {
	f, err := os.OpenFile(filepath.Join(d.dir, name), os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}
	if _, err := f.Write(data); err != nil {
		return err
	}
	return nil
}

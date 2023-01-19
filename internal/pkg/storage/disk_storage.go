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

func (d *DiskStorage) Save(name string, data []byte) (string, error) {
	path := filepath.Join(d.dir, name)
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return "", err
	}
	defer f.Close()
	if _, err := f.Write(data); err != nil {
		_ = f.Close()
		_ = os.Remove(path)
		return "", err
	}
	return path, nil
}

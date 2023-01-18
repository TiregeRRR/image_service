package store

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
	return os.WriteFile(filepath.Join(d.dir, name), data, 0666)
}

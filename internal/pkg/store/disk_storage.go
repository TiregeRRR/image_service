package store

type DiskStorage struct {
	dir string
}

func New(dir string) *DiskStorage {
	return &DiskStorage{
		dir: dir,
	}
}

// TODO
func Load(name string) ([]byte, error) {
	return nil, nil
}

// TODO
func Save(name string, data []byte) error {
	return nil
}

package storage

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSave(t *testing.T) {
	testCases := []struct {
		name string
		data []byte
	}{
		{
			name: "someName.txt",
			data: []byte("some data"),
		},
	}
	tmp := t.TempDir()
	r := require.New(t)
	for _, tc := range testCases {
		ds := New(tmp)
		path, err := ds.Save(tc.name, tc.data)
		r.NoError(err)
		expectedPath := filepath.Join(tmp, tc.name)
		r.Equal(expectedPath, path)
		b, err := os.ReadFile(expectedPath)
		r.NoError(err)
		r.Equal(b, tc.data)
	}
}

func TestLoad(t *testing.T) {
	testCases := []struct {
		name string
		data []byte
	}{
		{
			name: "someName.txt",
			data: []byte("some data"),
		},
	}
	tmp := t.TempDir()
	r := require.New(t)
	for _, tc := range testCases {
		path := filepath.Join(tmp, tc.name)
		err := os.WriteFile(path, tc.data, os.ModePerm)
		r.NoError(err)
		ds := New(tmp)
		b, err := ds.Load(tc.name)
		r.NoError(err)
		r.Equal(tc.data, b)
	}
}

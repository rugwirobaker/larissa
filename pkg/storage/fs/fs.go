package fs

import (
	"fmt"

	"github.com/rugwirobaker/larissa/pkg/storage"
	"github.com/rugwirobaker/larissa/pkg/types"
	"github.com/spf13/afero"
)

var _ (storage.Backend) = (*fsBackend)(nil)

// fsBackend is a filesystem implementation of larissa storage Backend
// that uses afero FileSystem Abstraction.
type fsBackend struct {
	rootDir string
	fs      afero.Fs
}

// NewBackend creates a new storage Backend instance.
func NewBackend(rootDir string, filesystem afero.Fs) (storage.Backend, error) {
	exists, err := afero.Exists(filesystem, rootDir)
	if err != nil {
		return nil, fmt.Errorf("could not check if root directory `%s` exists: %s", rootDir, err)
	}
	if !exists {
		return nil, fmt.Errorf("root directory `%s` does not exist", rootDir)
	}
	return &fsBackend{rootDir: rootDir, fs: filesystem}, nil
}

func (fs *fsBackend) Put(path string, content []byte) error {
	return nil
}
func (fs *fsBackend) Get(path string) (*types.Object, error) {
	return nil, nil
}
func (fs *fsBackend) Del(path string) error {
	return nil
}
func (fs *fsBackend) Exists(path string) bool {
	return false
}

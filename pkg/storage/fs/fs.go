package fs

import (
	"fmt"
	"path/filepath"

	"github.com/rugwirobaker/larissa/pkg/errors"

	"github.com/rugwirobaker/larissa/pkg/storage"
	"github.com/rugwirobaker/larissa/pkg/types"
	"github.com/spf13/afero"
)

var _ (storage.Backend) = (*backend)(nil)

// backend is a filesystem implementation of larissa storage Backend
// that uses afero FileSystem Abstraction.
type backend struct {
	rootDir    string
	filesystem afero.Fs
}

// NewBackend creates a new storage Backend instance.
func NewBackend(rootDir string, filesystem afero.Fs) (storage.Backend, error) {
	const op errors.Op = "fs.NewStorage"

	exists, err := afero.Exists(filesystem, rootDir)
	if err != nil {
		return nil, errors.E(op, fmt.Errorf("could not check if root directory `%s` exists: %s", rootDir, err))
	}
	if !exists {
		return nil, errors.E(op, fmt.Errorf("root directory `%s` does not exist", rootDir))
	}
	return &backend{rootDir: rootDir, filesystem: filesystem}, nil
}

// bucket gets bucket(subfolder)
func (fs *backend) bucketLocation(bucket string) string {
	return filepath.Join(fs.rootDir, bucket)
}

func (fs *backend) Put(file, bucket string, content []byte) error {
	const op errors.Op = "fs.Put"

	path := fs.bucketLocation(bucket)
	exists, err := afero.DirExists(fs.filesystem, path)
	if err != nil {
		return errors.E(op, err, errors.O(file), errors.B(bucket))
	}
	if !exists {
		fs.filesystem.MkdirAll(path, 0777)
	}
	err = afero.WriteFile(fs.filesystem, filepath.Join(path, file), content, 0666)
	if err != nil {
		return errors.E(op, err, errors.O(file), errors.B(bucket))
	}
	return nil
}
func (fs *backend) Get(file, bucket string) (*types.Object, error) {
	const op errors.Op = "fs.Get"
	return nil, nil
}
func (fs *backend) Del(file, bucket string) error {
	const op errors.Op = "fs.Del"
	return nil
}
func (fs *backend) Exists(file, bucket string) bool {
	const op errors.Op = "fs.Exists"
	return false
}

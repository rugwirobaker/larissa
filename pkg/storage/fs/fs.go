package fs

import (
	"context"
	"fmt"
	"os"
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
func (fs *backend) bucketLoc(bucket string) string {
	return filepath.Join(fs.rootDir, bucket)
}

func (fs *backend) fileLoc(bucket, file string) string {
	return filepath.Join(fs.rootDir, bucket, file)
}

func (fs *backend) Put(ctx context.Context, file, bucket string, content []byte) error {
	const op errors.Op = "fs.Put"

	path := fs.bucketLoc(bucket)
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
func (fs *backend) Get(ctx context.Context, file, bucket string) (*types.Object, error) {
	const op errors.Op = "fs.Get"

	path := fs.fileLoc(bucket, file)
	exists, err := afero.Exists(fs.filesystem, path)
	if err != nil {
		return nil, errors.E(op, err, errors.O(file), errors.B(bucket), errors.KindNotFound)
	}
	if !exists {
		return nil, errors.E(op, errors.O(file), errors.B(bucket), errors.KindNotFound)
	}

	content, err := afero.ReadFile(fs.filesystem, path)
	if err != nil {
		return nil, errors.E(op, err, errors.O(file), errors.B(bucket), errors.KindNotFound)
	}

	return &types.Object{Name: file, Content: content}, nil
}
func (fs *backend) Del(ctx context.Context, file, bucket string) error {
	const op errors.Op = "fs.Del"

	path := fs.fileLoc(bucket, file)
	exists, err := afero.Exists(fs.filesystem, path)
	if err != nil {
		return errors.E(op, err, errors.O(file), errors.B(bucket), errors.KindNotFound)
	}
	if !exists {
		return errors.E(op, errors.O(file), errors.B(bucket), errors.KindNotFound)
	}
	if err := fs.filesystem.Remove(path); err != nil {
		return errors.E(op, errors.O(file), errors.B(bucket), errors.KindNotFound)
	}

	return nil
}
func (fs *backend) Exists(ctx context.Context, file, bucket string) error {
	const op errors.Op = "fs.Exists"

	path := fs.fileLoc(bucket, file)

	exists, err := afero.Exists(fs.filesystem, path)
	if err != nil {
		return errors.E(op, err, errors.O(file), errors.B(bucket), errors.KindNotFound)
	}

	if !exists {
		return errors.E(op, err, errors.O(file), errors.B(bucket), errors.KindNotFound)
	}
	return nil
}

func (fs *backend) List(ctx context.Context, bucket string) ([]string, error) {
	const op errors.Op = "fs.List"

	path := fs.bucketLoc(bucket)

	exists, err := afero.Exists(fs.filesystem, path)
	if err != nil {
		return []string{}, errors.E(op, err, errors.B(bucket), errors.KindNotFound)
	}

	if !exists {
		return []string{}, errors.E(op, err, errors.B(bucket), errors.KindNotFound)
	}

	fileInfos, err := afero.ReadDir(fs.filesystem, path)

	if err != nil {
		if os.IsNotExist(err) {
			return []string{}, nil
		}

		return nil, errors.E(op, errors.B(bucket), err, errors.KindUnexpected)
	}

	objects := []string{}

	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			continue
		}
		name := fileInfo.Name()
		if name == "" {
			continue
		}
		objects = append(objects, fileInfo.Name())

	}
	return objects, nil
}

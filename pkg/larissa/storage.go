package larissa

// Backend describes the storage inteface
type Backend interface {
	Put(path string, content []byte) error
	Get(path string) (*Object, error)
	Del(path string) error
	Exists(path string) bool
}

// fileSystemBackend is a filesystem implementation of larissa storage Backend.
type fileSystemBackend struct{}

// NewBackend creates a new Backend instance.
func NewBackend() Backend {
	return &fileSystemBackend{}
}

func (fs *fileSystemBackend) Put(path string, content []byte) error { return nil }
func (fs *fileSystemBackend) Get(path string) (*Object, error)      { return nil, nil }
func (fs *fileSystemBackend) Del(path string) error                 { return nil }
func (fs *fileSystemBackend) Exists(path string) bool               { return false }

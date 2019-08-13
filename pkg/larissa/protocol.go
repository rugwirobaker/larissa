package larissa

import (
	"github.com/rugwirobaker/larissa/pkg/storage"
	"github.com/rugwirobaker/larissa/pkg/types"
)

// Protocol descirbes the user acceible interface to end users
type Protocol interface {
	Put(file, bucket string, content []byte) error
	Get(file, bucket string) (*types.Object, error)
	Del(file, bucket string) error
	Exists(file, bucket string) error
}

var _ (Protocol) = (*protocol)(nil)

type protocol struct {
	backend storage.Backend
}

// New creates a new larissa protocol
func New(backend storage.Backend) Protocol {
	return &protocol{backend}
}

func (proctl *protocol) Put(file, bucket string, content []byte) error {
	return proctl.backend.Put(file, bucket, content)
}

func (proctl *protocol) Get(file, bucket string) (*types.Object, error) {
	return proctl.backend.Get(file, bucket)
}

func (proctl *protocol) Del(file, bucket string) error {
	return proctl.backend.Del(file, bucket)
}

func (proctl *protocol) Exists(file, bucket string) error {
	return proctl.backend.Exists(file, bucket)
}

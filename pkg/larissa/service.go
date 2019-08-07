package larissa

import (
	"github.com/rugwirobaker/larissa/pkg/storage"
	"github.com/rugwirobaker/larissa/pkg/types"
)

// Service descirbes the user acceible interface to end users
type Service interface {
	Put(path string, content []byte) error
	Get(path string) (*types.Object, error)
	Del(path string) error
	Exists(path string) bool
}

var _ (Service) = (*service)(nil)

type service struct {
	backend storage.Backend
}

// New creates a new larissa service
func New(backend storage.Backend) Service {
	return &service{backend}
}

func (svc *service) Put(path string, content []byte) error {
	return svc.backend.Put(path, content)
}

func (svc *service) Get(path string) (*types.Object, error) {
	return svc.backend.Get(path)
}

func (svc *service) Del(path string) error {
	return svc.backend.Del(path)
}

func (svc *service) Exists(path string) bool {
	return svc.backend.Exists(path)
}

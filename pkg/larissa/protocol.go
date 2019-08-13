package larissa

import (
	"github.com/rugwirobaker/larissa/pkg/storage"
	"github.com/rugwirobaker/larissa/pkg/types"
)

// Service descirbes the user acceible interface to end users
type Service interface {
	Put(file, bucket string, content []byte) error
	Get(file, bucket string) (*types.Object, error)
	Del(file, bucket string) error
	Exists(file, bucket string) bool
}

var _ (Service) = (*service)(nil)

type service struct {
	backend storage.Backend
}

// New creates a new larissa service
func New(backend storage.Backend) Service {
	return &service{backend}
}

func (svc *service) Put(file, bucket string, content []byte) error {
	return svc.backend.Put(file, bucket, content)
}

func (svc *service) Get(file, bucket string) (*types.Object, error) {
	return svc.backend.Get(file, bucket)
}

func (svc *service) Del(file, bucket string) error {
	return svc.backend.Del(file, bucket)
}

func (svc *service) Exists(file, bucket string) bool {
	return svc.backend.Exists(file, bucket)
}

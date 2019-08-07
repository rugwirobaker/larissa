package storage

import (
	"github.com/rugwirobaker/larissa/pkg/types"
)

// Backend describes the storage inteface
type Backend interface {
	Put(path string, content []byte) error
	Get(path string) (*types.Object, error)
	Del(path string) error
	Exists(path string) bool
}

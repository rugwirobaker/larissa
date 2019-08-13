package storage

import (
	"github.com/rugwirobaker/larissa/pkg/types"
)

// Backend describes the larissa storage inteface
type Backend interface {
	Put(file, bucket string, content []byte) error
	Get(file, bucket string) (*types.Object, error)
	Del(file, bucket string) error
	Exists(file, bucket string) error
}

package storage

import (
	"context"

	"github.com/rugwirobaker/larissa/pkg/types"
)

// Backend describes the larissa storage inteface
type Backend interface {
	Put(ctx context.Context, file, bucket string, content []byte) error
	Get(ctx context.Context, file, bucket string) (*types.Object, error)
	Del(ctx context.Context, file, bucket string) error
	Exists(ctx context.Context, file, bucket string) error
	List(ctx context.Context, bucket string) ([]string, error)
}

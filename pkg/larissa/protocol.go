package larissa

import (
	"context"

	"github.com/rugwirobaker/larissa/pkg/errors"

	"github.com/rugwirobaker/larissa/pkg/storage"
	"github.com/rugwirobaker/larissa/pkg/types"
)

// Protocol descirbes the user accessible interface to end users
type Protocol interface {
	Put(ctx context.Context, file, bucket string, content []byte) error
	Get(ctx context.Context, file, bucket string) (*types.Object, error)
	Del(ctx context.Context, file, bucket string) error
	Exists(ctx context.Context, file, bucket string) error
	List(ctx context.Context, bucket string) ([]string, error)
}

var _ (Protocol) = (*protocol)(nil)

type protocol struct {
	backend storage.Backend
}

// New creates a new larissa protocol
func New(backend storage.Backend) Protocol {
	return &protocol{backend}
}

func (proctl *protocol) Put(ctx context.Context, file, bucket string, content []byte) error {
	const op errors.Op = "protocol.Put"

	_, err := ext(content)
	if err != nil {
		return errors.E(op, err, errors.O(file), errors.B(bucket), errors.KindBadRequest)
	}
	return proctl.backend.Put(ctx, file, bucket, content)
}

func (proctl *protocol) Get(ctx context.Context, file, bucket string) (*types.Object, error) {
	return proctl.backend.Get(ctx, file, bucket)
}

func (proctl *protocol) Del(ctx context.Context, file, bucket string) error {
	return proctl.backend.Del(ctx, file, bucket)
}

func (proctl *protocol) Exists(ctx context.Context, file, bucket string) error {
	return proctl.backend.Exists(ctx, file, bucket)
}

func (proctl *protocol) List(ctx context.Context, bucket string) ([]string, error) {
	return proctl.backend.List(ctx, bucket)
}

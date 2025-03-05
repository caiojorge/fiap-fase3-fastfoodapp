package sharedports

import "context"

type RepositoryInterface[T any] interface {
	Create(ctx context.Context, entity T) error
	Update(ctx context.Context, entity T) error
	Find(ctx context.Context, id string) (T, error)
	FindAll(ctx context.Context) ([]T, error)
	Delete(ctx context.Context, id string) error
}

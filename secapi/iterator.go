package secapi

import (
	"context"
	"errors"
	"io"

	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/schema"
	"github.com/eu-sovereign-cloud/go-sdk/pkg/types"
)

// IteratorFunc is a function that can be used to fetch a collection of items.
type IteratorFunc[T types.ResourceType] func(ctx context.Context, skipToken *string) ([]T, *schema.ResponseMetadata, error)

// Iterator is a generic iterator that can be used to iterate over a collection of items.
type Iterator[T types.ResourceType] struct {
	fn   IteratorFunc[T]
	ptr  int
	data []T
	meta schema.ResponseMetadata
}

// NewIterator creates a new iterator.
func NewIterator[T types.ResourceType](fn IteratorFunc[T]) *Iterator[T] {
	return &Iterator[T]{
		fn: fn,
	}
}

// Next returns the next item in the iterator. If there are no more items, it returns io.EOF.
func (i *Iterator[T]) Next(ctx context.Context) (*T, error) {
	// If we have items in the buffer, return the next one
	if i.ptr < len(i.data) {
		result := &i.data[i.ptr]
		i.ptr++
		return result, nil
	}

	// If we don't have a skip token and the buffer is empty, we're done
	if i.meta.SkipToken == nil && len(i.data) > 0 {
		return nil, io.EOF
	}

	// Fetch more items
	newData, newMeta, err := i.fn(ctx, i.meta.SkipToken)
	if err != nil {
		return nil, err
	}

	// Update the buffer and skip token
	i.data = newData
	i.meta = *newMeta

	// If we have items in the buffer, return the first one
	if len(i.data) > 0 {
		result := &i.data[0]
		i.ptr = 1
		return result, nil
	}

	return nil, io.EOF
}

// All returns all items in the iterator.
func (i *Iterator[T]) All(ctx context.Context) ([]*T, error) {
	var items []*T
	for {
		item, err := i.Next(context.Background())
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

// Returns the current response metadata
func (i *Iterator[T]) Metadata() schema.ResponseMetadata {
	return i.meta
}

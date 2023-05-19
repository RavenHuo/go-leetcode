package base

import (
	"context"
)

type BloomFilter interface {
	Add(ctx context.Context, key []byte) error
	Adds(ctx context.Context, keys [][]byte) error
	Exists(ctx context.Context, key []byte) (bool, error)
	BatchExists(ctx context.Context, key [][]byte) ([]bool, error)
	Clear()
}

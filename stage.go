package main

import "context"

type Stage[T Request] interface {
	Execute(ctx context.Context, req *T) (*T, error)
}

package main

import "context"

type ExecuteFunc[T Request] func(context.Context, *T) (*T, error)

func (f ExecuteFunc[T]) Execute(ctx context.Context, req *T) (*T, error) {
	return f(ctx, req)
}

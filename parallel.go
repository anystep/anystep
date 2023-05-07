package main

import (
	"context"

	"golang.org/x/sync/errgroup"
)

type parallel[T Request] struct {
	stages []Stage[T]
	merge  MergeFunc[T]
}

type MergeFunc[T Request] func(ctx context.Context, req *T, resps ...*T) (*T, error)

func Parallel[T Request](merge MergeFunc[T], stages ...Stage[T]) *parallel[T] {
	return &parallel[T]{
		stages: stages,
		merge:  merge,
	}
}

func (p *parallel[T]) Execute(ctx context.Context, req *T) (*T, error) {
	resps := make([]*T, len(p.stages))
	g, groupCtx := errgroup.WithContext(ctx)

	for i := range p.stages {
		i := i
		g.Go(func() error {
			req2 := req
			resp, err := p.stages[i].Execute(groupCtx, req2)
			if err != nil {
				return err
			}
			resps[i] = resp
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return nil, err
	}

	return p.merge(ctx, req, resps...)
}

package main

import "context"

type series[T Request] struct {
	stages []Stage[T]
}

func Series[T Request](stages ...Stage[T]) *series[T] {
	return &series[T]{stages: stages}
}

func (s *series[T]) Execute(ctx context.Context, req *T) (*T, error) {
	var err error
	resp := req

	for _, stage := range s.stages {
		resp, err = stage.Execute(ctx, resp)
		if err != nil {
			return nil, err
		}
	}

	return resp, nil
}

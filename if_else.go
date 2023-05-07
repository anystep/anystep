package main

import "context"

type ifElse[T Request] struct {
	condition func(ctx context.Context, req *T) (bool, error)
	ifTrue    Stage[T]
	ifFalse   Stage[T]
}

func IfElse[T Request](
	condition func(ctx context.Context, req *T) (bool, error),
	ifTrue Stage[T],
	ifFalse Stage[T],
) *ifElse[T] {
	return &ifElse[T]{
		condition: condition,
		ifTrue:    ifTrue,
		ifFalse:   ifFalse,
	}
}

func (ie *ifElse[T]) Execute(ctx context.Context, req *T) (*T, error) {
	condResult, err := ie.condition(ctx, req)
	if err != nil {
		return nil, err
	}

	if condResult {
		return ie.ifTrue.Execute(ctx, req)
	}

	return ie.ifFalse.Execute(ctx, req)
}

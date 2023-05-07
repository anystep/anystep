package main

type Middleware[T Request] func(stage Stage[T]) Stage[T]

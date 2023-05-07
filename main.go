package main

import (
	"context"
	"fmt"
)

func Result() Stage[Request] {
	return ExecuteFunc[Request](func(ctx context.Context, req *Request) (*Request, error) {

		for k, v := range req.Data {
			fmt.Println(k, v)
		}

		return req, nil
	})
}

func ReadInput(field string) Stage[Request] {
	return ExecuteFunc[Request](func(ctx context.Context, req *Request) (*Request, error) {

		var input string
		fmt.Scanln(&input)

		req.Data[field] = input

		return req, nil
	})
}

func Pipeline(dependencies ...string) Stage[Request] {
	return Series(
		ReadInput(dependencies[0]),
		ReadInput(dependencies[1]),
		Result(),
	)
}

func main() {
	Pipeline("field1", "field2").Execute(context.Background(), &Request{
		Data: make(map[string]string),
	})
}

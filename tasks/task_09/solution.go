package main

import (
	"context"
	"errors"
	"sync"
)

type Result[R any] struct {
	Value R
	Err   error
}

var ErrInvalidWorkers = errors.New("workers must be positive")

func ParallelMap[T any, R any](ctx context.Context, workers int, in []T,
	fn func(context.Context, T) (R, error)) ([]Result[R], error) {
	if workers <= 0 {
		return nil, ErrInvalidWorkers
	}

	if err := ctx.Err(); err != nil {
		return nil, err
	}

	n := len(in)

	if n == 0 {
		return []Result[R]{}, nil
	}

	works := make(chan int, n)
	res := make([]Result[R], n)

	grts := min(workers, n)
	var wg sync.WaitGroup

	for range grts {
		wg.Go(func() {
			innerWorker(ctx, fn, works, in, res)
		})
	}

	for i := range n {
		works <- i
	}
	close(works)

	wg.Wait()

	if err := ctx.Err(); err != nil {
		return nil, err
	}

	return res, nil
}

func innerWorker[T any, R any](ctx context.Context, fn func(context.Context, T) (R, error),
	works <-chan int, in []T, res []Result[R]) {
	for i := range works {
		if err := ctx.Err(); err != nil {
			break
		}
		result, err := fn(ctx, in[i])
		if err != nil {
			res[i].Err = err
		} else {
			res[i].Value = result
		}
	}
}

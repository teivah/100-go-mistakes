package main

import (
	"context"
	"sync"

	"golang.org/x/sync/errgroup"
)

func handler1(ctx context.Context, circles []Circle) ([]Result, error) {
	results := make([]Result, len(circles))
	wg := sync.WaitGroup{}
	wg.Add(len(results))

	for i, circle := range circles {
		i := i
		circle := circle
		go func() {
			defer wg.Done()

			result, err := foo(ctx, circle)
			if err != nil {
			}
			// ?
			results[i] = result
		}()
	}

	wg.Wait()
	// ...
	return results, nil
}

func handler2(ctx context.Context, circles []Circle) ([]Result, error) {
	results := make([]Result, len(circles))
	g, ctx := errgroup.WithContext(ctx)

	for i, circle := range circles {
		i := i
		circle := circle
		g.Go(func() error {
			result, err := foo(ctx, circle)
			if err != nil {
				return err
			}
			results[i] = result
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return nil, err
	}
	return results, nil
}

func foo(context.Context, Circle) (Result, error) {
	return Result{}, nil
}

type (
	Circle struct{}
	Result struct{}
)

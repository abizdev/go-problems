package taskrunner

import (
	"context"
	"errors"
	"sync"
)

type Task func(context.Context) (any, error)

type TaskResult struct {
	Value any
	Err   error
}

func RunTasks(ctx context.Context, concurrency int, tasks []Task) ([]TaskResult, error) {

	jobs := make(chan int, len(tasks))
	result := make([]TaskResult, len(tasks))

	wg := sync.WaitGroup{}

	if concurrency <= 0 {
		return result, errors.New("cuncurrency can'n be less or equal to 0")
	}

	if len(tasks) == 0 {
		return result, nil
	}

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for {
				select {
				case <-ctx.Done():
					return

				case index, ok := <-jobs:
					if !ok {
						return
					}

					if ctx.Err() != nil {
						return
					}

					if tasks[index] == nil {
						result[index] = TaskResult{
							Err: errors.New("nil task"),
						}
						continue
					}

					value, err := tasks[index](ctx)
					result[index] = TaskResult{
						Value: value,
						Err:   err,
					}

				}
			}
		}()
	}

	for index, _ := range tasks {
		if ctx.Err() != nil {
			break
		}

		jobs <- index
	}

	close(jobs)
	wg.Wait()

	if err := ctx.Err(); err != nil {
		return result, err
	}

	return result, nil
}

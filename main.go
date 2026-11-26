package main

import (
	"context"
	"errors"
	"fmt"

	// merge "go-problems/merge-channel"
	successfulRequest "go-problems/first-successful-request"
	requestFailover "go-problems/request-failover"
	taskRunnner "go-problems/task-runner"

	// requestFailover "go-problems/request-failover"
	"net/http"
	"time"
)

func failOverResult() {
	client := &http.Client{}
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*500)
	defer cancel()

	urls := []string{
		// "https://jsonplaceholder.typicode.com/posts",
		// "https://jsonplaceholder.typicode.com/comments",
		// "https://jsonplaceholder.typicode.com/albums",
	}

	fmt.Println(requestFailover.RequestWithFailover(&ctx, client, urls))
}

func firstSuccessFulRequestCall() {
	http := http.Client{}
	urls := []string{
		"https://jsonplaceholder.typicode.com/posts",
		"https://jsonplaceholder.typicode.com/comments",
		"https://jsonplaceholder.typicode.com/albums",
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	result, err := successfulRequest.FirstSuccessfulRequest(ctx, &http, urls)
	fmt.Println(result, err)
}

func taskRunnerCall() {
	ctx := context.Background()
	tasks := []taskRunnner.Task{
		func(ctx context.Context) (any, error) {
			return "First", nil
		},

		func(ctx context.Context) (any, error) {
			return nil, errors.New("something went wrong")
		},

		func(ctx context.Context) (any, error) {
			return "Third", nil
		},
	}

	result, err := taskRunnner.RunTasks(ctx, 3, tasks)
	fmt.Println(result, err)
}

func main() {
	// merge.MergeChannels()
	// failOverResult()
	// reateLimiter.NewRateLimiter(5, 3)
	// firstSuccessFulRequestCall()
	taskRunnerCall()
}

package main

import (
	"context"
	"fmt"

	// merge "go-problems/merge-channel"
	reateLimiter "go-problems/rate-limiter"
	requestFailover "go-problems/request-failover"
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

func main() {
	// merge.MergeChannels()
	// failOverResult()
	reateLimiter.NewRateLimiter(5, 3)
}

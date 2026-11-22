package rateLimiter

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

type RateLimiter struct {
	ticker *time.Ticker
}

func NewRateLimiter(n int, d time.Duration) (*RateLimiter, error) {

	if n <= 0 || d <= 0 {
		return &RateLimiter{}, errors.New("введенные параментры не правильны, должны быть только положительные числа и время")
	}

	wg := &sync.WaitGroup{}
	limit := time.Second / time.Duration(d)
	limiter := &RateLimiter{ticker: time.NewTicker(limit)}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(limit))
	defer cancel()

	wg.Add(n)
	for i := 1; i <= n; i++ {
		fmt.Println(i)
		limiter.Wait(ctx, wg)
	}

	wg.Wait()

	return limiter, nil
}

func (r *RateLimiter) Wait(ctx context.Context, wg *sync.WaitGroup) error {
	defer wg.Done()
	fmt.Println("waiting...")

	select {
	case <-ctx.Done():
		return errors.New("context closed")
	}
}

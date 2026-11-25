package firstSuccessfulRequest

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sync"
)

type FetchResult struct {
	URL  string
	Body []byte
}

func worker(
	wg *sync.WaitGroup,
	result chan FetchResult,
	client *http.Client,
	url string,
) {
	defer wg.Done()

	resp, err := client.Get(url)
	if err != nil {
		fmt.Println("get error", err)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Read error:", err)
	}

	result <- FetchResult{
		URL:  url,
		Body: body,
	}

	resp.Body.Close()
}

func FirstSuccessfulRequest(ctx context.Context, client *http.Client, urls []string) (FetchResult, error) {
	if len(urls) == 0 {
		return FetchResult{}, errors.New("no urls")
	}

	wg := &sync.WaitGroup{}
	ch := make(chan FetchResult, 1)

	wg.Add(len(urls))

	for _, url := range urls {
		go worker(wg, ch, client, url)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for {
		select {
		case <-ctx.Done():
			return FetchResult{}, ctx.Err()
		case value, ok := <-ch:
			if !ok {
				return value, errors.New("error in all requests")
			}

			return value, nil
		}

	}
}

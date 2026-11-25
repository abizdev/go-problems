package requestFailover

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type FailoverResult struct {
	URL  string
	Body []byte
}	

func RequestWithFailover(ctx *context.Context, client *http.Client, urls []string) (FailoverResult, error) {
	var result FailoverResult
	context, _ := context.WithCancel(*ctx)

	if len(urls) == 0 {
		return result, errors.New("no urls")
	}

	select {
	case <-context.Done():
		fmt.Println("canceled")
		break
	default:
		for _, url := range urls {
			resp, err := client.Get(url)
			if err != nil {
				continue
			}

			body, err := io.ReadAll(resp.Body)
			result.URL = url
			result.Body = body

			break
		}
	}

	if result.URL == "" {
		return result, errors.New("all ulrs return error")
	}

	return result, nil
}

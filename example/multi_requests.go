package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
)

const (
	invalidUrl = "http://localhost:3000/"
	urlLength  = 5
)

func generateUrls() []string {
	urls := make([]string, urlLength)
	for i := 0; i < urlLength; i++ {
		urls = append(urls, "https://google.com")
	}

	return urls
}

// makeConcurrentRequestV2 using ctx and wait group
func makeConcurrentRequestV2() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	urls := generateUrls()
	urls = append(urls, invalidUrl)

	resultChan := make(chan interface{})
	var wg sync.WaitGroup
	wg.Add(len(urls))

	startTime := time.Now().Unix()

	for _, url := range urls {
		go func(
			wg *sync.WaitGroup,
			ctx context.Context,
			url string,
		) {
			defer wg.Done()

			// Check if the context has been canceled
			if ctx.Err() != nil {
				return
			}

			// Make the HTTP request
			res, err := http.Get(url)
			if err != nil {
				fmt.Println("Error occurred, canceling all requests:", err)
				cancel() // Cancel all requests on error
				return
			}

			select {
			case <-ctx.Done():
				// If context is canceled, stop processing
				return
			case resultChan <- res: // Send the result if not canceled
			}
		}(&wg, ctx, url)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// Process results
	for result := range resultChan {
		fmt.Println("Received result:", result)
	}

	endTime := time.Now().Unix()
	fmt.Printf("Took makeConcurrentRequestV2: %ds \n", endTime-startTime)
}

// makeConcurrentRequestV1 using error group
func makeConcurrentRequestV1() {
	urls := generateUrls()
	urls = append(urls, invalidUrl)

	resultChan := make(chan interface{})
	defer close(resultChan)

	g, ctx := errgroup.WithContext(context.Background())
	g.SetLimit(2)

	for _, url := range urls {
		g.Go(func() error {
			select {
			case <-ctx.Done():
				fmt.Println("Context Done")
				return nil
			default:
				res, err := http.Get(url)
				if err != nil {
					fmt.Println("Error occurred, canceling all requests:", err)
					return err
				}

				resultChan <- res
				return nil
			}
		})
	}

	if err := g.Wait(); err != nil {
		fmt.Print("Error occured: ", err.Error())
		return
	}

	for result := range resultChan {
		fmt.Println("Received result:", result)
	}
}

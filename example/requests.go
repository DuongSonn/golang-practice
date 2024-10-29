package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

const (
	invalidUrl = "http://localhost:3000/"
	urlLength  = 5
)

func generateUrls() []string {
	urls := make([]string, urlLength)
	for i := 0; i < urlLength; i++ {
		urls = append(urls, "https://example.com")
	}

	return urls
}

func makeConcurrentRequest() {
	urls := generateUrls()

	startTime := time.Now().Unix()
	var wg sync.WaitGroup
	wg.Add(len(urls))

	for _, url := range urls {
		go func() {
			defer wg.Done()

			_, err := http.Get(url)
			if err != nil {
				fmt.Println("Error: \n", err.Error())
			}
		}()
	}

	wg.Wait()
	endTime := time.Now().Unix()
	fmt.Printf("Took makeConcurrentRequest: %ds \n", endTime-startTime)
}

func makeSingleRequest() {
	startTime := time.Now().Unix()
	for _, url := range generateUrls() {
		_, err := http.Get(url)
		if err != nil {
			fmt.Println("Error: \n", err.Error())
		}
	}
	endTime := time.Now().Unix()
	fmt.Printf("Took makeSingleRequest: %ds \n", endTime-startTime)
}

// makeConcurrentRequestV2 stop when the 1st error is met
func makeConcurrentRequestV2() {
	urls := generateUrls()
	resultChan := make(chan interface{})
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	wg.Add(len(urls) + 1)

	startTime := time.Now().Unix()

	go func() {
		defer wg.Done()

		if ctx.Err() != nil {
			return
		}

		// Make the request
		res, err := http.Get(invalidUrl)
		if err != nil {
			fmt.Println("Error occurred, canceling all requests:", err)
			cancel()
			return
		}

		select {
		case <-ctx.Done():
			return
		case resultChan <- res:
		}
	}()

	for _, url := range urls {
		go func(url string) {
			defer wg.Done()

			if ctx.Err() != nil {
				return
			}

			res, err := http.Get(url)
			if err != nil {
				fmt.Println("Error occurred, canceling all requests:", err)
				cancel()
				return
			}

			select {
			case <-ctx.Done():
				return
			case resultChan <- res:
			}
		}(url)
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

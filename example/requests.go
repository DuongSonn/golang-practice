package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

const (
	invalidUrl = "http://localhost:3000/"
	urlLength  = 10
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

	startTime := time.Now().Unix()

	errChan := make(chan error)
	var wg sync.WaitGroup
	wg.Add(len(urls) + 1)

	go func(
		errChan chan error,
	) {
		defer wg.Done()

		err, ok := <-errChan
		if !ok {
			fmt.Println("Channel closed")
			return
		}
		if err != nil {
			fmt.Println("Error Detected")
			return
		}

		if _, err := http.Get(invalidUrl); err != nil {
			close(errChan)
			fmt.Println("Error: \n", err.Error())
		}

	}(errChan)

	for _, url := range urls {
		go func(
			errChan chan error,
		) {
			defer wg.Done()

			err, ok := <-errChan
			if !ok {
				fmt.Println("Channel closed")
				return
			}
			if err != nil {
				fmt.Println("Error Detected")

				return
			}

			if _, err := http.Get(url); err != nil {
				close(errChan)
				fmt.Println("Error: \n", err.Error())
			}
		}(errChan)
	}

	wg.Wait()
	endTime := time.Now().Unix()
	fmt.Printf("Took makeConcurrentRequestV2: %ds \n", endTime-startTime)
}

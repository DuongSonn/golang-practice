package main

import (
	"fmt"
	"net/http"
)

type result struct {
	Error    error
	Response *http.Response
}

func errorPatter() {
	checkStatus := func(done <-chan interface{}, urls ...string) <-chan result {
		results := make(chan result)
		go func() {
			defer close(results)
			for _, url := range urls {
				resp, err := http.Get(url)
				res := result{
					Response: resp,
					Error:    err,
				}
				select {
				case <-done:
					return
				case results <- res:
				}
			}
		}()

		return results
	}

	done := make(chan interface{})
	defer close(done)

	urls := []string{"https://www.google.com", "https://badhost"}
	for result := range checkStatus(done, urls...) {
		if result.Error != nil {
			fmt.Printf("error: %v", result.Error)
			continue
		}
		fmt.Printf("Response: %v\n", result.Response.Status)
	}
}

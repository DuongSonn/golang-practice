package main

import (
	"fmt"
	"strconv"
	"time"
)

func rateLimit(rateLimitChan <-chan int) {

}

func main() {
	rateLimitChan := make(chan int, 100)
	defer close(rateLimitChan)

	for i := 0; i < 100; i++ {
		go func() {
			fmt.Println("Begin insert. Waiting... " + strconv.Itoa(i))
			rateLimitChan <- i
			fmt.Println("Finish Insert: " + strconv.Itoa(i))
		}()
	}

	for i := range rateLimitChan {
		fmt.Println("Handling: ", strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		fmt.Println("Finish Handling: ", strconv.Itoa(i))
	}
}

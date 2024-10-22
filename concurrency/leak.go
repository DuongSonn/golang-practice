package main

import (
	"fmt"
	"time"
)

/*
NOTE:  If a goroutine is responsible for creating a goroutine, it is also responsible for ensuring it can stop the goroutine
*/

/*
The above code will cause deadlock because when go loop through the nil channel. It will loop forever => Never run to defer func and completed will not be closed
*/
func leakProblem() {
	doWork := func(strings <-chan string) <-chan interface{} {
		completed := make(chan interface{})
		go func() {
			defer fmt.Println("doWork exited.")
			defer close(completed)
			for s := range strings {
				// Do something interesting
				fmt.Println(s)
			}
		}()
		return completed
	}
	<-doWork(nil)
	// Perhaps more work is done here
	fmt.Println("Done.")
}

/*
The above solution setup a done channel. After 1 second close the done channel so when the do work receive the done channel close event it will break the loop
*/
func leakSolution() {
	doWork := func(
		done <-chan interface{},
		strings <-chan string,
	) <-chan interface{} {
		terminated := make(chan interface{})
		go func() {
			defer fmt.Println("doWork exited.")
			defer close(terminated)
			for {
				select {
				case s := <-strings:
					// Do something interesting
					fmt.Println(s)
				case _, ok := <-done:
					if !ok {
						return
					}
				}
			}
		}()
		return terminated
	}

	done := make(chan interface{})
	terminated := doWork(done, nil)
	go func() {
		// Cancel the operation after 1 second.
		time.Sleep(1 * time.Second)
		fmt.Println("Canceling doWork goroutine...")
		close(done)
	}()

	<-terminated
	fmt.Println("Done.")

}

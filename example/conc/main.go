package main

import (
	"fmt"
	"time"

	"github.com/sourcegraph/conc"
)

func task() {
	fmt.Println("Doing Task ...")
	time.Sleep(2 * time.Second)
	fmt.Println("Done Task")
}

func main() {
	var wg conc.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Go(task)
	}
	// panics with a nice stacktrace
	wg.Wait()
}

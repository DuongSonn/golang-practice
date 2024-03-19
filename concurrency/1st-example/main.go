package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	words := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

	wg.Add(len(words))

	for i, char := range words {
		go printSomeThing(fmt.Sprintf("%d - %s", i, char), &wg)
	}

	wg.Wait()

	wg.Add(1)
	printSomeThing("Hello World 2", &wg)
}

func printSomeThing(s string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println(s)
}

package main

import "fmt"

/*
func adHocConfinement:
- data variable is available globally but only access inside loopData func
*/
func adHocConfinement() {
	data := make([]int, 4)
	loopData := func(handleData chan<- int) {
		defer close(handleData)
		for i := range data {
			handleData <- data[i]
		}
	}

	handleData := make(chan int)
	go loopData((handleData))

	for num := range handleData {
		fmt.Println(num)
	}
}

/*
func lexicalConfinement:
- we have 2 func chanOwner and consumer.
The chanOwner create the channel and handle writing inside the channel => Prevent other routine to write into the channel
The consumer only read data from the channel
- Another way is split the data into small part and each routine handle 1 part separately
*/
func lexicalConfinement() {
	chanOwner := func() <-chan int {
		results := make(chan int, 5)
		go func() {
			defer func() {
				fmt.Println("Closing channel!")
				close(results)
			}()
			for i := 0; i <= 5; i++ {
				results <- i
			}
		}()

		return results
	}

	consumer := func(results <-chan int) {
		for result := range results {
			fmt.Printf("Received %d\n", result)
		}

		fmt.Println("Done receiving!")
	}

	results := chanOwner()
	consumer(results)
}

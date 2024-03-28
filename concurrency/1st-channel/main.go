package main

import (
	"fmt"
	"strings"
)

func shout(ping <-chan string, pong chan<- string) {
	for {
		s := <-ping
		pong <- fmt.Sprintf("%s!!!", strings.ToUpper(s))
	}
}

func main() {
	ping := make(chan string)
	pong := make(chan string)

	go shout(ping, pong)

	fmt.Println("Type something and press ENTER (enter Q to quit)")
	for {
		fmt.Print("->")

		var input string
		_, _ = fmt.Scanln(&input)
		if input == strings.ToLower("q") {
			break
		}

		ping <- input
		fmt.Println(<-pong)
	}

	close(ping)
	close(pong)
}

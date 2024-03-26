package main

import (
	"fmt"
	"sync"
	"time"
)

type Philosopher struct {
	name      string
	rightFork int
	leftFork  int
}

var philosophers = []Philosopher{
	{name: "Plato", rightFork: 0, leftFork: 4},
	{name: "Sam", rightFork: 1, leftFork: 0},
	{name: "Peter", rightFork: 2, leftFork: 1},
	{name: "Candy", rightFork: 3, leftFork: 2},
	{name: "Carl", rightFork: 4, leftFork: 3},
}

var hunger = 3
var eatTime = 1 * time.Second
var think = 3 * time.Second
var sleepTime = 1 * time.Second
var finishOrder = 1

func main() {
	fmt.Println("Dining Philosophers Problem")
	fmt.Println("-----------------------------")
	fmt.Println("The table is empty")

	dine()

	fmt.Println("The table is empty")
}

func dine() {
	wg := &sync.WaitGroup{}
	wg.Add(len(philosophers))

	seated := &sync.WaitGroup{}
	seated.Add(len(philosophers))

	forks := make(map[int]*sync.Mutex)
	finishMutex := &sync.Mutex{}

	for i := 0; i < len(philosophers); i++ {
		forks[i] = &sync.Mutex{}
	}

	for i := 0; i < len(philosophers); i++ {
		go diningProblem(philosophers[i], wg, forks, seated, finishMutex)
	}

	wg.Wait()
}

func diningProblem(p Philosopher, wg *sync.WaitGroup, forks map[int]*sync.Mutex, seated *sync.WaitGroup, finishMutex *sync.Mutex) {
	defer wg.Done()

	fmt.Printf("%s is seated at the table.\n", p.name)
	seated.Done()

	seated.Wait()

	for i := 0; i < hunger; i++ {
		if p.leftFork < p.rightFork {
			forks[p.leftFork].Lock()
			fmt.Printf("\t%s take the left fork.\n", p.name)

			forks[p.rightFork].Lock()
			fmt.Printf("\t%s take the right fork.\n", p.name)
		} else {
			forks[p.rightFork].Lock()
			fmt.Printf("\t%s take the right fork.\n", p.name)
			forks[p.leftFork].Lock()
			fmt.Printf("\t%s take the left fork.\n", p.name)
		}

		fmt.Printf("\t %s is eating.\n", p.name)
		time.Sleep(eatTime)

		fmt.Printf("\t %s is thinking.\n", p.name)
		time.Sleep(think)

		forks[p.leftFork].Unlock()
		forks[p.rightFork].Unlock()

		fmt.Printf("\t%s put down the forks.\n", p.name)
	}

	finishMutex.Lock()
	fmt.Printf("%s finish number %d\n", p.name, finishOrder)
	finishOrder++
	finishMutex.Unlock()

	fmt.Println(p.name, "is sleeping")
	fmt.Println(p.name, "lef the table")
}

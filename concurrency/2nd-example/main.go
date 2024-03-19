package main

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

type Income struct {
	Source string
	Amount int
}

func main() {
	var totalBalance int
	var mutex sync.Mutex
	incomes := []Income{
		{"a", 100},
		{"b", 200},
		{"c", 300},
		{"d", 400},
		{"e", 500},
		{"f", 600},
	}

	fmt.Println("Initial balance: ", totalBalance)

	wg.Add(len(incomes))
	for _, income := range incomes {
		go func(income Income, m *sync.Mutex) {
			defer wg.Done()

			for i := 0; i < 52; i++ {
				m.Lock()
				totalBalance += income.Amount

				fmt.Printf("Balance: %d in Week: %d from Source: %s \n", totalBalance, i, income.Source)
				m.Unlock()

			}
		}(income, &mutex)
	}
	wg.Wait()

	fmt.Println("Final balance: ", totalBalance)
}

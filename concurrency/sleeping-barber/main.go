package main

import (
	"fmt"
	"math/rand"
	"time"
)

var seatingCapacity = 10
var arrivalRate = 100
var cutDuration = 1000 * time.Microsecond
var timeOpen = 10 * time.Second

type BarberShop struct {
	ShopCapacity    int
	HaircutDuration time.Duration
	NumberOfBarber  int
	BarberDoneChan  chan bool
	ClientChan      chan string
	Open            bool
}

func (s *BarberShop) addBarber(barber string) {
	s.NumberOfBarber++

	go func() {
		isSleeping := false
		fmt.Printf("Barber %s is waiting for clients\n", barber)

		for {
			if len(s.ClientChan) == 0 {
				fmt.Printf("Barber %s go to sleep\n", barber)
				isSleeping = true
			}

			client, ok := <-s.ClientChan
			if ok {
				if isSleeping {
					fmt.Printf("Barber %s wakes up\n", barber)
					isSleeping = false
				}

				s.cutHair(barber, client)
			} else {
				s.sendBarberHome(barber)
				return
			}
		}
	}()
}

func (s *BarberShop) cutHair(barber, client string) {
	fmt.Printf("Barber %s is cutting %s's hair\n", barber, client)
	time.Sleep(s.HaircutDuration)
	fmt.Printf("Barber %s finished cutting %s's hair\n", barber, client)
}

func (s *BarberShop) sendBarberHome(barber string) {
	fmt.Printf("Barber %s is going home\n", barber)
	s.BarberDoneChan <- true
}

func (s *BarberShop) closeShopForDay() {
	fmt.Println("The shop is closed for the day")

	close(s.ClientChan)
	s.Open = false

	for i := 1; i <= s.NumberOfBarber; i++ {
		<-s.BarberDoneChan
	}

	close(s.BarberDoneChan)
	fmt.Println("-----------------------------")
	fmt.Println("The barber shop is closed, everyone is gone")
}

func (s *BarberShop) addClient(client string) {
	fmt.Println("Adding new client to waiting list ", client)
	if s.Open {
		select {
		case s.ClientChan <- client:
			fmt.Println("Client", client, "added to queue")
		default:
			fmt.Println("The shop is full, client", client, "unable to join")
		}
	} else {
		fmt.Println("The shop is closed, no more clients can be added")
	}
}

func main() {
	rand.Seed(time.Now().Unix())

	fmt.Println("The sleeping barber problem")
	fmt.Println("-----------------------------")

	clientChan := make(chan string, seatingCapacity)
	doneChan := make(chan bool)

	shop := &BarberShop{
		ShopCapacity:    seatingCapacity,
		HaircutDuration: cutDuration,
		NumberOfBarber:  0,
		BarberDoneChan:  doneChan,
		ClientChan:      clientChan,
		Open:            true,
	}
	fmt.Println("The shop is open for the day")

	shop.addBarber("Barber1")

	shopClosing := make(chan bool)
	closed := make(chan bool)
	go func() {
		<-time.After(timeOpen)
		shopClosing <- true
		shop.closeShopForDay()
		closed <- true
	}()

	i := 1
	go func() {
		for {
			rand := rand.Int() % (2 * arrivalRate)
			select {
			case <-shopClosing:
				return
			case <-time.After(time.Duration(rand) * time.Millisecond):
				shop.addClient(fmt.Sprintf("Client%d", i))
				i++
			}
		}
	}()

	<-closed
}

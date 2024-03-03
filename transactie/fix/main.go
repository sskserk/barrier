package main

import (
	"fmt"
	"sync"
)

type Account struct {
	amount int
	lock   sync.Mutex
}

func (account *Account) Inc(amount int) {
	account.amount += amount
}

func (account *Account) Dec(amount int) {
	account.amount -= amount
}

func (account *Account) Lock() {
	account.lock.Lock()
}

func (account *Account) Unlock() {
	account.lock.Unlock()
}

const numberOfTransactions = 100

func main() {
	waitChannel := make(chan bool, numberOfTransactions)

	var a = Account{amount: 1000, lock: sync.Mutex{}}
	var b = Account{amount: 1000, lock: sync.Mutex{}}

	transactionsStartSignal := sync.WaitGroup{}
	transactionsStartSignal.Add(numberOfTransactions * 2)

	routineA := func(amount int) {
		transactionsStartSignal.Done()
		transactionsStartSignal.Wait()
		a.Lock()
		b.Lock()

		defer a.Unlock()
		defer b.Unlock()

		a.Dec(amount) // a = a - amount
		b.Inc(amount) // b = b - amount

		waitChannel <- true
	}

	routineB := func(amount int) {
		transactionsStartSignal.Done()
		transactionsStartSignal.Wait()

		a.Lock()
		b.Lock()

		defer a.Unlock()
		defer b.Unlock()

		b.Dec(amount) // b = b - amount
		a.Inc(amount) // a = a + amount

		waitChannel <- true
	}

	for i := 0; i < numberOfTransactions; i++ {
		go routineA(10)
		go routineB(10)
	}

	for i := 0; i < numberOfTransactions * 2; i++ {
		select {

		case <-waitChannel:
		}
	}
	//ime.Sleep(1 * time.Second)

	fmt.Println(a, b)
}

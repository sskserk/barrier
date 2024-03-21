package main

import (
	"fmt"
	"sync"
)

type Account struct {
	id        int
	amount    int
	suspended bool
	lock      sync.Mutex
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

func (account *Account) SetSuspended(isSuspended bool) {
	account.suspended = isSuspended
}

func (account *Account) IsSuspended() bool {
	return account.suspended
}

const numberOfTransactions = 100

func main() {
	waitChannel := make(chan bool, numberOfTransactions*2)

	var a = Account{id: 1, amount: 1000, lock: sync.Mutex{}}
	var b = Account{id: 2, amount: 1000, lock: sync.Mutex{}}

	transactionsStartSignal := sync.WaitGroup{}
	transactionsStartSignal.Add(numberOfTransactions * 2)

	transferRoutine := func(sourceAccount *Account, destinationAccount *Account, amount int) {
		transactionsStartSignal.Done()
		transactionsStartSignal.Wait()

		leftAccount := sourceAccount
		rightAccount := destinationAccount
		if sourceAccount.id > destinationAccount.id {
			leftAccount = destinationAccount
			rightAccount = sourceAccount
		}
		leftAccount.Lock()
		rightAccount.Lock()

		defer leftAccount.Unlock()
		defer rightAccount.Unlock()

		if !sourceAccount.IsSuspended() && !destinationAccount.IsSuspended() {
			sourceAccount.Dec(amount)      // left = left - amount
			destinationAccount.Inc(amount) // right = right + amount
		}

		waitChannel <- true
	}

	for i := 0; i < numberOfTransactions; i++ {
		go transferRoutine(&a, &b, 10)
		go transferRoutine(&b, &a, 10)
	}

	for i := 0; i < numberOfTransactions*2; i++ {
		<-waitChannel
	}

	fmt.Println(a.amount, b.amount)
}

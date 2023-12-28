package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	N = 3
)

func main() {
	readyToStart := sync.WaitGroup{}
	readyToStart.Add(N)

	shotSignal := sync.WaitGroup{}
	shotSignal.Add(1)

	for i := 0; i < N; i++ {
//		wgReady.Add(1)

		go routine(i, &readyToStart, &shotSignal)
	}

//	wgReady.Wait()

	time.Sleep(5 * time.Second)
	readyToStart.Wait()

	fmt.Println(time.Now(), "Ready, steady, go...")

	
	shotSignal.Done()
//	wg.Done()
//	time.Sleep(5 * time.Second)
}

func routine(id int, readyToStart *sync.WaitGroup, shotSignal *sync.WaitGroup) {
	fmt.Printf("Gorotuine [%d] is ready ro run...\n", id)
	readyToStart.Done()

	shotSignal.Wait() // wait for the shot

	fmt.Println(time.Now(), id)

	
}

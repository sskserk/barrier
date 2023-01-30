package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	N = 100
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)

	wgReady := sync.WaitGroup{}

	for i := 0; i < N; i++ {

		wgReady.Add(1)

		go routine(i, &wg, &wgReady)
	}

	wgReady.Wait()

	time.Sleep(5 * time.Second)

	fmt.Println(time.Now(), "Ready, steady, go...")
	wg.Done()
	time.Sleep(5 * time.Second)
}

func routine(id int, wg *sync.WaitGroup, wgReady *sync.WaitGroup) {
	fmt.Printf("Gorotuine [%d] is ready ro run...\n", id)
	wgReady.Done()

	wg.Wait()

	fmt.Println(time.Now(), id)
}

package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	BARRIERS_NUMBER = 10 // number of cycles/barriers
	RUNNERS_NUMBER  = 5  // number of runners
)

func main() {
	readyToStartConfirmation := sync.WaitGroup{}
	starterPistol := sync.WaitGroup{}
	barrierReached := sync.WaitGroup{}
	proceedToTheNextStartPermission := sync.WaitGroup{}

	readyToStartConfirmation.Add(RUNNERS_NUMBER)
	starterPistol.Add(1)

	// spin-up runners
	for i := 0; i < RUNNERS_NUMBER; i++ {
		go runner(i, BARRIERS_NUMBER, &readyToStartConfirmation, &starterPistol, &barrierReached, &proceedToTheNextStartPermission)
	}

	// run cycles
	for barrierNumber := 0; barrierNumber < BARRIERS_NUMBER; barrierNumber++ {
		barrierReached.Add(RUNNERS_NUMBER)

		// make sure all runners are ready
		readyToStartConfirmation.Wait()
		logMark("====== Shot to start from the barrier [%d] ======", barrierNumber)

		readyToStartConfirmation.Add(RUNNERS_NUMBER)

		// fire the signal
		starterPistol.Done()
		proceedToTheNextStartPermission.Add(1)

		// wait for all runners to reach the next barrier
		barrierReached.Wait()
		starterPistol.Add(1)

		logMark("====== Barrier [%d] reached by all runners ======", barrierNumber+1)

		// grant permission to prepare for the next start
		proceedToTheNextStartPermission.Done()
	}
}

func runner(runnerNumber int,
	maxNumberOfBarriers int,
	readyToStartConfirmation *sync.WaitGroup,
	starterPistol *sync.WaitGroup,
	barrierReached *sync.WaitGroup,
	proceedToTheNextStartPermission *sync.WaitGroup,
) {
	for barrierNumber := 0; barrierNumber < maxNumberOfBarriers; barrierNumber++ {
		// confirm readiness to run
		logMark("Runner [%d] ready to run from the barrier [%d]", runnerNumber, barrierNumber)
		readyToStartConfirmation.Done()

		// wait for the shot to start running
		starterPistol.Wait()
		logMark("Runner [%d] started from the barrier [%d]...", runnerNumber, barrierNumber)

		// do something. Every runner has different velocity
		time.Sleep(time.Duration(runnerNumber+1) * time.Second)

		// finish of the current sprint is reched
		barrierReached.Done()
		logMark("Runner [%d] reached the barrier [%d]", runnerNumber, barrierNumber+1)

		// wait signal to proceed to start of the next barrier
		proceedToTheNextStartPermission.Wait()
	}
}

// print a log message regardless of the actual moment at which the logging function physically writes the log message
func logMark(format string, args ...interface{}) {
	fmt.Println(fmt.Sprintf("%-40s", time.Now().UTC()), fmt.Sprintf(format, args...))
}

package main

type Scheduler struct {
	ListOfTaskIDs []int
}

func (sc *Scheduler) Run() {
	// Do something....
}

func main() {
	scheduler := Scheduler{ListOfTaskIDs: []int{0, 1}}

	go scheduler.Run()

}

package main

import (
	"sync"
)

func main() {
	n := 5

	wg := new(sync.WaitGroup)
	wg.Add(n)

	forkChans := make([]chan ForkMsg, n)
	philChans := make([]chan string, n)

	for i := range forkChans {
		forkChans[i] = make(chan ForkMsg)
		go SpawnFork(forkChans[i])
	}

	for i := range philChans {
		philChans[i] = make(chan string)
		go SpawnPhilo(wg, forkChans[i])
	}

	wg.Wait()
}

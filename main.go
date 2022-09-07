package main

import "sync"

type ForkMsg struct {
	message  string
	response bool
}

func spawnFork(ch chan ForkMsg) {
	isClean := false
	msg := <-ch
	switch msg.message {
	case "isClean":
		ch <- ForkMsg{"isClean", isClean}
	case "setIsClean":
		isClean = msg.response
		ch <- ForkMsg{"isClean", isClean}
	}

}

func main() {
	n := 5

	wg := new(sync.WaitGroup)

	forkChans := make([]chan ForkMsg, n)
	philChans := make([]chan string, n)

	for i := range forkChans {
		forkChans[i] = make(chan ForkMsg)
	}

	for i := range forkChans {
		philChans[i] = make(chan string)
	}

	wg.Wait()

}

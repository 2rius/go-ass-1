package main

import (
	"sync"
	"time"
)

type Philo struct {
	id int
	eaten int
	state
	left Fork
	right Fork
	defered [2]bool
}

func SpawnPhilo(wg *sync.WaitGroup, ch chan ForkMsg) {
	time.Sleep(2 * time.Second)
	wg.Done()
}

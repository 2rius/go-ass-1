package main

import (
	"fmt"
	"sync"
)

func prompt() (int, int) {
	var philCount, mealCount int

	fmt.Print("How many philosophers and chopsticks? (default: 5, minimum: 3): ")
	if _, err := fmt.Scan(&philCount); err != nil {
		panic(err)
	} else {
		if philCount <= 2 {
			panic("Number of philosophers have to be atleast 3")
		}
	}

	fmt.Print("How many meals each? (default: 3, minimum: 1): ")
	if _, err := fmt.Scan(&mealCount); err != nil {
		panic(err)
	} else {
		if mealCount <= 0 {
			panic("Number of meals have to be atleast 1")
		}
	}

	return philCount, mealCount
}

func main() {
	fmt.Println()

	n := 5
	mealCount := 3

	// n, mealCount = prompt()

	wg := new(sync.WaitGroup)
	wg.Add(n)

	forkChans := make([]chan Msg, n)
	philChans := make([]chan Msg, n)

	for i := range forkChans {
		forkChans[i] = make(chan Msg)
		fork := Fork{i, forkChans[i], false}
		go SpawnFork(fork)
	}

	for i := range philChans {
		philChans[i] = make(chan Msg)
	}

	for i := range philChans {
		forks := [2]chan Msg{nil, nil}

		if i == 0 {
			forks[LEFT] = forkChans[i]
			forks[RIGHT] = forkChans[n-1]
		} else if i == n-1 {
		} else {
			forks[LEFT] = forkChans[i]
		}

		// Remember orientation
		neighbours := [3]chan Msg{philChans[Mod(i+1, n)], philChans[Mod(i-1, n)], philChans[i]}

		phil := Philo{i, 0, THINKING, forks, neighbours, [2]chan Msg{nil, nil}}

		go SpawnPhilo(wg, phil, mealCount)
	}

	wg.Wait()

	fmt.Println()
}

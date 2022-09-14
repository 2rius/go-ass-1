package main

import (
	"fmt"
	"sync"
)

//Prompting the user to input how many philosophers and how many meals they should eat
func prompt() (int, int) {
	var philCount, mealCount int

	fmt.Print("How many philosophers and chopsticks? (default: 5, minimum: 3): ")

	for {
		if _, err := fmt.Scan(&philCount); err != nil {
			fmt.Println("Invalid input. Please try again.")
		} else {
			if philCount < 3 {
				fmt.Println("Please enter a number greater than 2")
			} else {
				break
			}
		}
	}

	fmt.Print("How many meals each? (default: 3, minimum: 1): ")

	for {
		if _, err := fmt.Scan(&mealCount); err != nil {
			fmt.Println("Invalid input. Please try again.")
		} else {
			if mealCount < 1 {
				fmt.Println("Please enter a number greater than 0")
			} else {
				break
			}
		}
	}

	return philCount, mealCount
}

//Spawning the philosophers and chopsticks and starting the simulation of the dining philosophers problem
func main() {
	n := 5
	mealCount := 3

	n, mealCount = prompt()

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
		} else if i != n-1 {
			forks[LEFT] = forkChans[i]
		}

		neighbours := [3]chan Msg{philChans[Mod(i+1, n)], philChans[Mod(i-1, n)], philChans[i]}

		phil := Philo{i, 0, THINKING, forks, neighbours, [2]chan Msg{nil, nil}}

		go SpawnPhilo(wg, phil, mealCount)
	}

	wg.Wait()
}

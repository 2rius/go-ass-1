package main

import (
	"fmt"
	"sync"
)

func SpawnPhilo(wg *sync.WaitGroup, ph Philo, meals int) {
	calledWg := false

	for {
		if ph.eaten >= meals {
			fmt.Printf("(id: %d) has eaten all meals!\n", ph.id)

			// [TODO] gets called again
			if !calledWg {
				wg.Done()
				calledWg = true
			}
		}

		switch ph.state {
		case THINKING:
			fmt.Printf("(id: %d) THINKING\n", ph.id)
			ph.thinking()
		case HUNGRY:
			fmt.Printf("(id: %d) HUNGRY\n", ph.id)
			ph.hungry()
		case EATING:
			ph.eaten++
			fmt.Printf("(id: %d) EATING: %d\n", ph.id, ph.eaten)
			ph.eating()
		}
	}
}

type Philo struct {
	id    int
	eaten int
	state
	forks      [2]chan Msg
	neighbours [3]chan Msg
	defered    [2]chan Msg
}

func (ph *Philo) thinking() {
	for {
		select {
		case request := <-ph.neighbours[ME]:
			switch request.message {

			case "giveyourleft":
				// fmt.Printf("%d giving my left\n", ph.id)
				ph.forks[LEFT] <- Msg{"request", "clean!", nil}
				msg := Msg{"response", "here", ph.forks[LEFT]}
				request.recv <- msg
				ph.forks[LEFT] = nil
				ph.defered[LEFT] = nil

			case "giveyourright":
				// fmt.Printf("%d giving my right\n", ph.id)
				ph.forks[RIGHT] <- Msg{"request", "clean!", nil}
				msg := Msg{"response", "here", ph.forks[RIGHT]}
				request.recv <- msg
				ph.forks[RIGHT] = nil
				ph.defered[RIGHT] = nil
			}
		default:
		}

		left, right := ph.checkForks()

		// [TODO] rækkefølge
		switch {
		// If either is dirty we're still thinking
		case left == DIRTY || right == DIRTY:
			ph.state = THINKING // Still - redundent

		// If both clean we're prepared to eat
		case left == CLEAN && right == CLEAN:
			ph.state = EATING
			return

		// If we have none - we are hungry
		case left == NONE || right == NONE:
			fallthrough

		// If either is clean we're still hungry
		case left == CLEAN || right == CLEAN:
			ph.state = HUNGRY
			return
		}
	}
}

func (ph *Philo) hungry() {
	var msgLeft Msg
	var msgRight Msg

	if ph.forks[LEFT] == nil {
		// fmt.Printf("%d give me my left!\n", ph.id)
		msgLeft = Msg{"request", "giveyourright", make(chan Msg)}
		ph.neighbours[LEFT] <- msgLeft
	}

	if ph.forks[RIGHT] == nil {
		// fmt.Printf("%d give me my right!\n", ph.id)
		msgRight = Msg{"request", "giveyourleft", make(chan Msg)}
		ph.neighbours[RIGHT] <- msgRight
	}

	for {
		select {
		case request := <-ph.neighbours[ME]:
			// fmt.Printf("%d NOT YET\n", ph.id)
			switch request.message {
			case "giveyourleft":
				ph.defered[LEFT] = request.recv
			case "giveyourright":
				ph.defered[RIGHT] = request.recv
			}

		case response := <-msgLeft.recv:
			// fmt.Printf("%d getting my left\n", ph.id)
			ph.forks[LEFT] = response.recv

		case response := <-msgRight.recv:
			// fmt.Printf("%d getting my right\n", ph.id)
			ph.forks[RIGHT] = response.recv
		}


		if ph.forks[LEFT] != nil && ph.forks[RIGHT] != nil {
			ph.state = EATING
			return
		}
	}
}

func (ph *Philo) eating() {
	if ph.defered[LEFT] != nil {
		// fmt.Printf("%d giving my left to defered\n", ph.id)
		ph.defered[LEFT] <- Msg{"response", "here", ph.forks[LEFT]}
		ph.forks[LEFT] = nil
		ph.defered[LEFT] = nil
	} else {
		ph.forks[LEFT] <- Msg{"request", "dirty!", nil}
	}

	if ph.defered[RIGHT] != nil {
		// fmt.Printf("%d giving my right to defered\n", ph.id)
		ph.defered[RIGHT] <- Msg{"response", "here", ph.forks[RIGHT]}
		ph.forks[RIGHT] = nil
		ph.defered[RIGHT] = nil
	} else {
		ph.forks[RIGHT] <- Msg{"request", "dirty!", nil}
	}

	ph.state = THINKING
}

func (ph *Philo) checkForks() (forkstate, forkstate) {
	var left, right forkstate

	if ph.forks[LEFT] == nil {
		left = NONE
	} else {
		msgLeft := Msg{"request", "isClean?", make(chan Msg)}
		ph.forks[LEFT] <- msgLeft

		answer := <-msgLeft.recv

		switch answer.message {
		case "true":
			left = CLEAN
		case "false":
			left = DIRTY
		}
	}

	if ph.forks[RIGHT] == nil {
		right = NONE
	} else {
		msgRight := Msg{"request", "isClean?", make(chan Msg)}
		ph.forks[RIGHT] <- msgRight

		answer := <-msgRight.recv

		switch answer.message {
		case "true":
			right = CLEAN
		case "false":
			right = DIRTY
		}
	}

	return left, right
}

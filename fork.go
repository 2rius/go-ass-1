package main

func SpawnFork(ch chan ForkMsg) {
	fork := Fork{false, ch}

	msg := <-fork.channel
	switch msg.message {
	case "isClean":
		ch <- ForkMsg{"isClean", fork.isClean}
	case "setIsClean":
		fork.isClean = msg.response
		ch <- ForkMsg{"isClean", fork.isClean}
	}
}

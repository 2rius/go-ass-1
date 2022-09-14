package main

import (
	"strconv"
)

// Constructor contains fork id, channel, and isClean
type Fork struct {
	id      int
	channel chan Msg
	isClean bool
}

// SpawnFork spawns a fork, and listens for messages
func SpawnFork(fork Fork) {
	for {
		msg := <-fork.channel

		switch {
		case msg.message == "isClean?":
			msg.recv <- Msg{"response", strconv.FormatBool(fork.isClean), nil}
		case msg.message == "clean!":
			fork.isClean = true
		case msg.message == "dirty!":
			fork.isClean = false
		}
	}
}

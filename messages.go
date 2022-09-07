package main

// Fork Msg
type Fork struct {
	isClean bool
	channel chan ForkMsg
}

type ForkMsg struct {
	message  string
	response bool
}

// Philo Msg
type state int
const (
	THINKING state = iota
	HUNGRY = iota
	EATING = iota
)

type msgtype int
const (
	SEND msgtype = iota
	RECEIVE = iota
)

type PhiloMsg struct {
	msgtype
	msg string
}

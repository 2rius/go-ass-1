package main

// Forkstate
type forkstate int

const (
	NONE  forkstate = iota
	DIRTY           = iota
	CLEAN           = iota
)

// Philo state
type state int

const (
	THINKING state = iota
	HUNGRY         = iota
	EATING         = iota
)

// Orientation
type orientation int

const (
	LEFT  orientation = iota
	RIGHT             = iota
	ME                = iota
)

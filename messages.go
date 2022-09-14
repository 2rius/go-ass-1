package main

// Enum for message types
type msgtype int

const (
	SEND    msgtype = iota
	RECEIVE         = iota
)

// Blueprint for a message
type Msg struct {
	msgtype string
	message string
	recv    chan Msg
}

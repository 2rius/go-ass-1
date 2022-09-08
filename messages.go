package main

type msgtype int

const (
	SEND    msgtype = iota
	RECEIVE         = iota
)

type Msg struct {
	msgtype string
	message string
	recv    chan Msg
}

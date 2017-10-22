package main

type state uint8
type player interface {
	init([][]byte)
	move(string, *game)
}

// States
const (
	ALIVE state = iota
	FALLING
	DIGGING
	TRAPPED
)

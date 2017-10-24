package main

type player interface {
	init([][]byte)
	move(string, *game)
}

type position struct{ x, y int }
type state uint8

// States
const (
	ALIVE state = iota
	FALLING
	DIGGING
	TRAPPED
)

package game

type Player interface {
	init([][]byte)
	Move(string, *Game)
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

package main

import "log"

type game struct {
	lvl     *level
	players map[*client]*player
}

type player interface {
	move()
}

func newGame() *game {
	return &game{
	//players: make([]player, 0, 2),
	}
}

func (g *game) start() {
	g.lvl = newLevel(1)
	log.Println(g)
}

func (g *game) stop() {
}

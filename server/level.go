package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

type level struct {
	num  int
	grid [][]byte
}

// Tiles
const (
	// TODO Order/Rename
	EMPTY        = ' '
	RUNNER       = '&'
	GUARD        = '0'
	BRICK        = '#'
	SOLIDBRICK   = '@'
	FALSEBRICK   = 'X'
	LADDER       = 'H'
	ESCAPELADDER = 'S'
	ROPE         = '-'
	GOLD         = '$'
)

func newLevel(num int) (*level, error) {
	filename := fmt.Sprintf("levels/%03d.lvl", num)

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return &level{num, bytes.Split(content, []byte("\n"))}, nil
}

func (l *level) print() {
	for _, row := range l.grid {
		fmt.Println(string(row))
	}
}

func (l *level) emptyBelow(pos position) bool {
	return l.grid[pos.y+1][pos.x] == EMPTY
}

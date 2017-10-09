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

type position struct{ x, y int }

const (
	//SPACE         = ' '
	//NORMAL_BRICK  = '#'
	//SOLID_BRICK   = '@'
	//NORMAL_LADDER = 'H'
	//ROPE          = '-'
	//FALSE_BRICK   = 'X'
	//ESCAPE_LADDER = 'S'
	//GOLD          = '$'
	GUARD  = '0'
	RUNNER = '&'
)

// TODO
func newLevel() *level {
	level := &level{}
	go level.init(1)
	return level
}

func (l *level) init(num int) error {
	filename := fmt.Sprintf("levels/%03d.lvl", num)
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	l.num = num
	l.grid = bytes.Split(content, []byte("\n"))
	return nil
}

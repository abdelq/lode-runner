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

const (
	SPACE         = ' '
	RUNNER        = '@'
	GUARD         = '&'
	GOLD          = '$'
	ROPE          = '-'
	NORMAL_BRICK  = '#'
	SOLID_BRICK   = 'X'
	FAKE_BRICK    = 'V'
	NORMAL_LADDER = 'H'
	ESCAPE_LADDER = '|'
)

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

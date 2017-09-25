package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

type level struct {
	num  uint8
	grid [][]byte
}

const (
	RUNNER = '@'
)

func (l *level) init(num uint8) error {
	filename := fmt.Sprintf("levels/%03d.lvl", num)
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	l.num = num
	l.grid = bytes.Split(content, []byte("\n"))
	return nil
}

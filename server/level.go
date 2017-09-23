package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
)

// TODO enum w/ grid signification elements

type level struct {
	num  uint8
	grid [][]byte
}

func newLevel(num uint8) *level {
	filename := fmt.Sprintf("levels/%03d.lvl", num)
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println(err)
		return nil
	}

	return &level{num, bytes.Split(content, []byte("\n"))}
}

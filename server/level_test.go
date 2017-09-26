package main

import "testing"

func TestInit(t *testing.T) {
	lvl := &level{}
	for i := 1; i <= 150; i++ {
		if err := lvl.init(i); err != nil {
			t.Error(err)
		}
	}
}

package main

import "testing"

func TestInit(t *testing.T) {
	lvl := &level{}
	for i := 1; i <= 150; i++ {
		if err := lvl.init(i); err != nil {
			t.Error(err)
			//continue
		}

		if len(lvl.grid) != 16 {
			t.Fail()
			//continue
		}
		for _, row := range lvl.grid {
			if len(row) != 26 {
				t.Fail()
				//continue
			}
		}
	}
}

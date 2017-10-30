package game

import "testing"

// TODO Rename
var emptyBelowTests = []struct {
	pos          position
	isEmptyBelow bool
}{
	{position{15, 3}, true},
	{position{7, 1}, false},
	{position{24, 10}, true},
	{position{14, 8}, false},
	{position{7, 10}, true},
	{position{11, 11}, false},
}

// TODO Every single level
func TestNewLevel(t *testing.T) {}

func TestEmptyBelow(t *testing.T) {
	lvl, err := newLevel(1)
	if err != nil {
		t.Error(err)
		return
	}

	for _, test := range emptyBelowTests {
		isEmptyBelow := lvl.emptyBelow(test.pos)
		if isEmptyBelow != test.isEmptyBelow {
			t.Errorf("level 1, position %+v: %t, want %t",
				test.pos, isEmptyBelow, test.isEmptyBelow)
		}
	}
}

func TestValidMove(t *testing.T) {
	lvl, _ := newLevel(1)
	lvl.validMove(position{0, 0}, UP)
}

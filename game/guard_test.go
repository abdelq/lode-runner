package game

import "testing"

var guardInitTests = []struct {
	lvl       int
	positions []position
}{
	{1, []position{{5, 6}, {23, 6}, {14, 9}}},
	{2, []position{{4, 3}, {20, 5}, {4, 8}}},
	{3, []position{{7, 4}, {15, 7}, {19, 9}}},
	{4, []position{{6, 6}, {22, 6}, {14, 10}}},
	{5, []position{{24, 1}, {1, 5}, {4, 8}, {15, 11}}},
}

// TODO Transform into test helper
func posInPositions(pos position, positions []position) bool {
	for _, position := range positions {
		if pos == position {
			return true
		}
	}
	return false
}

func TestGuardInit(t *testing.T) {
	guard := &guard{}
	for _, test := range guardInitTests {
		lvl, err := newLevel(test.lvl)
		if err != nil {
			t.Error(err)
			return
		}
		guard.init(lvl.grid)

		if !posInPositions(guard.pos, test.positions) {
			t.Fail()
			//t.Errorf("level %d: %+v, want %+v", test.lvl, runner.pos, test.pos)
		}
	}
}

// TODO
func TestGuardMove(t *testing.T) {}

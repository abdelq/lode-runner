package main

import "testing"

var initTests = []struct {
	lvl int
	pos position
}{
	{1, position{14, 14}},
	{2, position{12, 14}},
	{3, position{10, 14}},
	{4, position{18, 14}},
	{5, position{15, 14}},
}

func TestRunnerInit(t *testing.T) {
	runner := &runner{}
	for _, test := range initTests {
		runner.init(newLevel(test.lvl).grid)
		if runner.pos != test.pos {
			t.Errorf("%+v, want %+v", runner.pos, test.pos) // TODO
		}
	}
}

// TODO
func TestRunnerMove(t *testing.T) {}

// TODO
func TestRunnerDig(t *testing.T) {}

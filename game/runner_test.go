package game

import "testing"

var runnerInitTests = []struct {
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
	runner := new(runner)
	for _, test := range runnerInitTests {
		lvl, err := newLevel(test.lvl)
		if err != nil {
			t.Error(err)
			return
		}
		runner.init(lvl.grid)

		if runner.pos != test.pos {
			t.Errorf("level %d: %+v, want %+v", test.lvl, runner.pos, test.pos)
		}
	}
}

// TODO
func TestRunnerMove(t *testing.T) {}

// TODO
func TestRunnerDig(t *testing.T) {}

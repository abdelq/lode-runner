package game

import (
	"errors"
	"sync"

	msg "github.com/abdelq/lode-runner/message"
)

type Runner struct {
	name     string
	pos      position
	state    uint8
	action   action
	health   uint8
	startLvl uint8
	out      chan *msg.Message
}

func (r *Runner) Join(game *Game) error {
	if game.runner != nil {
		return errors.New("runner already joined")
	}
	if game.hasPlayer(r.name) {
		return errors.New("name already used")
	}

	game.runner = r
	/*game.broadcast <- &msg.Message{"join",
		[]byte(`{"name": "", "room": "", "role": "runner"}`), // FIXME
	}*/

	if game.filled() {
		if r.startLvl > 0 {
			go game.start(int(r.startLvl))
		} else {
			go game.start(1)
		}
	}

	return nil
}

func (r *Runner) Leave(game *Game) {
	//game.runner = nil // XXX Commented bc of generated panics
	/*game.broadcast <- &msg.Message{"leave",
		[]byte(`{"name": "", "room": "", "role": "runner"}`), // FIXME
	}*/

	if game.Started() {
		game.stop(GUARD)
	}
}

func (r *Runner) init(players sync.Map) {
	r.action = action{}
	players.Range(func(pos, tile interface{}) bool {
		if tile.(uint8) == RUNNER {
			r.pos = pos.(position)
			return false
		}
		return true
	})
}

// FIXME FIXME FIXME FIXME
func (r *Runner) move(dir uint8, game *Game) {
	// Stop falling if needed
	if r.state == FALLING && r.pos.y+1 < game.level.height() {
		var nextTile = game.level.tiles[r.pos.y+1][r.pos.x]

		if nextTile == BRICK || nextTile == BLOCK ||
			nextTile == LADDER || nextTile == HLADDER {
			r.state = ALIVE
		}
	}

	if r.state == FALLING && dir != DOWN {
		dir = DOWN
	}

	// FIXME
	var newPos position
	switch dir {
	case NONE:
		newPos = position{r.pos.x, r.pos.y}
	case UP:
		newPos = position{r.pos.x, r.pos.y - 1}
	case LEFT:
		newPos = position{r.pos.x - 1, r.pos.y}
	case DOWN:
		newPos = position{r.pos.x, r.pos.y + 1}
	case RIGHT:
		newPos = position{r.pos.x + 1, r.pos.y}
	}

	// Stop falling
	if r.state == FALLING && r.pos.y+1 < game.level.height() {
		var nextTile = game.level.tiles[r.pos.y+1][r.pos.x]

		if nextTile == ROPE {
			r.state = ALIVE
		}
	}

	if !game.level.validMove(r.pos, newPos, dir) {
		if r.state == FALLING {
			r.state = ALIVE
		}
		return
	}

	if game.level.goldCollected() &&
		game.level.getTiles()[newPos.y][newPos.x] == HLADDER {
		game.start(game.level.num + 1)
		return
	}

	r.collectGold(r.pos, game.level)
	game.level.players.Delete(r.pos)
	r.pos.x, r.pos.y = newPos.x, newPos.y // FIXME

	// Collision checking
	if _, ok := game.level.players.Load(r.pos); ok {
		r.health--

		if r.health == 0 {
			game.stop(GUARD) // TODO Goroutine?
			return
		}

		//game.start(game.level.num) // TODO Goroutine or not ?
		game.restart()
		return
		// TODO Reset
	}

	game.level.players.Store(r.pos, tile(RUNNER))

	if game.level.emptyBelow(r.pos) &&
		game.level.tiles[r.pos.y][r.pos.x] != ROPE &&
		game.level.tiles[r.pos.y][r.pos.x] != LADDER &&
		game.level.tiles[r.pos.y][r.pos.x] != HLADDER &&
		game.level.tiles[r.pos.y][r.pos.x] != GOLD {
		r.state = FALLING
	}
}

// FIXME FIXME FIXME FIXME
func (r *Runner) dig(dir uint8, game *Game) {
	if r.state == FALLING && r.pos.y+1 < game.level.height() {
		var nextTile = game.level.tiles[r.pos.y+1][r.pos.x]

		if nextTile == BRICK || nextTile == BLOCK ||
			nextTile == LADDER || nextTile == HLADDER {
			r.state = ALIVE
		} else {
			r.move(DOWN, game)
			return
		}
	}

	// FIXME
	var digPos position
	if dir == RIGHT {
		digPos = position{r.pos.x + 1, r.pos.y + 1}
	} else if dir == LEFT {
		digPos = position{r.pos.x - 1, r.pos.y + 1}
	} else {
		return // XXX Should be catched b4 getting here
	}

	// FIXME FIXME
	if game.level.validDig(digPos) {
		//r.state = DIGGING
		game.level.holes[digPos] = 8
		game.level.tiles[digPos.y][digPos.x] = EMPTY
	}
}

func (r *Runner) Move(direction uint8) {
	r.action = action{MOVE, direction}
}

func (r *Runner) Dig(direction uint8) {
	r.action = action{DIG, direction}
}

// FIXME FIXME FIXME
func (r *Runner) collectGold(pos position, lvl *level) {
	for i, p := range lvl.gold {
		if p == pos {
			lvl.gold[i] = lvl.gold[len(lvl.gold)-1]
			lvl.gold = lvl.gold[:len(lvl.gold)-1]
			return
		}
	}
}

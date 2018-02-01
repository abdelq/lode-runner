package game

import (
	"errors"

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

func (r *Runner) init(players map[position]tile) {
	r.action = action{}
	for pos, tile := range players {
		if tile == RUNNER {
			r.pos = pos
			return
		}
	}
}

// FIXME FIXME FIXME FIXME
func (r *Runner) move(dir uint8, game *Game) {
	// Stop falling if needed
	if r.state == FALLING && r.pos.y+1 < game.level.height()-1 {
		var nextTile = game.level.tiles[r.pos.y+1][r.pos.x]

		if nextTile == BRICK ||
			nextTile == SOLIDBRICK ||
			nextTile == LADDER {
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
	if r.state == FALLING && r.pos.y+1 < game.level.height()-1 {
		var nextTile = game.level.tiles[r.pos.y+1][r.pos.x]

		if nextTile == ROPE {
			r.state = ALIVE
		}
	}

	var validMove = game.level.validMove(r.pos, newPos, dir)

	if !validMove {
		if r.state == FALLING {
			r.state = ALIVE
		}
		return
	}

	if game.level.goldCollected() && newPos.y < 0 {
		game.start(game.level.num + 1)
		return
	}

	r.collectGold(r.pos, game.level)
	delete(game.level.players, r.pos)
	r.pos.x, r.pos.y = newPos.x, newPos.y // FIXME

	// Collision checking
	if _, ok := game.level.players[r.pos]; ok {
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

	game.level.players[r.pos] = RUNNER

	if game.level.emptyBelow(r.pos) && game.level.tiles[r.pos.y][r.pos.x] != ROPE && game.level.tiles[r.pos.y][r.pos.x] != GOLD {
		r.state = FALLING
	}
}

// FIXME FIXME FIXME FIXME
func (r *Runner) dig(dir uint8, game *Game) {
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

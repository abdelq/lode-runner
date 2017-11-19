package game

import (
	"errors"
	"log"
	"time"

	msg "github.com/abdelq/lode-runner/message"
)

type Runner struct {
	name   string
	pos    position
	state  state
	action action
	health uint8
	out    chan *msg.Message
}

func (r *Runner) Add(game *Game) error {
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
		go game.start(1)
	}

	return nil
}

func (r *Runner) Remove(game *Game) {
	game.runner = nil
	/*game.broadcast <- &msg.Message{"leave",
		[]byte(`{"name": "", "room": "", "role": "runner"}`), // FIXME
	}*/

	if game.Started() {
		game.stop(GUARD)
	}
}

func (r *Runner) init(players map[position]tile) {
	for pos, tile := range players {
		if tile == RUNNER {
			r.pos = pos
			return
		}
	}
}

// FIXME FIXME FIXME FIXME
func (r *Runner) move(dir direction, game *Game) {
	if r.state == DIGGING {
		return
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

	if r.pos.y+1 >= game.level.height()-1 {
		return
	}

	var nextTile = game.level.tiles[r.pos.y+1][r.pos.x]
	var validMove = game.level.validMove(r.pos, newPos, dir)

	// Stop falling
	if r.state == FALLING && (nextTile == LADDER || nextTile == ESCAPELADDER || nextTile == ROPE) {
		r.state = ALIVE
	}

	if !validMove { // FIXME
		//log.Println("invalid move")
		if r.state == FALLING {
			r.state = ALIVE
		}
		return
	}

	if game.level.getTiles()[r.pos.y+1][r.pos.x] == ESCAPELADDER {
		game.start(game.level.num + 1)
		return
	}

	/*if newPos.y < 0 {
		//if game.level.escape[] { // TP2
		game.start(game.level.num + 1)
		return
	}*/

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

		game.start(game.level.num) // TODO Goroutine or not ?
		return
		// TODO Reset
	}

	game.level.players[r.pos] = RUNNER

	if game.level.emptyBelow(r.pos) && game.level.tiles[r.pos.y][r.pos.x] != ROPE {
		r.state = FALLING
	}
}

// FIXME FIXME FIXME FIXME
func (r *Runner) dig(dir direction, game *Game) {
	// FIXME
	var digPos position
	if dir == RIGHT {
		digPos = position{r.pos.x + 1, r.pos.y + 1}
	} else {
		digPos = position{r.pos.x - 1, r.pos.y + 1}
	}

	// FIXME FIXME
	if game.level.validDig(digPos) {
		//r.state = DIGGING
		game.level.tiles[digPos.y][digPos.x] = EMPTY

		digDuration, err := time.ParseDuration("1000ms") // TODO Using flag ?
		if err != nil {
			log.Println(err)
			digDuration, _ = time.ParseDuration("1000ms") // TODO Forced to default
		}

		time.AfterFunc(digDuration, func() {
			if tile, ok := game.level.players[digPos]; ok && tile == RUNNER {
				r.health--

				if r.health == 0 {
					game.stop(GUARD) // TODO Goroutine?
					return
				}

				game.start(game.level.num) // TODO Goroutine or not ?
				return
			}
			// FIXME For guard
			game.level.tiles[digPos.y][digPos.x] = BRICK
		})
	}
}

// FIXME FIXME FIXME
func (r *Runner) UpdateAction(actionType uint8, direction direction) {
	r.action = action{actionType, direction}
	// FIXME Dig should only accept left/right (DO THIS HERE)
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

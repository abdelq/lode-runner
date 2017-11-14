package game

import (
	"errors"
	"log"
	"time"
)

import "fmt"

type Runner struct {
	Name   string
	pos    *position
	state  state
	health uint8
	Action Action
}

func (r *Runner) Add(game *Game) error {
	if game.runner != nil {
		return errors.New("runner already joined")
	}
	if game.hasPlayer(r.Name) {
		return errors.New("name already used")
	}

	game.runner = r
	//game.broadcast <- msg.NewMessage("join", r.Name) // FIXME Join Msg ?

	if game.filled() {
		go game.start(1)
	}

	return nil
}

func (r *Runner) Remove(game *Game) {
	game.runner = nil
	//game.broadcast <- msg.NewMessage("leave", r.Name) // FIXME Join Msg ?

	if game.Started() {
		go game.stop()
	}
}

func (r *Runner) init(landmarks map[position]tile) {
	for pos, tile := range landmarks {
		if tile == RUNNER {
			r.pos = &pos
			return
		}
	}
}

// TODO Broadcast
func (r *Runner) Move(dir direction, game *Game) {
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

	var nextTile = game.level.tiles[r.pos.y][r.pos.x]
	var validMove = game.level.validMove(*r.pos, newPos, dir)

	// Stop falling
	if r.state == FALLING && (!validMove || nextTile == LADDER || nextTile == ESCAPELADDER) {
		r.state = ALIVE
	}

	if !validMove { // FIXME

		if r.state == FALLING {
			r.state = ALIVE
		}

		return
	}

	if newPos.y < 0 {
		game.start(game.level.num + 1)
		return
	}

	// FIXME
	//fmt.Println("validmove")
	//fmt.Println(r.pos)
	//fmt.Println(*r.pos)
	//delete(game.level.gold, *r.pos) // FIXME
	//game.level.gold[i] = game.level.gold[len(a)-1]
	game.level.collectGold(*r.pos)
	delete(game.level.players, *r.pos)
	r.pos.x, r.pos.y = newPos.x, newPos.y // FIXME

	// Collision checking
	if _, ok := game.level.players[*r.pos]; ok {
		r.health--

		if r.health == 0 {
			game.stop() // TODO Goroutine?
			return
		}

		game.start(game.level.num) // TODO Goroutine or not ?
		return
		// TODO Reset
	}

	game.level.players[*r.pos] = RUNNER

	if game.level.emptyBelow(*r.pos) && game.level.tiles[r.pos.y][r.pos.x] != ROPE {
		r.state = FALLING
	}

	fmt.Println(game.level.String())
}

// TODO Broadcast
func (r *Runner) Dig(dir direction, game *Game) {
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

		digDuration, err := time.ParseDuration("320ms") // TODO Using flag ?
		if err != nil {
			log.Println(err)
			digDuration, _ = time.ParseDuration("320ms") // TODO Forced to default
		}

		time.AfterFunc(digDuration, func() {
			if tile, ok := game.level.players[digPos]; ok && tile == RUNNER {
				r.health--

				if r.health == 0 {
					game.stop() // TODO Goroutine?
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

func (r *Runner) UpdateAction(actionType string, direction direction) {
	r.Action = Action{actionType, direction}
}

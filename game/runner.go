package game

import (
	"errors"
	"log"
	"time"
	//msg "github.com/abdelq/lode-runner/message"
)

import "fmt"

type Runner struct {
	Name   string
	pos    *position
	state  state
	health uint8 // TODO Use
}

func (r *Runner) Add(game *Game) error {
	if game.Runner != nil {
		return errors.New("runner already joined")
	}
	if game.hasPlayer(r.Name) {
		return errors.New("name already used")
	}

	game.Runner = r
	//game.broadcast <- msg.NewMessage("join", r.Name) // FIXME Join Msg ?

	if game.filled() {
		go game.start(1)
	}

	return nil
}

func (r *Runner) Remove(game *Game) {
	game.Runner = nil
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

	if !game.Level.validMove(*r.pos, newPos, dir) { // FIXME
		//fmt.Println("notvalidmove")
		if r.state == FALLING {
			r.state = ALIVE
		}
		return
	}

	// FIXME
	//fmt.Println("validmove")
	//fmt.Println(r.pos)
	//fmt.Println(*r.pos)
	//delete(game.Level.gold, *r.pos) // FIXME
	//game.Level.gold[i] = game.Level.gold[len(a)-1]
	game.Level.collectGold(*r.pos)
	delete(game.Level.players, *r.pos)
	r.pos.x, r.pos.y = newPos.x, newPos.y // FIXME

	// Collision checking
	if _, ok := game.Level.players[*r.pos]; ok {
		r.health--
		game.start(game.Level.num) // TODO Goroutine or not ?
		return
		// TODO Reset
	}

	game.Level.players[*r.pos] = RUNNER

	if dir == DOWN || (game.Level.emptyBelow(*r.pos) && game.Level.tiles[r.pos.y][r.pos.x] != ROPE) {
		r.state = FALLING
	}

	fmt.Println(game.Level.String())
}

// TODO Broadcast
func (r *Runner) Dig(dir direction, lvl *level) {
	// FIXME
	var digPos position
	if dir == RIGHT {
		digPos = position{r.pos.x + 1, r.pos.y + 1}
	} else {
		digPos = position{r.pos.x - 1, r.pos.y + 1}
	}

	// FIXME FIXME
	if lvl.validDig(digPos) {
		//r.state = DIGGING
		lvl.tiles[digPos.y][digPos.x] = EMPTY

		digDuration, err := time.ParseDuration("320ms") // TODO Using flag ?
		if err != nil {
			log.Println(err)
			digDuration, _ = time.ParseDuration("320ms") // TODO Forced to default
		}

		time.AfterFunc(digDuration, func() {
			lvl.tiles[digPos.y][digPos.x] = BRICK
			// TODO Check if player in position and kill him/respawn him
		})
	}
}

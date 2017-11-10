package game

import "errors"

type Runner struct {
	Name  string
	pos   position
	state state
}

// TODO Maybe just use landmarks directly
func (r *Runner) init(game *Game) {
	for pos, tile := range game.Level.landmarks {
		if tile == RUNNER {
			r.pos = pos
			delete(game.Level.landmarks, pos) // TODO Test
			return
		}
	}
}

func (r *Runner) Move(dir direction, lvl *level) {
	if r.state == DIGGING {
		return
	}

	if r.state == FALLING && dir != DOWN {
		dir = DOWN
	}

	// TODO Improve
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

	if !lvl.validMove(r.pos, newPos, dir) {
		if r.state == FALLING {
			r.state = ALIVE
		}
		return
	}

	r.pos.x, r.pos.y = newPos.x, newPos.y // TODO Improve

	if dir == DOWN || lvl.emptyBelow(r.pos) {
		r.state = FALLING
	}

	// TODO
	//gfx_move_sprite(HERO, orig, hero.pos)
	//game.check_collisions()
}

func (r *Runner) Dig(dir direction, lvl *level) {} // TODO TODO

func (r *Runner) Add(game *Game) error {
	if game.Runner != nil {
		return errors.New("runner already joined")
	}
	if game.hasPlayer(r.Name) {
		return errors.New("name already used")
	}

	game.Runner = r
	//game.broadcast <- newJoinMessage(p.name, 0) // TODO

	if game.filled() {
		go game.start()
	}

	return nil
}

func (r *Runner) Remove(game *Game) {
	game.Runner = nil
	//game.broadcast <- newLeaveMessage(p.name, 0) // TODO

	if game.Started() {
		go game.stop()
	}
}
